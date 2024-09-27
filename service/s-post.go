package service

import (
	"context"
	"fmt"
	pb "post-servic/genproto/post"
)

func (p *PostService) CreatePost(ctx context.Context, in *pb.Post) (*pb.PostResponse, error) {
	res, err := p.post.CreatePost(in)
	if err != nil {
		p.logger.Error("failed to post post", err)
		return nil, err
	}
	return res, nil
}

func (p *PostService) UpdatePost(ctx context.Context, in *pb.UpdateAPost) (*pb.PostResponse, error) {
	res, err := p.post.UpdatePost(in)
	if err != nil {
		p.logger.Error("failed to update post", err)
		return nil, err
	}
	return res, nil
}

func (p *PostService) GetPostByID(ctx context.Context, in *pb.PostId) (*pb.PostResponse, error) {
	res, err := p.post.GetPostByID(in)
	if err != nil {
		p.logger.Error("failed to GetPostByID", err)
		return nil, err
	}
	return res, nil
}

func (p *PostService) ListPosts(ctx context.Context, in *pb.PostList) (*pb.PostListResponse, error) {
	res, err := p.post.ListPosts(in)
	if err != nil {
		p.logger.Error("failed to ListPosts", err)
		return nil, err
	}
	return res, nil
}

func (p *PostService) DeletePost(ctx context.Context, in *pb.PostId) (*pb.Message, error) {
	res, err := p.post.DeletePost(in)
	if err != nil {
		p.logger.Error("failed to DeletePost", err)
		return nil, err
	}
	return res, nil
}

func (p *PostService) AddImageToPost(ctx context.Context, in *pb.ImageUrl) (*pb.PostResponse, error) {
	res, err := p.post.AddImageToPost(in)
	if err != nil {
		p.logger.Error("failed to AddImageToPost", err)
		return nil, err
	}
	return res, nil
}

func (p *PostService) RemoveImageFromPost(ctx context.Context, in *pb.ImageUrl) (*pb.Message, error) {
	res, err := p.post.RemoveImageFromPost(in)
	if err != nil {
		p.logger.Error("failed to RemoveImageFromPost", err)
		return nil, err
	}
	return res, nil
}

func (p *PostService) GetPostByCountry(ctx context.Context, in *pb.PostCountry) (*pb.PostListResponse, error) {
	res, err := p.post.GetPostByCountry(in)
	fmt.Println(in)
	if err != nil {
		p.logger.Error("failed to GetPostByCountry", err)
		return nil, err
	}
	return res, nil
}
