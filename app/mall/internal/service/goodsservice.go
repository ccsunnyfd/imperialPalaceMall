package service

import (
	"context"
	pb "imperialPalaceMall/api/mall/v1"
	"imperialPalaceMall/app/mall/internal/biz"
	"imperialPalaceMall/app/pkg/filters"
)

type GoodsServiceService struct {
	pb.UnimplementedGoodsServiceServer
	uc *biz.GoodsUsecase
}

func NewGoodsServiceService(uc *biz.GoodsUsecase) *GoodsServiceService {
	return &GoodsServiceService{
		uc: uc,
	}
}

func (s *GoodsServiceService) CreateGoods(ctx context.Context, req *pb.CreateGoodsRequest) (*pb.CreateGoodsReply, error) {
	return &pb.CreateGoodsReply{}, nil
}
func (s *GoodsServiceService) UpdateGoods(ctx context.Context, req *pb.UpdateGoodsRequest) (*pb.UpdateGoodsReply, error) {
	return &pb.UpdateGoodsReply{}, nil
}
func (s *GoodsServiceService) DeleteGoods(ctx context.Context, req *pb.DeleteGoodsRequest) (*pb.DeleteGoodsReply, error) {
	return &pb.DeleteGoodsReply{}, nil
}
func (s *GoodsServiceService) GetGoods(ctx context.Context, req *pb.GetGoodsRequest) (*pb.GetGoodsReply, error) {
	g, err := s.uc.GetGoodsDetail(ctx, req.Id.Value)
	if err != nil {
		//if errors.Is(err, biz.ErrGoodsNotFound) {
		//	return nil, pb.ErrorGoodsNotFound("goods %d not found", req.Id)
		//}
		return nil, err
	}

	var infoList = make([]*pb.GoodsInfo, 0, len(g.Infos))

	for _, info := range g.Infos {
		infoList = append(infoList, &pb.GoodsInfo{
			Id:      info.Id,
			Kind:    info.Kind,
			Content: info.Content,
		})
	}

	return &pb.GetGoodsReply{
		Goods: &pb.GoodsDetail{
			Id:         g.Id,
			SpuNo:      g.SpuNo,
			GoodsName:  g.GoodsName,
			GoodsDesc:  g.GoodsDesc,
			StartPrice: g.StartPrice,
			CategoryId: g.CategoryId,
			BrandId:    g.BrandId,
			Infos:      infoList,
		},
	}, nil
}
func (s *GoodsServiceService) ListGoods(ctx context.Context, req *pb.ListGoodsRequest) (*pb.ListGoodsReply, error) {
	f := filters.NewFilters(req.Page, req.PageSize, 1, 20)

	gg, metadata, err := s.uc.List(ctx, req.CategoryId, f)
	if err != nil {
		return nil, err
	}
	rs := make([]*pb.GoodsSimplify, 0)
	for _, x := range gg {
		rs = append(rs, &pb.GoodsSimplify{
			Id:         x.Id,
			SpuNo:      x.SpuNo,
			GoodsName:  x.GoodsName,
			StartPrice: x.StartPrice,
			CategoryId: x.CategoryId,
			BrandId:    x.BrandId,
			Info: &pb.GoodsInfo{
				Id:      x.Info.Id,
				Kind:    x.Info.Kind,
				Content: x.Info.Content,
			},
		})
	}
	return &pb.ListGoodsReply{
		GoodsList: rs,
		Metadata: &pb.Metadata{
			CurrentPage:  metadata.CurrentPage,
			PageSize:     metadata.PageSize,
			FirstPage:    metadata.FirstPage,
			LastPage:     metadata.LastPage,
			TotalRecords: metadata.TotalRecords,
		},
	}, nil
}

func (s *GoodsServiceService) GetSKUs(ctx context.Context, req *pb.GetSKUsRequest) (*pb.GetSKUsReply, error) {
	skus, attrs, err := s.uc.GetGoodsSKUs(ctx, req.Id.Value)
	if err != nil {
		return nil, err
	}

	var skuList = make([]*pb.GoodsSKU, 0, len(skus))
	for _, sku := range skus {
		skuList = append(skuList, &pb.GoodsSKU{
			Id:            sku.Id,
			GoodsId:       sku.GoodsId,
			GoodsAttrPath: sku.GoodsAttrPath,
			Price:         sku.Price,
			Stock:         sku.Stock,
		})
	}

	var attrList = make([]*pb.GoodsAttr, 0, len(attrs))
	for _, attr := range attrs {
		var attrValueList = make([]*pb.GoodsAttrValue, 0, len(attr.AttrValues))
		for _, v := range attr.AttrValues {
			attrValueList = append(attrValueList, &pb.GoodsAttrValue{
				Id:        v.Id,
				AttrKeyId: v.AttrKeyId,
				AttrValue: v.AttrValue,
			})
		}
		attrList = append(attrList, &pb.GoodsAttr{
			Id:         attr.Id,
			GoodsId:    attr.GoodsId,
			AttrKey:    attr.AttrKey,
			AttrValues: attrValueList,
		})
	}

	return &pb.GetSKUsReply{
		Skus:  skuList,
		Attrs: attrList,
	}, nil
}

func (s *GoodsServiceService) GetGoodsAndSkuDetail(ctx context.Context, req *pb.GetGoodsAndSkuDetailRequest) (*pb.GetGoodsAndSkuDetailReply, error) {
	goodsAndSkuDetail, err := s.uc.GetGoodsAndSkuDetails(ctx, req.GoodsId.Value, req.SkuId.Value)
	if err != nil {
		return nil, err
	}

	return &pb.GetGoodsAndSkuDetailReply{
		GoodsName:     goodsAndSkuDetail.GoodsName,
		GoodsDesc:     goodsAndSkuDetail.GoodsDesc,
		GoodsImage:    goodsAndSkuDetail.GoodsImage,
		Price:         goodsAndSkuDetail.Price,
		Stock:         goodsAndSkuDetail.Stock,
		GoodsAttrPath: goodsAndSkuDetail.GoodsAttrPath,
	}, nil
}
