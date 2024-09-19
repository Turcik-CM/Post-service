package storage

import (
	pb "post-servic/genproto/post"
)

type PostStorage interface {
	CreatePost(*pb.Post) (*pb.PostResponse, error)
	UpdatePost(*pb.UpdateAPost) (*pb.PostResponse, error)
	GetPostByID(*pb.PostId) (*pb.PostResponse, error)
	ListPosts(*pb.PostList) (*pb.PostListResponse, error)
	DeletePost(*pb.PostId) (*pb.Message, error)
	RemoveImageFromPost(*pb.Post) (*pb.PostListResponse, error)
	GetPostByCountry(*pb.PostCountry) (*pb.PostListResponse, error)
}

type LikeStorage interface {
	AddLikePost(*pb.LikePost) (*pb.LikePost, error)
	DeleteLikePost(*pb.LikeId) (*pb.Message, error)
	AddLikeComment(*pb.LikeId) (*pb.PostListResponse, error)
	DeleteLikeComment(*pb.UserId) (*pb.Message, error)
	GetPostLikeCount(*pb.PostId) (*pb.Message, error)
	GetCommentLikeCount(*pb.Like) (*pb.PostListResponse, error)
	GetUsersWhichLikePost(*pb.PostId) (*pb.PostListResponse, error)
	GetUsersWhichLikeComment(*pb.CommentId) (*pb.PostListResponse, error)
}

type CommentStorage interface {
	CreateComment(*pb.CommentPost) (*pb.LikePost, error)
	UpdateComment(*pb.UpdateAComment) (*pb.Message, error)
	GetCommentByID(*pb.CommentId) (*pb.PostListResponse, error)
	GetCommentByUsername(*pb.UserPostId) (*pb.Message, error)
	ListComments(*pb.CommentList) (*pb.Message, error)
	DeleteComment(*pb.CommentId) (*pb.PostListResponse, error)
	GetCommentByPostID(*pb.PostId) (*pb.PostListResponse, error)
	GetAllUserComments(*pb.CommentId) (*pb.PostListResponse, error)
	GetMostlikeCommentPost(*pb.PostId) (*pb.PostListResponse, error)
}
