package service

import (
	"context"
	pb "post-servic/genproto/post"
)

func (l *PostService) AddLikePost(ctx context.Context, in *pb.LikePost) (*pb.LikeResponse, error) {
	res, err := l.like.AddLikePost(in)
	if err != nil {
		l.logger.Error("failed to post like", err)
		return nil, err
	}
	return res, nil
}

func (l *PostService) DeleteLikePost(ctx context.Context, in *pb.LikePost) (*pb.Message, error) {
	res, err := l.like.DeleteLikePost(in)
	if err != nil {
		l.logger.Error("failed to delete like", err)
		return nil, err
	}
	return res, nil
}

func (l *PostService) AddLikeComment(ctx context.Context, in *pb.LikeCommit) (*pb.LikeComResponse, error) {
	res, err := l.like.AddLikeComment(in)
	if err != nil {
		l.logger.Error("failed to add like comment", err)
		return nil, err
	}
	return res, nil
}

func (l *PostService) DeleteLikeComment(ctx context.Context, in *pb.LikeCommit) (*pb.Message, error) {
	res, err := l.like.DeleteLikeComment(in)
	if err != nil {
		l.logger.Error("failed to delete like comment", err)
		return nil, err
	}
	return res, nil
}

func (l *PostService) GetPostLikeCount(ctx context.Context, in *pb.PostId) (*pb.LikeCount, error) {
	res, err := l.like.GetPostLikeCount(in)
	if err != nil {
		l.logger.Error("failed to get post like count", err)
		return nil, err
	}
	return res, nil
}

func (l *PostService) GetMostLikedComment(ctx context.Context, in *pb.PostId) (*pb.LikeCount, error) {
	res, err := l.like.GetMostLikedComment(in)
	if err != nil {
		l.logger.Error("failed to get most liked comment", err)
		return nil, err
	}
	return res, nil
}

func (l *PostService) GetUsersWhichLikePost(ctx context.Context, in *pb.PostId) (*pb.Users, error) {
	res, err := l.like.GetUsersWhichLikePost(in)
	if err != nil {
		l.logger.Error("failed to get users like post", err)
		return nil, err
	}
	return res, nil
}

func (l *PostService) GetUsersWhichLikeComment(ctx context.Context, in *pb.CommentId) (*pb.Users, error) {
	res, err := l.like.GetUsersWhichLikeComment(in)
	if err != nil {
		l.logger.Error("failed to get users like comment", err)
		return nil, err
	}
	return res, nil
}
