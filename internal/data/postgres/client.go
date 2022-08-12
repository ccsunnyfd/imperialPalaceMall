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

func (c *Client) Close() error {
	if err := c.db.Close(); err != nil {
		return err
	}
	return nil
}
