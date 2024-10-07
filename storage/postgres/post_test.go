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
		Country:     "Uzbekistan",
		Location:    "11111",
		Title:       "11111",
		Hashtag:     "11111",
		Content:     "11111",
		ImageUrl:    "11111",
		Description: "nimadur",
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
		Id:      "f380a915-4bd8-434e-a699-7b82bbfb3906",
		Country: "Uzbekistan",
		Title:   "dodi2",
		Content: "Uzbekistan",
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
		Id: "f380a915-4bd8-434e-a699-7b82bbfb3906",
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
		//Limit: 5,
		Country: "zbe",
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
		Id: "f380a915-4bd8-434e-a699-7b82bbfb3906",
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
		PostId: "f380a915-4bd8-434e-a699-7b82bbfb3906",
		Url:    "dodi",
	}

	post := NewPostStorage(db)
	req, err := post.AddImageToPost(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req, "ssssssssss")
}

func TestRemoveImageFromPost(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.ImageUrl{
		PostId: "f380a915-4bd8-434e-a699-7b82bbfb3906",
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
		Country: "Uzbekistan",
	}

	post := NewPostStorage(db)

	req, err := post.GetPostByCountry(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}
