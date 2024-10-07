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
		PostId:    "f380a915-4bd8-434e-a699-7b82bbfb3906",
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
		UserId:  "5d2b4c0d-7fad-4d8a-b708-7124746ec4cd",
		Id:      "b894b386-c4a1-45df-8c89-d5edd2e18201",
		Content: "Toshket",
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
		CommentId: "b894b386-c4a1-45df-8c89-d5edd2e18201",
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
		CommentId: "b894b386-c4a1-45df-8c89-d5edd2e18201",
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
		Username: "c732f495-bc83-492a-bdae-e98fc430345c",
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
		Username: "c732f495-bc83-492a-bdae-e98fc430345c",
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
		//PostId: "ff3c798c-fa32-41d6-88e0-dcf5287aa5b2",
		Limit: 3,
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
		Id: "f380a915-4bd8-434e-a699-7b82bbfb3906",
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
		Id: "f380a915-4bd8-434e-a699-7b82bbfb3906",
	}
	comment := NewCommentStorage(db)
	req, err := comment.GetMostlikeCommentPost(&res)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
}
