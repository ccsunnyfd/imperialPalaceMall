package data

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/lib/pq"
	"imperialPalaceMall/app/cart/internal/conf"
	"imperialPalaceMall/app/cart/internal/data/ent"
	"imperialPalaceMall/app/cart/internal/data/ent/migrate"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewCartRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(entClient *ent.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "cart-service/data"))

	d := &Data{
		db: entClient,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "cart-service/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		fmt.Sprintf(
			"host=%s user=%s password='%s' dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
			conf.Database.Addr,
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Dbname,
			conf.Database.Port),
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
