package service

import (
	"context"
	pb "imperialPalaceMall/api/mall/v1"
	"imperialPalaceMall/internal/biz"
	"imperialPalaceMall/internal/filters"
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
	g, err := s.uc.Get(ctx, req.Id)
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
