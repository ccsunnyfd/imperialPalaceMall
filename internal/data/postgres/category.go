package postgres

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
)

type Category struct {
	Id           int64
	CategoryName string
}

type CategoryClient struct {
	client *Client
}

func NewCategoryClient(client *Client) *CategoryClient {
	return &CategoryClient{
		client: client,
	}
}

func (c *CategoryClient) Create(ctx context.Context, category *Category) (*Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CategoryClient) Update(ctx context.Context, category *Category) (*Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CategoryClient) Get(ctx context.Context, id int64) (*Category, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CategoryClient) Query(ctx context.Context) ([]*Category, error) {
	query := `
		SELECT id, category_name
		FROM categories
		ORDER BY id`

	rows, err := c.client.db.QueryContext(ctx, query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	categories := make([]*Category, 0)
	for rows.Next() {
		var category Category

		err := rows.Scan(
			&category.Id,
			&category.CategoryName,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
