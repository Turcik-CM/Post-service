package service

import (
	"context"
	pb "post-servic/genproto/post"
)

func (c *PostService) CreateComment(ctx context.Context, in *pb.CommentPost) (*pb.CommentResponse, error) {
	return nil, nil
}

func (c *PostService) UpdateComment(ctx context.Context, in *pb.UpdateAComment) (*pb.CommentResponse, error) {
	return nil, nil
}

func (c *PostService) GetCommentByID(ctx context.Context, in *pb.CommentId) (*pb.CommentResponse, error) {
	return nil, nil
}

func (c *PostService) GetCommentByUsername(ctx context.Context, in *pb.Username) (*pb.CommentResponse, error) {
	return nil, nil
}

func (c *PostService) ListComments(ctx context.Context, in *pb.CommentList) (*pb.CommentsR, error) {
	return nil, nil
}

func (c *PostService) DeleteComment(ctx context.Context, in *pb.CommentId) (*pb.Message, error) {
	return nil, nil
}

func (c *PostService) GetCommentByPostID(ctx context.Context, in *pb.PostId) (*pb.CommentsR, error) {
	return nil, nil
}

func (c *PostService) GetAllUserComments(ctx context.Context, in *pb.Username) (*pb.CommentsR, error) {
	return nil, nil
}

func (c *PostService) GetMostlikeCommentPost(ctx context.Context, in *pb.PostId) (*pb.CommentResponse, error) {
	return nil, nil
}
