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

func (l *LikeStorage) AddLikePost(in *pb.LikePost) (*pb.LikePost, error) {
	return nil, nil
}

func (l *LikeStorage) DeleteLikePost(in *pb.LikeId) (*pb.Message, error) {
	return nil, nil
}

func (l *LikeStorage) AddLikeComment(in *pb.LikeId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (l *LikeStorage) DeleteLikeComment(in *pb.UserId) (*pb.Message, error) {
	return nil, nil
}

func (l *LikeStorage) GetPostLikeCount(in *pb.PostId) (*pb.Message, error) {
	return nil, nil
}

func (l *LikeStorage) GetCommentLikeCount(in *pb.Like) (*pb.PostListResponse, error) {
	return nil, nil
}

func (l *LikeStorage) GetUsersWhichLikePost(in *pb.PostId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (l *LikeStorage) GetUsersWhichLikeComment(in *pb.CommentId) (*pb.PostListResponse, error) {
	return nil, nil
}
