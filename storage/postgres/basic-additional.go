package postgres

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
	pb "post-servic/genproto/post"
	"post-servic/storage"
)

type BasicAdditional struct {
	db *sqlx.DB
}

func NewBasicAdditional(db *sqlx.DB) storage.BasicAdditional {
	return &BasicAdditional{db}
}

func (b *BasicAdditional) GetUserRecommendation(in *pb.Username) (*pb.PostListResponse, error) {
	var res *pb.PostListResponse

	// getting last 10 posts which user liked
	var postID []string
	err := b.db.Select(&postID, "SELECT post_id from likes where user_id=$1 order by created_at desc limit 10", in.Username)
	if errors.Is(err, sql.ErrNoRows) {
		err = b.db.Select(&res, `SELECT id, user_id, nationality, location, title, hashtag, content,
       image_url, created_at, updated_at FROM posts order by created_at desc limit 20`)
		return res, err
	}
	if err != nil {
		return nil, err
	}

	// getting user ID from Post table
	var userId []string
	q := "SELECT user_id from posts where id in (?)"

	query, args, err := sqlx.In(q, postID)
	if err != nil {
		return nil, err
	}

	log.Println(query, args)

	query = b.db.Rebind(query)
	log.Println(query)

	err = b.db.Select(&userId, query, args...)
	if err != nil {
		return nil, err
	}

	// getting  nationality from Post table which user liked
	var nationality []string
	q = "SELECT nationality from posts where id in (?)"

	query, args, err = sqlx.In(q, postID)
	if err != nil {
		return nil, err
	}

	log.Println(query, args)

	query = b.db.Rebind(query)
	log.Println(query)

	err = b.db.Select(&nationality, query, args...)
	if err != nil {
		return nil, err
	}

	// getting  hashtag from Post table which user liked
	var hashtag []string
	q = "SELECT hashtag from posts where id in (?)"

	query, args, err = sqlx.In(q, postID)
	if err != nil {
		return nil, err
	}

	log.Println(query, args)

	query = b.db.Rebind(query)
	log.Println(query)

	err = b.db.Select(&hashtag, query, args...)
	if err != nil {
		return nil, err
	}

	// getting user recommended posts
	q = `SELECT id, user_id, nationality, location, title, hashtag, content, image_url, created_at, updated_at 
         FROM posts 
         WHERE nationality IN (?) AND user_id IN (?) AND hashtag IN (?) 
         order by created_at desc
         limit 20`

	query, args, err = sqlx.In(q, nationality, userId, hashtag)
	if err != nil {
		return nil, err
	}

	query = b.db.Rebind(query)
	err = b.db.Select(&res, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *BasicAdditional) GetPostsByUsername(in *pb.Username) (*pb.PostListResponse, error) {
	var res *pb.PostListResponse

	err := b.db.Select(&res, `SELECT id, user_id, nationality, location, title, hashtag, content, image_url,
       created_at, updated_at FROM posts WHERE user_id=$1 order by created_at desc`, in.Username)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (b *BasicAdditional) GetTrendsPost(in *pb.Void) (*pb.PostListResponse, error) {
	return nil, nil
}

func (b *BasicAdditional) SearchPost(in *pb.Search) (*pb.PostListResponse, error) {
	var res *pb.PostListResponse

	if in.Action == "" {
		return res, nil
	}

	in.Action = "%" + in.Action + "%"

	query := `SELECT id, user_id, nationality, location, title, hashtag, content, image_url,
       created_at, updated_at FROM posts WHERE title ilike $1 or description ilike $2 or hashtag $3`

	err := b.db.Select(&res, query, in.Action, in.Action, in.Action)
	if err != nil {
		return nil, err
	}

	return res, nil
}
