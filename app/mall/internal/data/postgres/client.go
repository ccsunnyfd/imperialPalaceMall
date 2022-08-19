package postgres

import "database/sql"

type Client struct {
	db       *sql.DB
	Category *CategoryClient
	Goods    *GoodsClient
}

func NewClient(db *sql.DB) *Client {
	client := &Client{db: db}
	client.Category = NewCategoryClient(client)
	client.Goods = NewGoodsClient(client)

	return client
}
