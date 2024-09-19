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

func (c *CommentStorage) CreateComment(in *pb.CommentPost) (*pb.CommentResponse, error) {
	return nil, nil
}

func (c *CommentStorage) UpdateComment(in *pb.UpdateAComment) (*pb.CommentResponse, error) {
	return nil, nil
}

func (c *CommentStorage) GetCommentByID(in *pb.CommentId) (*pb.CommentResponse, error) {
	return nil, nil
}

func (c *CommentStorage) GetCommentByUsername(in *pb.Username) (*pb.CommentResponse, error) {
	return nil, nil
}

func (c *CommentStorage) ListComments(in *pb.CommentList) (*pb.CommentsR, error) {
	return nil, nil
}

func (c *CommentStorage) DeleteComment(in *pb.CommentId) (*pb.Message, error) {
	return nil, nil
}

func (c *CommentStorage) GetCommentByPostID(in *pb.PostId) (*pb.CommentsR, error) {
	return nil, nil
}

func (c *CommentStorage) GetAllUserComments(in *pb.Username) (*pb.CommentsR, error) {
	return nil, nil
}

func (c *CommentStorage) GetMostlikeCommentPost(in *pb.PostId) (*pb.CommentResponse, error) {
	return nil, nil
}
