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
		Location:    "1111",
		Title:       "1111",
		Hashtag:     "1111",
		Content:     "1111",
		ImageUrl:    "1111",
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
		Id:      "2a05b627-d55b-42f3-b497-fed2b014fd1b",
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
		Id: "2a05b627-d55b-42f3-b497-fed2b014fd1b",
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
		Limit: 5,
		//Country: "Uzbekistan",
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
		Id: "a2083af4-ee47-47f9-b4ed-096c3b75934a",
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
		PostId: "2a05b627-d55b-42f3-b497-fed2b014fd1b",
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
		PostId: "2a05b627-d55b-42f3-b497-fed2b014fd1b",
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
