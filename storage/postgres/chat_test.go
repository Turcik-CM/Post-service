package postgres

import (
	"fmt"
	"github.com/google/uuid"
	pb "post-servic/genproto/post"
	"post-servic/pkg/config"
	"testing"
)

func TestStartMessaging(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	defer db.Close()

	res := pb.CreateChat{
		User1Id: uuid.New().String(),
		User2Id: uuid.New().String(),
	}

	chat := NewChatStorage(db)

	rep, err := chat.StartMessaging(&res)
	if err != nil {
		t.Errorf("Error starting messaging: %v", err)
	}
	fmt.Println(rep)
}

func TestSendMessage(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	defer db.Close()
	res := pb.CreateMassage{
		ChatId:      "85a8081f-7018-4a1c-bbd9-611ecac4a70c",
		SenderId:    uuid.New().String(),
		ContentType: "text",
		Content:     "dodi",
	}

	chat := NewChatStorage(db)
	rep, err := chat.SendMessage(&res)
	if err != nil {
		t.Errorf("Error sending message: %v", err)
	}
	fmt.Println(rep)
}

func TestGetChatMessages(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	defer db.Close()
	res := pb.List{
		ChatId: "f56d292e-55a4-4a86-bffe-c9eca93e9248",
		Limit:  0,
	}
	chat := NewChatStorage(db)
	rep, err := chat.GetChatMessages(&res)
	if err != nil {
		t.Errorf("Error getting messages: %v", err)
	}
	fmt.Println(rep)
}

func TestMessageMarcTrue(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	defer db.Close()
	res := pb.MassageTrue{
		ChatId: "f56d292e-55a4-4a86-bffe-c9eca93e9248",
	}
	chat := NewChatStorage(db)
	rep, err := chat.MessageMarcTrue(&res)
	if err != nil {
		t.Errorf("Error getting message: %v", err)
	}
	fmt.Println(rep)
}

func TestGetUserChats(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	defer db.Close()
	res := pb.Username{
		Username: "f56d292e-55a4-4a86-bffe-c9eca93e9248",
	}
	chat := NewChatStorage(db)
	rep, err := chat.GetUserChats(&res)
	if err != nil {
		t.Errorf("Error getting user: %v", err)
	}
	fmt.Println(rep)
}

func TestGetUnreadMessages(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	defer db.Close()
	res := pb.ChatId{
		ChatId: "f56d292e-55a4-4a86-bffe-c9eca93e9248",
	}
	chat := NewChatStorage(db)
	rep, err := chat.GetUnreadMessages(&res)
	if err != nil {
		t.Errorf("Error getting unread messages: %v", err)
	}
	fmt.Println(rep)
}

func TestUpdateMessage(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	defer db.Close()
	res := pb.UpdateMs{
		MessageId: "e96b6719-dacb-4729-a993-b1466df6ee9f",
		Text:      "dodi alo dodi",
	}
	chat := NewChatStorage(db)
	rep, err := chat.UpdateMessage(&res)
	if err != nil {
		t.Errorf("Error updating message: %v", err)
	}
	fmt.Println(rep)
}

func TestGetTodayMessages(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	defer db.Close()
	res := pb.ChatId{
		ChatId: "f56d292e-55a4-4a86-bffe-c9eca93e9248",
	}
	chat := NewChatStorage(db)
	rep, err := chat.GetTodayMessages(&res)
	if err != nil {
		t.Errorf("Error getting today messages: %v", err)
	}
	fmt.Println(rep)
}

func TestDeleteMessage(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	defer db.Close()
	res := pb.MassageId{
		MassageId: "0203a647-bae2-4501-acf7-fb008b30401e",
	}
	chat := NewChatStorage(db)
	rep, err := chat.DeleteMessage(&res)
	if err != nil {
		t.Errorf("Error deleting message: %v", err)
	}
	fmt.Println(rep)
}

func TestDeleteChat(t *testing.T) {
	cfg := config.Load()
	db, err := ConnectPostgres(cfg)
	if err != nil {
		t.Errorf("Error connecting to database: %v", err)
	}
	defer db.Close()
	res := pb.ChatId{
		ChatId: "f5697c48-3b72-4681-bb54-97d8d7e4ed5a",
	}
	chat := NewChatStorage(db)
	rep, err := chat.DeleteChat(&res)
	if err != nil {
		t.Errorf("Error deleting chat: %v", err)
	}
	fmt.Println(rep)
}
