package postgres

import (
	"fmt"
	"github.com/google/uuid"
	pb "post-servic/genproto/post"
	"post-servic/pkg/config"
	"testing"
)

func TestAddLikePost(t *testing.T) {
	cfg := config.Load()

	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}

	res := pb.LikePost{
		UserId: uuid.New().String(),
		PostId: "a2083af4-ee47-47f9-b4ed-096c3b75934a",
	}

	like := NewLikeStorage(db)

	req, err := like.AddLikePost(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestDeleteLikePost(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.LikePost{
		UserId: "08d32f24-fcc7-4238-aa75-6f7be3787e5d",
		PostId: "0a167e10-a32c-4df7-b88e-03c0e00301a5",
	}

	like := NewLikeStorage(db)
	req, err := like.DeleteLikePost(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestAddLikeComment(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}

	res := pb.LikePost{
		UserId: "2b1b7f64-aaf3-4edc-855b-59482d5ed687",
		PostId: "a2083af4-ee47-47f9-b4ed-096c3b75934a",
	}

	like := NewLikeStorage(db)
	req, err := like.AddLikePost(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestDeleteLikeComment(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.LikePost{
		UserId: "08d32f24-fcc7-4238-aa75-6f7be3787e5d",
		PostId: "0a167e10-a32c-4df7-b88e-03c0e00301a5",
	}
	like := NewLikeStorage(db)
	req, err := like.DeleteLikePost(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestGetPostLikeCount(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.PostId{
		Id: "0a167e10-a32c-4df7-b88e-03c0e00301a5",
	}
	like := NewLikeStorage(db)
	req, err := like.GetPostLikeCount(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestGetMostLikedComment(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.PostId{
		Id: "0a167e10-a32c-4df7-b88e-03c0e00301a5",
	}

	like := NewLikeStorage(db)
	req, err := like.GetMostLikedComment(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}
