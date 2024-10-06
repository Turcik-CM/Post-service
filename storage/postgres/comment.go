package postgres

import (
	"context"
	"database/sql"
	"fmt"
	pb "post-servic/genproto/post"
	"post-servic/storage"
	"strings"
	"time"

	"github.com/google/uuid"
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
	id := uuid.New()
	createdAt := time.Now()
	updatedAt := createdAt

	query := `
		INSERT INTO comments (id, user_id, post_id, content, created_at, updated_at) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	if err := c.db.QueryRow(query, id, in.UserId, in.PostId, in.Content, createdAt, updatedAt).Scan(&id); err != nil {
		return nil, err
	}

	return &pb.CommentResponse{
		Id:        id.String(),
		UserId:    in.UserId,
		PostId:    in.PostId,
		Content:   in.Content,
		CreatedAt: createdAt.String(),
		UpdatedAt: updatedAt.String(),
	}, nil
}

func (c *CommentStorage) UpdateComment(in *pb.UpdateAComment) (*pb.CommentResponse, error) {
	query := `UPDATE comments SET `
	args := []interface{}{}
	argIndex := 1
	updateFields := []string{}

	if in.Content != "" {
		updateFields = append(updateFields, fmt.Sprintf("content = $%d", argIndex))
		args = append(args, in.Content)
		argIndex++
	}

	if len(updateFields) > 0 {
		query += fmt.Sprintf("%s, updated_at = $%d", strings.Join(updateFields, ", "), argIndex)
		args = append(args, time.Now())
		argIndex++
	} else {
		return nil, fmt.Errorf("hech qanday maydon yangilanmadi")
	}

	query += fmt.Sprintf(" WHERE id = $%d AND user_id = $%d RETURNING id, user_id, post_id, content, created_at, updated_at", argIndex, argIndex+1)
	args = append(args, in.Id, in.UserId)

	var res pb.CommentResponse
	err := c.db.QueryRowContext(context.Background(), query, args...).Scan(
		&res.Id, &res.UserId, &res.PostId, &res.Content, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *CommentStorage) GetCommentByID(in *pb.CommentId) (*pb.CommentResponse, error) {
	query := `SELECT id, user_id, post_id, content, created_at, updated_at FROM comments WHERE id = $1`

	var comment pb.CommentResponse
	if err := c.db.QueryRow(query, in.CommentId).Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
		return nil, err
	}

	return &comment, nil
}

func (c *CommentStorage) DeleteComment(in *pb.CommentId) (*pb.Message, error) {
	query := `DELETE FROM comments WHERE id = $1`

	if _, err := c.db.Exec(query, in.CommentId); err != nil {
		return nil, err
	}

	return &pb.Message{Massage: "Comment deleted successfully"}, nil
}

func (c *CommentStorage) GetCommentByUsername(in *pb.Username) (*pb.CommentResponse, error) {
	query := `SELECT id, user_id, post_id, content, created_at, updated_at FROM comments WHERE user_id = $1`

	var comment pb.CommentResponse
	if err := c.db.QueryRow(query, in.Username).Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
		return nil, err
	}

	return &comment, nil
}

func (c *CommentStorage) ListComments(in *pb.CommentList) (*pb.CommentsR, error) {
	query := `SELECT id, user_id, post_id, content, created_at, updated_at FROM comments WHERE 1=1`
	args := []interface{}{}
	argIndex := 1

	if in.PostId != "" {
		query += fmt.Sprintf(" AND post_id = $%d", argIndex)
		args = append(args, in.PostId)
		argIndex++
	}

	if in.Limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argIndex)
		args = append(args, in.Limit)
		argIndex++
	}

	if in.Offset >= 0 {
		query += fmt.Sprintf(" OFFSET $%d", argIndex)
		args = append(args, in.Offset)
		argIndex++
	}

	rows, err := c.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*pb.CommentResponse
	for rows.Next() {
		var comment pb.CommentResponse
		if err := rows.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return &pb.CommentsR{Comments: comments}, nil
}

func (c *CommentStorage) GetCommentByPostID(in *pb.PostId) (*pb.CommentsR, error) {
	query := `SELECT id, user_id, post_id, content, created_at, updated_at FROM comments WHERE post_id = $1`
	rows, err := c.db.Query(query, in.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*pb.CommentResponse
	for rows.Next() {
		var comment pb.CommentResponse
		if err := rows.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return &pb.CommentsR{Comments: comments}, nil
}

func (c *CommentStorage) GetAllUserComments(in *pb.Username) (*pb.CommentsR, error) {
	query := `SELECT id, user_id, post_id, content, created_at, updated_at FROM comments WHERE user_id = $1`
	rows, err := c.db.Query(query, in.Username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*pb.CommentResponse
	for rows.Next() {
		var comment pb.CommentResponse
		if err := rows.Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}

	return &pb.CommentsR{Comments: comments}, nil
}

func (c *CommentStorage) GetMostlikeCommentPost(in *pb.PostId) (*pb.CommentResponse, error) {
	query := `
		SELECT c.id, c.user_id, c.post_id, c.content, c.created_at, c.updated_at as like_count
		FROM comments c
		LEFT JOIN likes l ON c.user_id = l.user_id -- to'g'ri bog'lanish comment_id bilan
		WHERE c.post_id = $1
		GROUP BY c.id, c.user_id, c.post_id, c.content, c.created_at, c.updated_at
		ORDER BY like_count DESC
		LIMIT 1`

	var comment pb.CommentResponse
	err := c.db.QueryRow(query, in.Id).Scan(
		&comment.Id,
		&comment.UserId,
		&comment.PostId,
		&comment.Content,
		&comment.CreatedAt,
		&comment.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &comment, nil
}
