package data

import (
	"context"
	"database/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	userPb "imperialPalaceMall/api/user/v1"
	"imperialPalaceMall/app/user/internal/conf"
	"imperialPalaceMall/app/user/internal/data/postgres"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewWxLoginServiceClient, NewUserRepo)

// Data .
type Data struct {
	rdb           *redis.Client
	db            *postgres.Client
	wxLoginClient userPb.UserHTTPClient
}

// NewData .
func NewData(conf *conf.Data, wxLoginClient userPb.UserHTTPClient, logger log.Logger) (*Data, func(), error) {
	logH := log.NewHelper(log.With(logger, "module", "user/data"))

	db, err := sql.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		logH.Fatalf("failed opening connection to db: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		logH.Fatalf("failed opening connection to db: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
	})
	rdb.AddHook(redisotel.TracingHook{})

	dbClient := postgres.NewClient(db)

	d := &Data{
		db:            dbClient,
		rdb:           rdb,
		wxLoginClient: wxLoginClient,
	}

	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func NewWxLoginServiceClient() (userPb.UserHTTPClient, func()) {
	conn, err := http.NewClient(
		context.Background(),
		http.WithEndpoint("https://api.weixin.qq.com/sns/jscode2session"),
		http.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return userPb.NewUserHTTPClient(conn), func() {
		if err := conn.Close(); err != nil {
			log.Error(err)
		}
	}
}
