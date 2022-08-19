package postgres

import "database/sql"

type Client struct {
	db   *sql.DB
	User *UserClient
}

func NewClient(db *sql.DB) *Client {
	client := &Client{db: db}
	client.User = NewUserClient(client)

	return client
}

func (c *Client) Close() error {
	if err := c.db.Close(); err != nil {
		return err
	}
	return nil
}
