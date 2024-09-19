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

func (l *LikeStorage) AddLikePost(in *pb.LikePost) (*pb.LikeResponse, error) {
	return nil, nil
}

func (l *LikeStorage) DeleteLikePost(in *pb.LikePost) (*pb.Message, error) {
	return nil, nil
}

func (l *LikeStorage) AddLikeComment(in *pb.LikePost) (*pb.LikeResponse, error) {
	return nil, nil
}

func (l *LikeStorage) DeleteLikeComment(in *pb.LikePost) (*pb.Message, error) {
	return nil, nil
}

func (l *LikeStorage) GetPostLikeCount(in *pb.PostId) (*pb.PostResponse, error) {
	return nil, nil
}

func (l *LikeStorage) GetCommentLikeCount(in *pb.PostId) (*pb.CommentResponse, error) {
	return nil, nil
}

func (l *LikeStorage) GetUsersWhichLikePost(in *pb.PostId) (*pb.Users, error) {
	return nil, nil
}

func (l *LikeStorage) GetUsersWhichLikeComment(in *pb.CommentId) (*pb.Users, error) {
	return nil, nil
}
