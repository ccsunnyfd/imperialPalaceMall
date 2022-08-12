package mock

import (
	"context"
)

type CategoryClient struct{}

func NewCategoryClient() *CategoryClient {
	return &CategoryClient{}
}

type Category struct {
	Id           int64
	CategoryName string
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
	titles := []string{"生活方式", "文创好物", "创意玩具", "微瑕特卖", "宫廷服饰", "甄选珠宝", "宫廷美玉"}

	res := make([]*Category, 0, len(titles))
	for i, t := range titles {
		res = append(res,
			&Category{
				Id:           int64(i) + 1,
				CategoryName: t,
			},
		)
	}
	return res, nil
}
