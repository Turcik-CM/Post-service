package postgres

import (
	"fmt"
	pb "post-servic/genproto/post"
	"post-servic/storage"

	"github.com/jmoiron/sqlx"
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
	// Check if user exists
	var userExists bool
	userCheckQuery := `SELECT EXISTS (SELECT 1 FROM users WHERE id = $1)`
	err := c.db.Get(&userExists, userCheckQuery, in.Username)
	if err != nil || !userExists {
		return nil, fmt.Errorf("user does not exist")
	}

	// Check if post exists
	var postExists bool
	postCheckQuery := `SELECT EXISTS (SELECT 1 FROM posts WHERE id = $1)`
	err = c.db.Get(&postExists, postCheckQuery, in.PostId)
	if err != nil || !postExists {
		return nil, fmt.Errorf("post does not exist")
	}

	// Create comment within a transaction
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO comments (username, post_id, content, created_at)
		VALUES ($1, $2, $3, now())
		RETURNING id, username, post_id, content, created_at, updated_at
	`

	var comment pb.CommentResponse
	err = tx.QueryRow(query, in.Username, in.PostId, in.Content).Scan(
		&comment.Id, &comment.Username, &comment.PostId, &comment.Content,
		&comment.CreatedAt, &comment.UpdatedAt,
	)
	if err != nil {
		tx.Rollback() // rollback if error occurs
		return nil, err
	}

	err = tx.Commit() // commit the transaction
	if err != nil {
		return nil, err
	}

	return &comment, nil
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
