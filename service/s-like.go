package service

import (
	"context"
	pb "post-servic/genproto/post"
)

func (l *PostService) AddLikePost(ctx context.Context, in *pb.LikePost) (*pb.LikeResponse, error) {
	return nil, nil
}

func (l *PostService) DeleteLikePost(ctx context.Context, in *pb.LikePost) (*pb.Message, error) {
	return nil, nil
}

func (l *PostService) AddLikeComment(ctx context.Context, in *pb.LikeCommit) (*pb.LikeComResponse, error) {
	return nil, nil
}

func (l *PostService) DeleteLikeComment(ctx context.Context, in *pb.LikeCommit) (*pb.Message, error) {
	return nil, nil
}

func (l *PostService) GetPostLikeCount(ctx context.Context, in *pb.PostId) (*pb.LikeCount, error) {
	return nil, nil
}

func (l *PostService) GetMostLikedComment(ctx context.Context, in *pb.PostId) (*pb.LikeCount, error) {
	return nil, nil
}

func (l *PostService) GetUsersWhichLikePost(ctx context.Context, in *pb.PostId) (*pb.Users, error) {
	return nil, nil
}

func (l *PostService) GetUsersWhichLikeComment(ctx context.Context, in *pb.CommentId) (*pb.Users, error) {
	return nil, nil
}
