package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "post-servic/genproto/post"
	"post-servic/storage"
)

type CommentStorage struct {
	db *sqlx.DB
}

func NewCommentStorage(db *sqlx.DB) storage.CommentStorage {
	return &CommentStorage{
		db: db,
	}
}

func (c *CommentStorage) CreateComment(*pb.CommentPost) (*pb.LikePost, error) {
	return nil, nil
}

func (c *CommentStorage) UpdateComment(*pb.UpdateAComment) (*pb.Message, error) {
	return nil, nil
}

func (c *CommentStorage) GetCommentByID(*pb.CommentId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (c *CommentStorage) GetCommentByUsername(*pb.UserPostId) (*pb.Message, error) {
	return nil, nil
}

func (c *CommentStorage) ListComments(*pb.CommentList) (*pb.Message, error) {
	return nil, nil
}

func (c *CommentStorage) DeleteComment(*pb.CommentId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (c *CommentStorage) GetCommentByPostID(*pb.PostId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (c *CommentStorage) GetAllUserComments(*pb.CommentId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (c *CommentStorage) GetMostlikeCommentPost(*pb.PostId) (*pb.PostListResponse, error) {
	return nil, nil
}
