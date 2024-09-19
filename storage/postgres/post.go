package postgres

import (
	"github.com/jmoiron/sqlx"
	pb "post-servic/genproto/post"
	"post-servic/storage"
)

type PostStorage struct {
	db *sqlx.DB
}

func NewPostStorage(db *sqlx.DB) storage.PostStorage {
	return &PostStorage{
		db: db,
	}
}

func (p *PostStorage) CreatePost(*pb.Post) (*pb.PostResponse, error) {
	return &pb.PostResponse{}, nil
}
func (p *PostStorage) UpdatePost(*pb.UpdateAPost) (*pb.PostResponse, error) {
	return &pb.PostResponse{}, nil
}
func (p *PostStorage) GetPostByID(*pb.PostId) (*pb.PostResponse, error) {
	return &pb.PostResponse{}, nil
}
func (p *PostStorage) ListPosts(*pb.PostList) (*pb.PostListResponse, error) {
	return &pb.PostListResponse{}, nil
}
func (p *PostStorage) DeletePost(*pb.PostId) (*pb.Message, error) {
	return &pb.Message{}, nil
}
func (p *PostStorage) RemoveImageFromPost(*pb.Post) (*pb.PostListResponse, error) {
	return &pb.PostListResponse{}, nil
}
func (p *PostStorage) GetPostByCountry(*pb.PostCountry) (*pb.PostListResponse, error) {
	return &pb.PostListResponse{}, nil
}
