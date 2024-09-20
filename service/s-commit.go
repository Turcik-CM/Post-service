package service

import (
	"context"
	pb "post-servic/genproto/post"
)

func (c *PostService) CreateComment(ctx context.Context, in *pb.CommentPost) (*pb.CommentResponse, error) {
	res, err := c.comment.CreateComment(in)
	if err != nil {
		c.logger.Error("failed to post commit", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) UpdateComment(ctx context.Context, in *pb.UpdateAComment) (*pb.CommentResponse, error) {
	res, err := c.comment.UpdateComment(in)
	if err != nil {
		c.logger.Error("failed to post comment", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) GetCommentByID(ctx context.Context, in *pb.CommentId) (*pb.CommentResponse, error) {
	res, err := c.comment.GetCommentByID(in)
	if err != nil {
		c.logger.Error("failed to post comment", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) GetCommentByUsername(ctx context.Context, in *pb.Username) (*pb.CommentResponse, error) {
	res, err := c.comment.GetCommentByUsername(in)
	if err != nil {
		c.logger.Error("failed to post comment by username", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) ListComments(ctx context.Context, in *pb.CommentList) (*pb.CommentsR, error) {
	res, err := c.comment.ListComments(in)
	if err != nil {
		c.logger.Error("failed to post comments", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) DeleteComment(ctx context.Context, in *pb.CommentId) (*pb.Message, error) {
	res, err := c.comment.DeleteComment(in)
	if err != nil {
		c.logger.Error("failed to delete comment", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) GetCommentByPostID(ctx context.Context, in *pb.PostId) (*pb.CommentsR, error) {
	res, err := c.comment.GetCommentByPostID(in)
	if err != nil {
		c.logger.Error("failed to post comment by post id", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) GetAllUserComments(ctx context.Context, in *pb.Username) (*pb.CommentsR, error) {
	res, err := c.comment.GetAllUserComments(in)
	if err != nil {
		c.logger.Error("failed to post comments", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) GetMostlikeCommentPost(ctx context.Context, in *pb.PostId) (*pb.CommentResponse, error) {
	res, err := c.comment.GetMostlikeCommentPost(in)
	if err != nil {
		c.logger.Error("failed to post comment", err)
		return nil, err
	}
	return res, nil
}
