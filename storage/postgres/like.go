package postgres

import (
	"context"
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
	h := NewPostStorage(l.db)

	o := pb.PostId{
		Id: in.PostId,
	}

	_, err := h.GetPostByID(&o)
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO likes (user_id, post_id, created_at) 
	          VALUES ($1, $2, NOW())
	          RETURNING user_id, post_id`

	var res pb.LikeResponse
	err = l.db.QueryRowContext(context.Background(), query, in.UserId, in.PostId).Scan(
		&res.UserId, &res.PostId)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (l *LikeStorage) DeleteLikePost(in *pb.LikePost) (*pb.Message, error) {
	h := NewPostStorage(l.db)

	o := pb.PostId{
		Id: in.PostId,
	}

	_, err := h.GetPostByID(&o)
	if err != nil {
		return nil, err
	}

	query := `DELETE FROM likes WHERE user_id = $1 AND post_id = $2`

	_, err = l.db.ExecContext(context.Background(), query, in.UserId, in.PostId)
	if err != nil {
		return nil, err
	}

	return &pb.Message{
		Massage: "Like muvaffaqiyatli o'chirildi.",
	}, nil
}

func (l *LikeStorage) AddLikeComment(in *pb.LikeCommit) (*pb.LikeComResponse, error) {
	h := NewCommentStorage(l.db)

	o := pb.CommentId{
		CommentId: in.CommitId,
	}

	_, err := h.GetCommentByID(&o)
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO comment_like (user_id, comment_id, created_at) 
	          VALUES ($1, $2, NOW()) 
	          RETURNING user_id, comment_id`

	var res pb.LikeComResponse
	err = l.db.QueryRowContext(context.Background(), query, in.UserId, in.CommitId).Scan(
		&res.UserId, &res.CommitId)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (l *LikeStorage) DeleteLikeComment(in *pb.LikeCommit) (*pb.Message, error) {
	h := NewCommentStorage(l.db)

	o := pb.CommentId{
		CommentId: in.CommitId,
	}

	_, err := h.GetCommentByID(&o)
	if err != nil {
		return nil, err
	}

	query := `DELETE FROM comment_like WHERE user_id = $1 AND comment_id = $2`

	_, err = l.db.ExecContext(context.Background(), query, in.UserId, in.CommitId)
	if err != nil {
		return nil, err
	}

	return &pb.Message{
		Massage: "Like muvaffaqiyatli o'chirildi.",
	}, nil
}

func (l *LikeStorage) GetPostLikeCount(in *pb.PostId) (*pb.LikeCount, error) {
	query := `SELECT COUNT(*) FROM likes WHERE post_id = $1`

	var likeCount string
	err := l.db.QueryRowContext(context.Background(), query, in.Id).Scan(&likeCount)
	if err != nil {
		return nil, err
	}

	return &pb.LikeCount{
		Id:    in.Id,
		Count: likeCount,
	}, nil
}

func (l *LikeStorage) GetMostLikedComment(in *pb.PostId) (*pb.LikeCount, error) {
	query := `SELECT COUNT(*) FROM likes WHERE post_id = $1`

	var likeCount string
	err := l.db.QueryRowContext(context.Background(), query, in.Id).Scan(&likeCount)
	if err != nil {
		return nil, err
	}

	return &pb.LikeCount{
		Id:    in.Id,
		Count: likeCount,
	}, nil
}

func (l *LikeStorage) GetUsersWhichLikePost(in *pb.PostId) (*pb.Users, error) {
	return nil, nil
}

func (l *LikeStorage) GetUsersWhichLikeComment(in *pb.CommentId) (*pb.Users, error) {
	return nil, nil
}
