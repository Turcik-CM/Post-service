package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "post-servic/genproto/post"
	"post-servic/storage"
)

type LikeStorage struct {
	db *sqlx.DB
}

func NewLikeStorage(db *sqlx.DB) storage.LikeStorage {
	return &LikeStorage{
		db: db,
	}
}

func (l *LikeStorage) AddLikePost(*pb.LikePost) (*pb.LikePost, error) {
	return nil, nil
}

func (l *LikeStorage) DeleteLikePost(*pb.LikeId) (*pb.Message, error) {
	return nil, nil
}

func (l *LikeStorage) AddLikeComment(*pb.LikeId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (l *LikeStorage) DeleteLikeComment(*pb.UserId) (*pb.Message, error) {
	return nil, nil
}

func (l *LikeStorage) GetPostLikeCount(*pb.PostId) (*pb.Message, error) {
	return nil, nil
}

func (l *LikeStorage) GetCommentLikeCount(*pb.Like) (*pb.PostListResponse, error) {
	return nil, nil
}

func (l *LikeStorage) GetUsersWhichLikePost(*pb.PostId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (l *LikeStorage) GetUsersWhichLikeComment(*pb.CommentId) (*pb.PostListResponse, error) {
	return nil, nil
}
