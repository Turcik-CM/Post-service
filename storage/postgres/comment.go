package postgres

import (
	pb "post-servic/genproto/post"
	"post-servic/storage"
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

// CreateComment - yangi izoh qo'shish
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

// UpdateComment - izohni yangilash
func (c *CommentStorage) UpdateComment(in *pb.UpdateAComment) (*pb.CommentResponse, error) {
	updatedAt := time.Now()

	query := `
		UPDATE comments 
		SET content = $1, updated_at = $2 
		WHERE id = $3 RETURNING id, user_id, post_id, content, created_at, updated_at`

	var comment pb.CommentResponse
	if err := c.db.QueryRow(query, in.Content, updatedAt, in.Id).Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
		return nil, err
	}

	return &comment, nil
}

// GetCommentByID - ID bo'yicha izohni olish
func (c *CommentStorage) GetCommentByID(in *pb.CommentId) (*pb.CommentResponse, error) {
	query := `SELECT id, user_id, post_id, content, created_at, updated_at FROM comments WHERE id = $1`

	var comment pb.CommentResponse
	if err := c.db.QueryRow(query, in.CommentId).Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
		return nil, err
	}

	return &comment, nil
}

// DeleteComment - izohni o'chirish
func (c *CommentStorage) DeleteComment(in *pb.CommentId) (*pb.Message, error) {
	query := `DELETE FROM comments WHERE id = $1`

	if _, err := c.db.Exec(query, in.CommentId); err != nil {
		return nil, err
	}

	return &pb.Message{Massage: "Comment deleted successfully"}, nil
}

// GetCommentByUsername - foydalanuvchi nomi bo'yicha izohlarni olish
func (c *CommentStorage) GetCommentByUsername(in *pb.Username) (*pb.CommentResponse, error) {
	query := `SELECT id, user_id, post_id, content, created_at, updated_at FROM comments WHERE user_id = $1`

	var comment pb.CommentResponse
	if err := c.db.QueryRow(query, in.Username).Scan(&comment.Id, &comment.UserId, &comment.PostId, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
		return nil, err
	}

	return &comment, nil
}

// ListComments - izohlar ro'yxatini olish
func (c *CommentStorage) ListComments(in *pb.CommentList) (*pb.CommentsR, error) {
	query := `SELECT id, user_id, post_id, content, created_at, updated_at FROM comments WHERE post_id = $1 LIMIT $2 OFFSET $3`
	rows, err := c.db.Query(query, in.PostId, in.Limit, in.Offset)
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

// GetCommentByPostID - post ID bo'yicha izohlarni olish
func (c *CommentStorage) GetCommentByPostID(in *pb.PostId) (*pb.CommentsR, error) {
	return c.ListComments(&pb.CommentList{PostId: in.Id})
}

// GetAllUserComments - foydalanuvchining barcha izohlarini olish
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

// GetMostlikeCommentPost - eng ko'p yoqtirilgan izohni olish
func (c *CommentStorage) GetMostlikeCommentPost(in *pb.PostId) (*pb.CommentResponse, error) {
	// Eng ko'p yoqtirilgan izohni olish uchun so'rov
	query := `
		SELECT c.id, c.user_id, c.post_id, c.content, c.created_at, c.updated_at, COUNT(l.user_id) as like_count
		FROM comments c
		LEFT JOIN likes l ON c.id = l.comment_id
		WHERE c.post_id = $1
		GROUP BY c.id
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
			return nil, nil // Agar izoh topilmasa, nil qaytarish
		}
		return nil, err // Boshqa xatoliklar
	}

	return &comment, nil
}
