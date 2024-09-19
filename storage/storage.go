package storage

import (
	pb "post-servic/genproto/post"
)

type PostStorage interface {
	CreatePost(in *pb.Post) (*pb.PostResponse, error)
	UpdatePost(in *pb.UpdateAPost) (*pb.PostResponse, error)
	GetPostByID(in *pb.PostId) (*pb.PostResponse, error)
	ListPosts(in *pb.PostList) (*pb.PostListResponse, error)
	DeletePost(in *pb.PostId) (*pb.Message, error)
	AddImageToPost(in *pb.ImageUrl) (*pb.PostResponse, error)
	RemoveImageFromPost(in *pb.ImageUrl) (*pb.PostListResponse, error)
	GetPostByCountry(in *pb.PostCountry) (*pb.PostListResponse, error)
}

type LikeStorage interface {
	AddLikePost(in *pb.LikePost) (*pb.LikeResponse, error)
	DeleteLikePost(in *pb.LikePost) (*pb.Message, error)
	AddLikeComment(in *pb.LikePost) (*pb.LikeResponse, error)
	DeleteLikeComment(in *pb.LikePost) (*pb.Message, error)
	GetPostLikeCount(in *pb.PostId) (*pb.PostResponse, error)
	GetCommentLikeCount(in *pb.PostId) (*pb.CommentResponse, error)
	GetUsersWhichLikePost(in *pb.PostId) (*pb.Users, error)
	GetUsersWhichLikeComment(in *pb.CommentId) (*pb.Users, error)
}

type CommentStorage interface {
	CreateComment(in *pb.CommentPost) (*pb.CommentResponse, error)
	UpdateComment(in *pb.UpdateAComment) (*pb.CommentResponse, error)
	GetCommentByID(in *pb.CommentId) (*pb.CommentResponse, error)
	GetCommentByUsername(in *pb.Username) (*pb.CommentResponse, error)
	ListComments(in *pb.CommentList) (*pb.CommentsR, error)
	DeleteComment(in *pb.CommentId) (*pb.Message, error)
	GetCommentByPostID(in *pb.PostId) (*pb.CommentsR, error)
	GetAllUserComments(in *pb.Username) (*pb.CommentsR, error)
	GetMostlikeCommentPost(in *pb.PostId) (*pb.CommentResponse, error)
}
