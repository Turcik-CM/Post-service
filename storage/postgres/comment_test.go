package postgres

import (
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	pb "post-servic/genproto/post"
	"post-servic/pkg/config"
	"testing"
	"time"
)

func TestCreateComment(t *testing.T) {
	cfg := config.Load()

	db, err := ConnectPostgres(cfg)

	if err != nil {
		t.Fatal(err)
	}

	res := pb.CommentPost{
		UserId:    uuid.New().String(),
		PostId:    "a2083af4-ee47-47f9-b4ed-096c3b75934a",
		Content:   "Toshket",
		CreatedAt: time.Now().String(),
	}

	comment := NewCommentStorage(db)

	req, err := comment.CreateComment(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestUpdateComment(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.UpdateAComment{
		UserId:  "650cbaea-8ab4-4930-aa32-aeb098956562",
		Id:      "fa10902e-b599-4850-9024-8c20163b9634",
		Content: "dddd",
	}
	comment := NewCommentStorage(db)
	req, err := comment.UpdateComment(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestGetCommentByID(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.CommentId{
		CommentId: "fa10902e-b599-4850-9024-8c20163b9634",
	}
	comment := NewCommentStorage(db)
	req, err := comment.GetCommentByID(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestDeleteComment(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.CommentId{
		CommentId: "94eda20d-7a75-4cdd-9a2d-ca4c97089626",
	}
	comment := NewCommentStorage(db)
	req, err := comment.DeleteComment(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestGetAllUserComments(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.Username{
		Username: "5cb6e285-050b-42ed-b1dd-6bf549a14d6e",
	}
	comment := NewCommentStorage(db)
	req, err := comment.GetAllUserComments(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestGetCommentByUsername(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.Username{
		Username: "650cbaea-8ab4-4930-aa32-aeb098956562",
	}

	comment := NewCommentStorage(db)
	req, err := comment.GetCommentByUsername(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestListComments(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}

	res := pb.CommentList{
		PostId: "ff3c798c-fa32-41d6-88e0-dcf5287aa5b2",
		Limit:  1,
		//Offset: 0,
	}
	comment := NewCommentStorage(db)
	req, err := comment.ListComments(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

func TestGetCommentByPostID(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.PostId{
		Id: "ff3c798c-fa32-41d6-88e0-dcf5287aa5b2",
	}
	comment := NewCommentStorage(db)
	req, err := comment.GetCommentByPostID(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}

// ishlamadi
func TestGetMostlikeCommentPost(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Fatal(err)
	}
	res := pb.PostId{
		Id: "ff3c798c-fa32-41d6-88e0-dcf5287aa5b2",
	}
	comment := NewCommentStorage(db)
	req, err := comment.GetMostlikeCommentPost(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}
