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
	RemoveImageFromPost(in *pb.ImageUrl) (*pb.Message, error)
	GetPostByCountry(in *pb.PostCountry) (*pb.PostListResponse, error)
}

type LikeStorage interface {
	AddLikePost(in *pb.LikePost) (*pb.LikeResponse, error)
	DeleteLikePost(in *pb.LikePost) (*pb.Message, error)
	AddLikeComment(in *pb.LikeCommit) (*pb.LikeComResponse, error)
	DeleteLikeComment(in *pb.LikeCommit) (*pb.Message, error)
	GetPostLikeCount(in *pb.PostId) (*pb.LikeCount, error)
	GetMostLikedComment(in *pb.PostId) (*pb.LikeCount, error)
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

type ChatStorage interface {
	StartMessaging(in *pb.CreateChat) (*pb.ChatResponse, error)
	SendMessage(in *pb.CreateMassage) (*pb.MassageResponse, error)
	GetChatMessages(in *pb.List) (*pb.MassageResponseList, error)
	MessageMarcTrue(in *pb.MassageTrue) (*pb.Message, error)
	GetUserChats(in *pb.Username) (*pb.ChatResponseList, error)
	GetUnreadMessages(in *pb.ChatId) (*pb.MassageResponseList, error)
	UpdateMessage(in *pb.UpdateMs) (*pb.MassageResponse, error)
	GetTodayMessages(in *pb.ChatId) (*pb.MassageResponseList, error)
	DeleteMessage(in *pb.MassageId) (*pb.Message, error)
	DeleteChat(in *pb.ChatId) (*pb.Message, error)
}

type BasicAdditional interface {
	GetUserRecommendation(in *pb.Username) (*pb.PostListResponse, error)
	GetPostsByUsername(in *pb.Username) (*pb.PostListResponse, error)
	GetTrendsPost(in *pb.Void) (*pb.PostListResponse, error)
	SearchPost(in *pb.Search) (*pb.PostListResponse, error)
}
