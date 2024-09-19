package service

import (
	pb "post-servic/genproto/post"
)

func (p *PostService) CreatePost(*pb.Post) (*pb.PostResponse, error) {
	return &pb.PostResponse{}, nil
}
func (p *PostService) UpdatePost(*pb.UpdateAPost) (*pb.PostResponse, error) {
	return &pb.PostResponse{}, nil
}
func (p *PostService) GetPostByID(*pb.PostId) (*pb.PostResponse, error) {
	return &pb.PostResponse{}, nil
}
func (p *PostService) ListPosts(*pb.PostList) (*pb.PostListResponse, error) {
	return &pb.PostListResponse{}, nil
}
func (p *PostService) DeletePost(*pb.PostId) (*pb.Message, error) {
	return &pb.Message{}, nil
}
func (p *PostService) RemoveImageFromPost(*pb.Post) (*pb.PostListResponse, error) {
	return &pb.PostListResponse{}, nil
}
func (p *PostService) GetPostByCountry(*pb.PostCountry) (*pb.PostListResponse, error) {
	return &pb.PostListResponse{}, nil
}
