package service

import (
	pb "post-servic/genproto/post"
)

func (c *PostService) CreateComment(*pb.CommentPost) (*pb.LikePost, error) {
	return nil, nil
}

func (c *PostService) UpdateComment(*pb.UpdateAComment) (*pb.Message, error) {
	return nil, nil
}

func (c *PostService) GetCommentByID(*pb.CommentId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (c *PostService) GetCommentByUsername(*pb.UserPostId) (*pb.Message, error) {
	return nil, nil
}

func (c *PostService) ListComments(*pb.CommentList) (*pb.Message, error) {
	return nil, nil
}

func (c *PostService) DeleteComment(*pb.CommentId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (c *PostService) GetCommentByPostID(*pb.PostId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (c *PostService) GetAllUserComments(*pb.CommentId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (c *PostService) GetMostlikeCommentPost(*pb.PostId) (*pb.PostListResponse, error) {
	return nil, nil
}
