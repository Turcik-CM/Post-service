package service

import (
	"context"
	pb "post-servic/genproto/post"
)

func (p *PostService) GetUserRecommendation(ctx context.Context, in *pb.Username) (*pb.PostListResponse, error) {
	res, err := p.additional.GetUserRecommendation(in)
	if err != nil {
		p.logger.Error("Error in Getting User Recommendation", "err", err)
		return nil, err
	}

	return res, nil
}

func (p *PostService) GetPostsByUsername(ctx context.Context, in *pb.Username) (*pb.PostListResponse, error) {
	res, err := p.additional.GetPostsByUsername(in)
	if err != nil {
		p.logger.Error("Error in Getting Posts By Username", "err", err)
		return nil, err
	}

	return res, nil
}

func (p *PostService) GetTrendsPost(ctx context.Context, in *pb.Void) (*pb.PostListResponse, error) {
	res, err := p.additional.GetTrendsPost(in)
	if err != nil {
		p.logger.Error("Error in Getting TrendsPost", "err", err)
		return nil, err
	}

	return res, nil
}

func (p *PostService) SearchPost(ctx context.Context, in *pb.Search) (*pb.PostListResponse, error) {
	res, err := p.additional.SearchPost(in)
	if err != nil {
		p.logger.Error("Error in Searching Post", "err", err)
		return nil, err
	}

	return res, nil
}
