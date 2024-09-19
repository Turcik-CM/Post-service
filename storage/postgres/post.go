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

func (p *PostStorage) CreatePost(in *pb.Post) (*pb.PostResponse, error) {
	return &pb.PostResponse{}, nil
}
func (p *PostStorage) UpdatePost(in *pb.UpdateAPost) (*pb.PostResponse, error) {
	return &pb.PostResponse{}, nil
}
func (p *PostStorage) GetPostByID(in *pb.PostId) (*pb.PostResponse, error) {
	return &pb.PostResponse{}, nil
}
func (p *PostStorage) ListPosts(in *pb.PostList) (*pb.PostListResponse, error) {
	return &pb.PostListResponse{}, nil
}
func (p *PostStorage) DeletePost(in *pb.PostId) (*pb.Message, error) {
	return &pb.Message{}, nil
}
func (p *PostStorage) AddImageToPost(in *pb.ImageUrl) (*pb.PostResponse, error) {
	return &pb.PostResponse{}, nil
}
func (p *PostStorage) RemoveImageFromPost(in *pb.ImageUrl) (*pb.PostListResponse, error) {
	return &pb.PostListResponse{}, nil
}
func (p *PostStorage) GetPostByCountry(in *pb.PostCountry) (*pb.PostListResponse, error) {
	return &pb.PostListResponse{}, nil
}
