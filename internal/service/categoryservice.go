package service

import (
	"context"
	"imperialPalaceMall/internal/biz"

	pb "imperialPalaceMall/api/mall/v1"
)

type CategoryServiceService struct {
	pb.UnimplementedCategoryServiceServer
	uc *biz.CategoryUsecase
}

func NewCategoryServiceService(uc *biz.CategoryUsecase) *CategoryServiceService {
	return &CategoryServiceService{
		uc: uc,
	}
}

func (s *CategoryServiceService) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryReply, error) {
	return &pb.CreateCategoryReply{}, nil
}
func (s *CategoryServiceService) UpdateCategory(ctx context.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryReply, error) {
	return &pb.UpdateCategoryReply{}, nil
}
func (s *CategoryServiceService) DeleteCategory(ctx context.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryReply, error) {
	return &pb.DeleteCategoryReply{}, nil
}
func (s *CategoryServiceService) GetCategory(ctx context.Context, req *pb.GetCategoryRequest) (*pb.GetCategoryReply, error) {
	return &pb.GetCategoryReply{}, nil
}
func (s *CategoryServiceService) ListCategory(ctx context.Context, req *pb.ListCategoryRequest) (*pb.ListCategoryReply, error) {
	cc, err := s.uc.ListCategories(ctx)
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
