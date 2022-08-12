package postgres

import (
	"context"
	_ "github.com/lib/pq"
	"imperialPalaceMall/internal/filters"
)

type GoodsInfo struct {
	Id      int64
	Kind    int32
	Content string
}

type Goods struct {
	Id         int64
	SpuNo      string
	GoodsName  string
	StartPrice float64
	CategoryId int64
	BrandId    int64
}

type GoodsDetail struct {
	Id         int64
	SpuNo      string
	GoodsName  string
	StartPrice float64
	CategoryId int64
	BrandId    int64
	Infos      []*GoodsInfo
}

type GoodsSimplify struct {
	Id         int64
	SpuNo      string
	GoodsName  string
	StartPrice float64
	CategoryId int64
	BrandId    int64
	Info       GoodsInfo
}

type GoodsClient struct {
	client *Client
}

func NewGoodsClient(client *Client) *GoodsClient {
	return &GoodsClient{
		client: client,
	}
}

func (c *GoodsClient) Create(ctx context.Context, goods *Goods) (*Goods, error) {
	//TODO implement me
	panic("implement me")
}

func (c *GoodsClient) Update(ctx context.Context, goods *Goods) (*Goods, error) {
	//TODO implement me
	panic("implement me")
}

func (c *GoodsClient) Get(ctx context.Context, id int64) (*GoodsDetail, error) {
	//query := `
	//	SELECT a.id, a.spu_no, a.goods_name, a.start_price, a.category_id, a.brand_id,
	//		(SELECT jsonb_agg(jsonb_build_object('id', b.id, 'kind', b.kind, 'content', b.content) order by b.id)
	//			FROM goods_info as b WHERE b.goods_id = $1 GROUP BY b.goods_id)
	//		as info
	//	FROM goods as a
	//	WHERE a.id=$1`
	query := `SELECT a.id, a.spu_no, a.goods_name, a.start_price, a.category_id, a.brand_id, b.id as goods_info_id, b.kind, b.content
	FROM goods as a, goods_info as b WHERE a.id = b.goods_id AND a.id = $1 ORDER BY b.id ASC`

	rows, err := c.client.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var goods GoodsDetail

	i := 0
	for rows.Next() {
		var g GoodsSimplify
		err := rows.Scan(
			&g.Id,
			&g.SpuNo,
			&g.GoodsName,
			&g.StartPrice,
			&g.CategoryId,
			&g.BrandId,
			&g.Info.Id,
			&g.Info.Kind,
			&g.Info.Content,
		)
		if err != nil {
			return nil, err
		}
		if i == 0 {
			goods = GoodsDetail{
				Id:         g.Id,
				SpuNo:      g.SpuNo,
				GoodsName:  g.GoodsName,
				StartPrice: g.StartPrice,
				CategoryId: g.CategoryId,
				BrandId:    g.BrandId,
				Infos: []*GoodsInfo{
					&GoodsInfo{
						Id:      g.Info.Id,
						Kind:    g.Info.Kind,
						Content: g.Info.Content,
					},
				},
			}
		} else {
			goods.Infos = append(goods.Infos, &GoodsInfo{
				Id:      g.Info.Id,
				Kind:    g.Info.Kind,
				Content: g.Info.Content,
			})
		}

		i++
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	//if err != nil {
	//	switch {
	//	case errors.Is(err, sql.ErrNoRows):
	//		return nil, biz.ErrGoodsNotFound
	//	default:
	//		return nil, err
	//	}
	//}

	return &goods, nil
}

func (c *GoodsClient) Query(ctx context.Context, categoryId int64, f filters.Filters) ([]*GoodsSimplify, filters.Metadata, error) {
	// 取三张图片中的一张就行
	query := `
		 SELECT count(*) OVER(), s.id, s.spu_no, s.goods_name, s.start_price, s.category_id, s.brand_id, s.goods_info_id, s.kind, s.content
			FROM (
				SELECT a.id, a.spu_no, a.goods_name, a.start_price, a.category_id, a.brand_id, b.id as goods_info_id, b.kind, b.content
				, row_number() over (partition by a.id order by b.id asc) as group_idx
				FROM goods as a LEFT JOIN goods_info as b ON a.id = b.goods_id                                       
				 WHERE (a.category_id = $1 OR $1 = 0) AND b.kind = 0
			) s
			WHERE s.group_idx = 1
		 ORDER BY s.id DESC
		 LIMIT  $2 OFFSET $3`

	args := []interface{}{categoryId, f.Limit(), f.Offset()}
	rows, err := c.client.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, filters.Metadata{}, err
	}

	defer rows.Close()

	var totalRecords int64 = 0
	goodsList := make([]*GoodsSimplify, 0)
	for rows.Next() {
		var goods GoodsSimplify

		err := rows.Scan(
			&totalRecords,
			&goods.Id,
			&goods.SpuNo,
			&goods.GoodsName,
			&goods.StartPrice,
			&goods.CategoryId,
			&goods.BrandId,
			&goods.Info.Id,
			&goods.Info.Kind,
			&goods.Info.Content,
		)
		if err != nil {
			return nil, filters.Metadata{}, err
		}
		goodsList = append(goodsList, &goods)
	}

	if err = rows.Err(); err != nil {
		return nil, filters.Metadata{}, err
	}

	metadata := filters.CalculateMetadata(totalRecords, f.Page, f.PageSize)
	return goodsList, metadata, nil
}
