package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	weixinPb "imperialPalaceMall/api/weixin/v1"
	"imperialPalaceMall/app/pkg/weixin"
	"imperialPalaceMall/app/user/internal/biz"
	"imperialPalaceMall/app/user/internal/conf"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRedis, NewWxLoginAuther, NewWxDecrypter, NewUserRepo, NewAddressRepo)

// Data .
type Data struct {
	rdb           *redis.Client
	db            *gorm.DB
	wxLoginAuther *WxAuther
	wxDecrypter   *weixin.WXBizDataCrypt
	log           *log.Helper
}

// NewData .
func NewData(db *gorm.DB, rdb *redis.Client, wxLoginAuther *WxAuther, wxDecrypter *weixin.WXBizDataCrypt, logger log.Logger) (*Data, func(), error) {
	logH := log.NewHelper(log.With(logger, "module", "user-service/data"))

	d := &Data{
		db:            db,
		rdb:           rdb,
		wxLoginAuther: wxLoginAuther,
		wxDecrypter:   wxDecrypter,
		log:           logH,
	}

	return d, func() {}, nil
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	logH := log.NewHelper(log.With(logger, "module", "user-service/data/gorm"))

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s user=%s password='%s' dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			conf.Database.Addr,
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Dbname,
			conf.Database.Port),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		logH.Fatalf("failed opening connection to mysql: %v", err)
	}

	return db
}

func NewRedis(conf *conf.Data, logger log.Logger) (*redis.Client, func()) {
	logH := log.NewHelper(log.With(logger, "module", "order-service/data/redis"))

	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Redis.Addr,
		Password:     conf.Redis.Password,
		DB:           int(conf.Redis.Db),
		DialTimeout:  conf.Redis.DialTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		PoolSize:     10,
	})

	rdb.AddHook(redisotel.TracingHook{})

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()

	err := rdb.Ping(timeout).Err()
	if err != nil {
		logH.Fatalf("redis connect error: %v", err)
	}

	return rdb, func() {
		if err := rdb.Close(); err != nil {
			logH.Error(err)
		}
	}
}

//func NewTlsConf(conf *conf.Data) *tls.Config {
//	b, err := os.ReadFile("../../configs/cert/ca.crt")
//	if err != nil {
//		panic(err)
//	}
//	cp := x509.NewCertPool()
//	if !cp.AppendCertsFromPEM(b) {
//		panic("append certs error")
//	}
//
//	addrStrArr := strings.Split(conf.Weixin.Addr, "//")
//	if len(addrStrArr) != 2 {
//		panic("weixin server address not correct")
//	}
//
//	return &tls.Config{ServerName: addrStrArr[1], RootCAs: cp}
//}

func NewWxLoginAuther(conf *conf.Data, logger log.Logger) (*WxAuther, func()) {
	logH := log.NewHelper(log.With(logger, "module", "user-service/data/weixin-login"))

	conn, err := http.NewClient(
		context.Background(),
		http.WithEndpoint(conf.Weixin.Addr),
		//http.WithTLSConfig(tlsConf),
		http.WithTimeout(10*time.Second),
		http.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		logH.Fatalf("failed opening connection to weixin server: %v", err)
	}
	return NewWxAuther(&biz.Code2sessionRequest{
			Appid:     conf.Weixin.Appid,
			Secret:    conf.Weixin.Secret,
			GrantType: conf.Weixin.GrantType,
		},
			weixinPb.NewWeixinHTTPClient(conn)), func() {
			if err := conn.Close(); err != nil {
				logH.Error(err)
			}
		}
}

func NewWxDecrypter(conf *conf.Data) *weixin.WXBizDataCrypt {
	return &weixin.WXBizDataCrypt{
		AppId: conf.Weixin.Appid,
	}
}
