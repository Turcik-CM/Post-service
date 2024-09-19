package service

import (
	pb "post-servic/genproto/post"
)

func (c *PostService) CreateComment(in *pb.CommentPost) (*pb.CommentResponse, error) {
	return nil, nil
}

func (c *PostService) UpdateComment(in *pb.UpdateAComment) (*pb.CommentResponse, error) {
	return nil, nil
}

func (c *PostService) GetCommentByID(in *pb.CommentId) (*pb.CommentResponse, error) {
	return nil, nil
}

func (c *PostService) GetCommentByUsername(in *pb.Username) (*pb.CommentResponse, error) {
	return nil, nil
}

func (c *PostService) ListComments(in *pb.CommentList) (*pb.CommentsR, error) {
	return nil, nil
}

func (c *PostService) eleteComment(in *pb.CommentId) (*pb.Message, error) {
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
