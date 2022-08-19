package data

import (
	"context"
	"database/sql"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"imperialPalaceMall/app/mall/internal/conf"
	"imperialPalaceMall/app/mall/internal/data/postgres"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewCategoryRepo, NewGoodsRepo)

// Data .
type Data struct {
	//db  *mock.Client
	db  *postgres.Client
	log *log.Helper
}

func NewDB(conf *conf.Data, logger log.Logger) (*postgres.Client, func()) {
	logH := log.NewHelper(log.With(logger, "modules", "mall-srvice/data/sql"))

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

	return postgres.NewClient(db), func() {
		if err := db.Close(); err != nil {
			log.Error(err)
		}
	}
}

// NewData .
func NewData(db *postgres.Client, logger log.Logger) (*Data, func(), error) {
	logH := log.NewHelper(log.With(logger, "module", "mall/data"))

	d := &Data{
		db:  db,
		log: logH,
	}

	return d, func() {}, nil
}
