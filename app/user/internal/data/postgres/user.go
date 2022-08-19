package postgres

import (
	"context"
	_ "github.com/lib/pq"
)

type User struct {
	Id int64
}

type UserClient struct {
	client *Client
}

func NewUserClient(client *Client) *UserClient {
	return &UserClient{
		client: client,
	}
}

func (c *UserClient) Create(ctx context.Context, user *User) (*User, error) {
	//TODO implement me
	panic("implement me")
}

func (c *UserClient) Query(ctx context.Context) ([]*User, error) {
	//query := `
	//	SELECT id, category_name
	//	FROM categories
	//	ORDER BY id`
	//
	//rows, err := c.client.db.QueryContext(ctx, query)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	//
	//defer rows.Close()
	//
	//categories := make([]*Category, 0)
	//for rows.Next() {
	//	var category Category
	//
	//	err := rows.Scan(
	//		&category.Id,
	//		&category.CategoryName,
	//	)
	//	if err != nil {
	//		return nil, err
	//	}
	//	categories = append(categories, &category)
	//}
	//
	//if err = rows.Err(); err != nil {
	//	return nil, err
	//}
	//
	return []*User{}, nil
}
