package service

import (
	"log/slog"
	pb "post-servic/genproto/post"
	"post-servic/storage"
)

type PostService struct {
	post    storage.PostStorage
	like    storage.LikeStorage
	comment storage.CommentStorage
	logger  slog.Logger
	pb.UnimplementedPostServiceServer
}

func NewPostService(ps storage.PostStorage, ls storage.LikeStorage, cs storage.CommentStorage, log *slog.Logger) *PostService {
	return &PostService{
		post:    ps,
		like:    ls,
		comment: cs,
		logger:  *log,
	}
}
