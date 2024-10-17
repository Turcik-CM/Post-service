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
		PostId: "f380a915-4bd8-434e-a699-7b82bbfb3906",
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
		UserId: "f1225c22-48c6-4dc3-bfe6-5c6d48f03a2a",
		PostId: "f380a915-4bd8-434e-a699-7b82bbfb3906",
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
		UserId: "f1225c22-48c6-4dc3-bfe6-5c6d48f03a2a",
		PostId: "f380a915-4bd8-434e-a699-7b82bbfb3906",
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
	res := pb.LikeCommit{
		UserId:   "c732f495-bc83-492a-bdae-e98fc430345c",
		CommitId: "32a94376-4946-42cf-8e34-9b7f288539b6",
	}
	like := NewLikeStorage(db)
	req, err := like.DeleteLikeComment(&res)
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
		Id: "32a94376-4946-42cf-8e34-9b7f288539b6",
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
