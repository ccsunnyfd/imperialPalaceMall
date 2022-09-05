package postgres

import (
	"context"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	"imperialPalaceMall/app/pkg/filters"
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
	GoodsDesc  string
	StartPrice float64
	CategoryId int64
	BrandId    int64
}

type GoodsAttrKey struct {
	Id      int64
	AttrKey string
}

type GoodsAttrValue struct {
	Id        int64
	AttrKeyId int64
	AttrValue string
}

type GoodsSKU struct {
	Id            int64
	GoodsId       int64
	GoodsAttrPath []int64
	Price         int64
	Stock         int64
}

type GoodsAndSkuDetail struct {
	GoodsAttrPath []int64
	Price         int64
	Stock         int64
	GoodsName     string
	GoodsDesc     string
	GoodsImage    string
}

type GoodsSimplify struct {
	Id         int64
	SpuNo      string
	GoodsName  string
	GoodsDesc  string
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

func (c *GoodsClient) Get(ctx context.Context, id int64) (*Goods, error) {
	//query := `
	//	SELECT a.id, a.spu_no, a.goods_name, a.start_price, a.category_id, a.brand_id,
	//		(SELECT jsonb_agg(jsonb_build_object('id', b.id, 'kind', b.kind, 'content', b.content) order by b.id)
	//			FROM goods_info as b WHERE b.goods_id = $1 GROUP BY b.goods_id)
	//		as info
	//	FROM goods as a
	//	WHERE a.id=$1`
	//query := `SELECT a.id, a.spu_no, a.goods_name, a.goods_desc, a.start_price, a.category_id, a.brand_id, b.id as goods_info_id, b.kind, b.content
	//FROM goods as a, goods_info as b WHERE a.id = b.goods_id AND a.id = $1 ORDER BY b.id ASC`

	query := `SELECT id, spu_no, goods_name, goods_desc, start_price, category_id, brand_id FROM goods WHERE id = $1`

	var goods Goods

	err := c.client.db.QueryRowContext(ctx, query, id).Scan(
		&goods.Id,
		&goods.SpuNo,
		&goods.GoodsName,
		&goods.GoodsDesc,
		&goods.StartPrice,
		&goods.CategoryId,
		&goods.BrandId,
	)

	if err != nil {
		//switch {
		//case errors.Is(err, sql.ErrNoRows):
		//	return nil, biz.ErrGoodsNotFound
		//default:
		return nil, err
	}

	return &goods, nil
}

func (c *GoodsClient) GetGoodsInfo(ctx context.Context, goodsId int64) ([]*GoodsInfo, error) {
	query := `SELECT id, kind, content FROM goods_info WHERE goods_id = $1 ORDER BY id ASC`

	rows, err := c.client.db.QueryContext(ctx, query, goodsId)
	if err != nil {
		return nil, err
	}

	goodsInfoList := make([]*GoodsInfo, 0)
	for rows.Next() {
		var info GoodsInfo
		err := rows.Scan(
			&info.Id,
			&info.Kind,
			&info.Content,
		)
		if err != nil {
			return nil, err
		}

		goodsInfoList = append(goodsInfoList, &info)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	return goodsInfoList, nil
}

func (c *GoodsClient) GetSKUs(ctx context.Context, goodsId int64) ([]*GoodsSKU, error) {
	query := `SELECT id, goods_id, goods_attr_path, price, stock FROM goods_sku WHERE goods_id = $1 ORDER BY id ASC`

	rows, err := c.client.db.QueryContext(ctx, query, goodsId)
	if err != nil {
		return nil, err
	}

	goodsSKUList := make([]*GoodsSKU, 0)
	for rows.Next() {
		var sku GoodsSKU
		err := rows.Scan(
			&sku.Id,
			&sku.GoodsId,
			pq.Array(&sku.GoodsAttrPath),
			&sku.Price,
			&sku.Stock,
		)
		if err != nil {
			return nil, err
		}

		goodsSKUList = append(goodsSKUList, &sku)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	return goodsSKUList, nil
}

func (c *GoodsClient) GetGoodsAttrKeys(ctx context.Context, id int64) ([]*GoodsAttrKey, error) {
	query := `SELECT a.id, a.attr_key FROM goods_attr_key as a WHERE a.goods_id = $1 ORDER BY a.id ASC`

	rows, err := c.client.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}

	goodsAttrKeyList := make([]*GoodsAttrKey, 0)
	for rows.Next() {
		var k GoodsAttrKey
		err := rows.Scan(
			&k.Id,
			&k.AttrKey,
		)
		if err != nil {
			return nil, err
		}

		goodsAttrKeyList = append(goodsAttrKeyList, &k)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	return goodsAttrKeyList, nil
}

func (c *GoodsClient) GetGoodsAttrValues(ctx context.Context, attrKeys []int64) ([]*GoodsAttrValue, error) {
	query := `SELECT a.id, a.attr_key_id, a.attr_value FROM goods_attr_value as a WHERE a.attr_key_id IN $1 ORDER BY a.attr_key_id ASC, a.id ASC`
	rows, err := c.client.db.QueryContext(ctx, query, pq.Array(attrKeys))
	if err != nil {
		return nil, err
	}

	goodsAttrValueList := make([]*GoodsAttrValue, 0)
	for rows.Next() {
		var v GoodsAttrValue
		err := rows.Scan(
			&v.Id,
			&v.AttrKeyId,
			&v.AttrValue,
		)
		if err != nil {
			return nil, err
		}

		goodsAttrValueList = append(goodsAttrValueList, &v)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	return goodsAttrValueList, nil
}

func (c *GoodsClient) GetGoodsAttrValuesByKeyId(ctx context.Context, attrKeyId int64) ([]*GoodsAttrValue, error) {
	query := `SELECT a.id, a.attr_key_id, a.attr_value FROM goods_attr_value as a WHERE a.attr_key_id = $1 ORDER BY a.id ASC`
	rows, err := c.client.db.QueryContext(ctx, query, attrKeyId)
	if err != nil {
		return nil, err
	}

	goodsAttrValueList := make([]*GoodsAttrValue, 0)
	for rows.Next() {
		var v GoodsAttrValue
		err := rows.Scan(
			&v.Id,
			&v.AttrKeyId,
			&v.AttrValue,
		)
		if err != nil {
			return nil, err
		}

		goodsAttrValueList = append(goodsAttrValueList, &v)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()

	return goodsAttrValueList, nil
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

func (c *GoodsClient) GetSku(ctx context.Context, skuId int64) (*GoodsSKU, error) {
	query := `SELECT id, goods_id, goods_attr_path, price, stock FROM goods_sku WHERE id = $1`

	var sku GoodsSKU

	err := c.client.db.QueryRowContext(ctx, query, skuId).Scan(
		&sku.Id,
		&sku.GoodsId,
		pq.Array(&sku.GoodsAttrPath),
		&sku.Price,
		&sku.Stock,
	)

	if err != nil {
		//switch {
		//case errors.Is(err, sql.ErrNoRows):
		//	return nil, biz.ErrSkuNotFound
		//default:
		return nil, err
	}

	return &sku, nil
}
