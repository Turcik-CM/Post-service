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

func (c *CommentStorage) CreateComment(in *pb.CommentPost) (*pb.LikePost, error) {
	return nil, nil
}

func (c *CommentStorage) UpdateComment(in *pb.UpdateAComment) (*pb.Message, error) {
	return nil, nil
}

func (c *CommentStorage) GetCommentByID(in *pb.CommentId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (c *CommentStorage) GetCommentByUsername(in *pb.UserPostId) (*pb.Message, error) {
	return nil, nil
}

func (c *CommentStorage) ListComments(in *pb.CommentList) (*pb.Message, error) {
	return nil, nil
}

func (c *CommentStorage) DeleteComment(in *pb.CommentId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (c *CommentStorage) GetCommentByPostID(in *pb.PostId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (c *CommentStorage) GetAllUserComments(in *pb.CommentId) (*pb.PostListResponse, error) {
	return nil, nil
}

func (c *CommentStorage) GetMostlikeCommentPost(in *pb.PostId) (*pb.PostListResponse, error) {
	return nil, nil
}
