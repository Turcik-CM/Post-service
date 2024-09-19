package service

import (
	pb "post-servic/genproto/post"
)

func (l *PostService) AddLikePost(*pb.LikePost) (*pb.LikePost, error) {
	return nil, nil
}

func (l *PostService) DeleteLikePost(*pb.LikeId) (*pb.Message, error) {
	return nil, nil
}

func (l *PostService) AddLikeComment(*pb.LikeId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (l *PostService) DeleteLikeComment(*pb.UserId) (*pb.Message, error) {
	return nil, nil
}

func (l *PostService) GetPostLikeCount(*pb.PostId) (*pb.Message, error) {
	return nil, nil
}

func (l *PostService) GetCommentLikeCount(*pb.Like) (*pb.PostListResponse, error) {
	return nil, nil
}

func (l *PostService) GetUsersWhichLikePost(*pb.PostId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (l *PostService) GetUsersWhichLikeComment(*pb.CommentId) (*pb.PostListResponse, error) {
	return nil, nil
}
