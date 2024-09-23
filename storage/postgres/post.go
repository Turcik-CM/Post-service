package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	pb "post-servic/genproto/post"
	"post-servic/storage"
	"strings"
	"time"
)

type PostStorage struct {
	db *sqlx.DB
}

func NewPostStorage(db *sqlx.DB) storage.PostStorage {
	return &PostStorage{
		db: db,
	}
}

func (p *PostStorage) CreatePost(in *pb.Post) (*pb.PostResponse, error) {
	query1 := `SELECT name FROM hashtag WHERE name = $1`

	err := p.db.QueryRowContext(context.Background(), query1, in.Hashtag).Scan(&in.Hashtag)
	if err == sql.ErrNoRows {
		query2 := `INSERT INTO hashtag (name, description) VALUES ($1, $2)`
		_, err = p.db.ExecContext(context.Background(), query2, in.Hashtag, "Hashtag description")
		if err != nil {
			return nil, fmt.Errorf("failed to insert new hashtag: %v", err)
		}
	} else if err != nil {
		return nil, fmt.Errorf("failed to query hashtag: %v", err)
	}
	fmt.Println("dodi")

	query4 := `SELECT country FROM countries WHERE country = $1`

	err = p.db.QueryRowContext(context.Background(), query4, in.Country).Scan(&in.Country)
	if err == sql.ErrNoRows {
		query3 := `INSERT INTO countries (nationality, country, flag) VALUES ($1, $2, $3)`
		_, err = p.db.ExecContext(context.Background(), query3, "uz", in.Country, "dodi")
		if err != nil {
			return nil, fmt.Errorf("failed to insert new nationality: %v", err)
		}
	}
	fmt.Println("dodi")
	query := `INSERT INTO posts (user_id, country, location, title, hashtag, content, description) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7) 
	          RETURNING id, user_id, country, location, title, hashtag, content, created_at, description`

	var res pb.PostResponse
	err = p.db.QueryRowContext(context.Background(), query,
		in.UserId, in.Country, in.Location, in.Title, in.Hashtag, in.Content, in.Description).Scan(
		&res.Id,
		&res.UserId,
		&res.Country,
		&res.Location,
		&res.Title,
		&res.Hashtag,
		&res.Content,
		&res.CreatedAt,
		&res.Description)
	fmt.Println("dodi")
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %v", err)
	}

	return &res, nil
}

func (p *PostStorage) UpdatePost(in *pb.UpdateAPost) (*pb.PostResponse, error) {
	query := `UPDATE posts SET `
	args := []interface{}{}
	argIndex := 1
	updateFields := []string{}
	if in.Description != "" {
		updateFields = append(updateFields, "description")
		args = append(args, in.Description)
		argIndex++
	}

	if in.Country != "" {
		updateFields = append(updateFields, fmt.Sprintf("country = $%d", argIndex))
		args = append(args, in.Country)
		argIndex++
	}

	if in.Location != "" {
		updateFields = append(updateFields, fmt.Sprintf("location = $%d", argIndex))
		args = append(args, in.Location)
		argIndex++
	}

	if in.Title != "" {
		updateFields = append(updateFields, fmt.Sprintf("title = $%d", argIndex))
		args = append(args, in.Title)
		argIndex++
	}

	if in.Hashtag != "" {
		updateFields = append(updateFields, fmt.Sprintf("hashtag = $%d", argIndex))
		args = append(args, in.Hashtag)
		argIndex++
	}

	if in.Content != "" {
		updateFields = append(updateFields, fmt.Sprintf("content = $%d", argIndex))
		args = append(args, in.Content)
		argIndex++
	}

	if in.ImageUrl != "" {
		updateFields = append(updateFields, fmt.Sprintf("image_url = $%d", argIndex))
		args = append(args, in.ImageUrl)
		argIndex++
	}

	if len(updateFields) > 0 {
		query += fmt.Sprintf("%s, updated_at = $%d", strings.Join(updateFields, ", "), argIndex)
		args = append(args, time.Now())
		argIndex++
	} else {
		return nil, fmt.Errorf("hech qanday maydon yangilanmadi")
	}

	query += fmt.Sprintf(" WHERE id = $%d RETURNING id, user_id, country, location, title, hashtag, content, image_url, description,  created_at, updated_at", argIndex)
	args = append(args, in.Id)

	fmt.Println("dodi")
	var res pb.PostResponse
	err := p.db.QueryRowContext(context.Background(), query, args...).Scan(
		&res.Id, &res.UserId, &res.Country, &res.Location, &res.Title, &res.Hashtag, &res.Content, &res.ImageUrl, &res.Description, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		return nil, err
	}
	fmt.Println("dodi")

	return &res, nil
}

func (p *PostStorage) GetPostByID(in *pb.PostId) (*pb.PostResponse, error) {
	query := `SELECT id, user_id, country, location, title, hashtag, content, image_url, description, created_at, updated_at 
	          FROM posts WHERE id = $1`

	var res pb.PostResponse
	err := p.db.QueryRowContext(context.Background(), query, in.Id).Scan(
		&res.Id, &res.UserId, &res.Country, &res.Location, &res.Title, &res.Hashtag, &res.Content, &res.ImageUrl, &res.Description, &res.CreatedAt, &res.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &res, nil
}

func (p *PostStorage) ListPosts(in *pb.PostList) (*pb.PostListResponse, error) {
	query := "SELECT id, user_id, country, location, title, hashtag, content, image_url, description, created_at, updated_at FROM posts WHERE deleted_at = 0"
	args := []interface{}{}

	if in.Hashtag != "" {
		query += " AND hashtag = $1"
		args = append(args, in.Hashtag)
	}

	if in.Country != "" {
		query += " AND country = $" + fmt.Sprintf("%d", len(args)+1)
		args = append(args, in.Country)
	}

	if in.Limit > 0 {
		query += " LIMIT $" + fmt.Sprintf("%d", len(args)+1)
		args = append(args, in.Limit)
	}
	if in.Offset >= 0 {
		query += " OFFSET $" + fmt.Sprintf("%d", len(args)+1)
		args = append(args, in.Offset)
	}

	rows, err := p.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*pb.PostResponse
	for rows.Next() {
		var post pb.PostResponse
		if err := rows.Scan(&post.Id, &post.UserId, &post.Country, &post.Location, &post.Title,
			&post.Hashtag, &post.Content, &post.ImageUrl, &post.Description, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	return &pb.PostListResponse{
		Post: posts,
	}, nil
}

func (p *PostStorage) DeletePost(in *pb.PostId) (*pb.Message, error) {
	query := `update posts set deleted_at = date_part('epoch', current_timestamp)::INT
                  where id = $1 and deleted_at = 0`

	_, err := p.db.ExecContext(context.Background(), query, in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Message{
		Massage: "Post muvaffaqiyatli o'chirildi (soft delete)",
	}, nil
}

func (p *PostStorage) AddImageToPost(in *pb.ImageUrl) (*pb.PostResponse, error) {
	query := `UPDATE posts SET image_url = $1, updated_at = $2 WHERE id = $3`

	fmt.Println("dodi")

	_, err := p.db.ExecContext(context.Background(), query, in.Url, time.Now(), in.PostId)
	if err != nil {
		return nil, err
	}

	inn := pb.PostId{
		Id: in.PostId,
	}

	res, err := p.GetPostByID(&inn)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (p *PostStorage) RemoveImageFromPost(in *pb.ImageUrl) (*pb.Message, error) {
	query := `UPDATE posts SET image_url = 'no image' WHERE id = $1 RETURNING id`

	var postId string
	err := p.db.QueryRowContext(context.Background(), query, in.PostId).Scan(&postId)
	if err != nil {
		return nil, err
	}

	return &pb.Message{
		Massage: fmt.Sprintf("Post %s rasm muvaffaqiyatli o'chirildi", postId),
	}, nil
}

func (p *PostStorage) GetPostByCountry(in *pb.PostCountry) (*pb.PostListResponse, error) {
	query := `SELECT id, user_id, country, location, title, hashtag, content, image_url, description, created_at, updated_at 
	          FROM posts WHERE country = $1`

	rows, err := p.db.QueryContext(context.Background(), query, in.Country)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*pb.PostResponse
	for rows.Next() {
		var post pb.PostResponse
		if err := rows.Scan(&post.Id, &post.UserId, &post.Country, &post.Location, &post.Title,
			&post.Hashtag, &post.Content, &post.ImageUrl, &post.Description, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pb.PostListResponse{
		Post: posts,
	}, nil
}
