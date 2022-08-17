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
var ProviderSet = wire.NewSet(NewData, NewCategoryRepo, NewGoodsRepo)

// Data .
type Data struct {
	//db  *mock.Client
	db *postgres.Client
}

// NewData .
func NewData(conf *conf.Data, logger log.Logger) (*Data, func(), error) {
	logH := log.NewHelper(log.With(logger, "module", "mall/data"))

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

	client := postgres.NewClient(db)
	d := &Data{
		db: client,
	}

	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
