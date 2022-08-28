package service

import (
	"context"
	pb "imperialPalaceMall/api/shop/interface/v1"
)

func (s *ShopInterface) ListCategory(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryReply, error) {
	cc, err := s.catc.ListCategories(ctx)
	if err != nil {
		return nil, err
	}
	rs := make([]*pb.Category, 0)
	for _, x := range cc {
		rs = append(rs, &pb.Category{
			Id:           x.Id,
			CategoryName: x.CategoryName,
		})
	}
	return &pb.ListCategoryReply{
		Categories: rs,
	}, nil
}
