package postgres

import (
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	pb "post-servic/genproto/post"
	"post-servic/pkg/config"
	"testing"
)

func TestCreatePost(t *testing.T) {
	cfg := config.Load()

	db, err := ConnectPostgres(cfg)

	if err != nil {
		t.Fatal(err)
	}

	res := pb.Post{
		UserId:      uuid.New().String(),
		Nationality: "Toshkent",
		Location:    "dodi1",
		Title:       "dodi1",
		Hashtag:     "dodi",
		Content:     "dodi1",
	}

	post := NewPostStorage(db)

	req, err := post.CreatePost(&res)
	fmt.Println(req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req, "dodi")
}

func TestUpdatePost(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.UpdateAPost{
		Id:          "6622e177-582a-4d7a-bd78-1743f611c7ae",
		Nationality: "Toshkent",
		Title:       "dodi2",
		Content:     "Uzbekistan",
	}

	post := NewPostStorage(db)
	req, err := post.UpdatePost(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestGetPostByID(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}

	res := pb.PostId{
		Id: "6622e177-582a-4d7a-bd78-1743f611c7ae",
	}

	post := NewPostStorage(db)

	req, err := post.GetPostByID(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestListPosts(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.PostList{
		Limit:   5,
		Country: "Toshkent",
	}

	post := NewPostStorage(db)
	req, err := post.ListPosts(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req, "dodi")
}

func TestDeletePost(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}

	res := pb.PostId{
		Id: "67b99246-ed95-49dd-8cc6-8f0009bdaf0b",
	}

	post := NewPostStorage(db)
	req, err := post.DeletePost(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestAddImageToPost(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}

	res := pb.ImageUrl{
		PostId: "eca9c592-2468-4316-8221-26d1ac9411ab",
		Url:    "dodi",
	}

	post := NewPostStorage(db)
	req, err := post.AddImageToPost(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestRemoveImageFromPost(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.ImageUrl{
		PostId: "eca9c592-2468-4316-8221-26d1ac9411ab",
	}
	post := NewPostStorage(db)
	req, err := post.RemoveImageFromPost(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestGetPostByCountry(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.PostCountry{
		Nationality: "Toshkent",
	}

	post := NewPostStorage(db)

	req, err := post.GetPostByCountry(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}
