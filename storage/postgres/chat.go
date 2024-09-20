package postgres

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	pb "post-servic/genproto/post"
	"post-servic/storage"
	"time"
)

type ChatStorage struct {
	db *sqlx.DB
}

func NewChatStorage(db *sqlx.DB) storage.ChatStorage {
	return &ChatStorage{
		db: db,
	}
}

func (c *ChatStorage) StartMessaging(in *pb.CreateChat) (*pb.ChatResponse, error) {
	query := `INSERT INTO chat (user1_id, user2_id, created_at) 
	          VALUES ($1, $2, $3) 
	          RETURNING id, user1_id, user2_id, created_at`

	var res pb.ChatResponse
	err := c.db.QueryRowContext(context.Background(), query,
		in.User1Id, in.User2Id, time.Now()).Scan(
		&res.Id,
		&res.User1Id,
		&res.User2Id,
		&res.CratedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat: %v", err)
	}

	return &res, nil
}

func (c *ChatStorage) SendMessage(in *pb.CreateMassage) (*pb.MassageResponse, error) {
	query := `INSERT INTO messages (chat_id, sender_id, content_type, content, created_at) 
	          VALUES ($1, $2, $3, $4) 
	          RETURNING id, chat_id, sender_id, content_type, content, created_at`

	var res pb.MassageResponse
	err := c.db.QueryRowContext(context.Background(), query,
		in.ChatId, in.SenderId, in.ContentType, in.Content, time.Now()).Scan(
		&res.Id,
		&res.ChatId,
		&res.SenderId,
		&res.ContentType,
		&res.Content,
		&res.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat: %v", err)
	}

	return &res, nil
}

func (c *ChatStorage) GetChatMessages(in *pb.List) (*pb.MassageResponseList, error) {
	query := "SELECT id, chat_id, sender_id, content_type, content, created_at, updated_aut, is_read FROM messages WHERE deleted_at = 0"
	args := []interface{}{}

	if in.Limit > 0 {
		query += " LIMIT $" + fmt.Sprintf("%d", len(args)+1)
		args = append(args, in.Limit)
	}
	if in.Offset >= 0 {
		query += " OFFSET $" + fmt.Sprintf("%d", len(args)+1)
		args = append(args, in.Offset)
	}

	rows, err := c.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var massages []*pb.MassageResponse
	for rows.Next() {
		var massage pb.MassageResponse
		if err := rows.Scan(&massage.Id, &massage.ChatId, &massage.SenderId, &massage.ContentType, &massage.Content,
			&massage.CreatedAt, &massage.UpdatedAt, &massage.IsRead); err != nil {
			return nil, err
		}
		massages = append(massages, &massage)
	}

	return &pb.MassageResponseList{
		Massage: massages,
	}, nil
}

func (c *ChatStorage) MessageMarcTrue(in *pb.MassageTrue) (*pb.Message, error) {
	query := `UPDATE posts SET is_read = $1 WHERE id = $2 RETURNING id`

	var postId int64
	err := c.db.QueryRowContext(context.Background(), query, true, in.ChatId).Scan(&postId)
	if err != nil {
		return nil, err
	}

	return &pb.Message{
		Massage: "is read true",
	}, nil
}

func (c *ChatStorage) GetUserChats(in *pb.Username) (*pb.ChatResponseList, error) {
	query := `SELECT id, user1_id, user2_id FROM chats WHERE user_id = $1 ORDER BY created_at DESC`

	rows, err := c.db.QueryContext(context.Background(), query, in.Username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chats []*pb.ChatResponse
	for rows.Next() {
		var chat pb.ChatResponse
		err := rows.Scan(&chat.Id, &chat.User1Id, &chat.User2Id)
		if err != nil {
			return nil, err
		}
		chats = append(chats, &chat)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.ChatResponseList{
		Chat: chats,
	}, nil
}

func (c *ChatStorage) GetUnreadMessages(in *pb.ChatId) (*pb.MassageResponseList, error) {
	query := `SELECT id, chat_id, sender_id, content_type, content, created_at,  is_read
	          FROM messages WHERE chat_id = $1`

	rows, err := c.db.QueryContext(context.Background(), query, in.ChatId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var massages []*pb.MassageResponse
	for rows.Next() {
		var massage pb.MassageResponse
		if err := rows.Scan(&massage.Id, &massage.ChatId, &massage.SenderId, &massage.ContentType, &massage.Content,
			&massage.CreatedAt, &massage.IsRead); err != nil {
			return nil, err
		}
		massages = append(massages, &massage)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pb.MassageResponseList{
		Massage: massages,
	}, nil
}

func (c *ChatStorage) UpdateMessage(in *pb.UpdateMs) (*pb.MassageResponse, error) {
	query := `SELECT content_type
	          FROM messages WHERE id = $1`

	var content_type string
	var res pb.MassageResponse
	err := c.db.QueryRowContext(context.Background(), query, in.MessageId).Scan(
		&content_type)

	if content_type != "text" {
		return nil, fmt.Errorf("failed to update message content_type")
	}

	query1 := `UPDATE messages SET content = $1 WHERE id = $2 RETURNING id, chat_id, sender_id, content_type, content, created_at`

	err = c.db.QueryRowContext(context.Background(), query1, true, in.MessageId).Scan(&res.Id, &res.ChatId, &res.SenderId,
		&res.ContentType, &res.Content, &res.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *ChatStorage) GetTodayMessages(in *pb.ChatId) (*pb.MassageResponseList, error) {
	query := `
		SELECT id, chat_id, sender_id, content_type, content, is_read, created_at 
		FROM messages 
		WHERE chat_id = $1 AND created_at >= CURRENT_DATE
		ORDER BY created_at DESC
		LIMIT 20`

	rows, err := c.db.QueryContext(context.Background(), query, in.ChatId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*pb.MassageResponse
	for rows.Next() {
		var msg pb.MassageResponse
		err := rows.Scan(&msg.Id, &msg.ChatId, &msg.SenderId, &msg.ContentType, &msg.Content, &msg.IsRead, &msg.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &msg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &pb.MassageResponseList{
		Massage: messages,
	}, nil
}

func (c *ChatStorage) DeleteMessage(in *pb.MassageId) (*pb.Message, error) {
	query := `update messages set deleted_at = date_part('epoch', current_timestamp)::INT
                  where id = $1 and deleted_at = 0`

	_, err := c.db.ExecContext(context.Background(), query, in.MassageId)
	if err != nil {
		return nil, err
	}

	return &pb.Message{
		Massage: "Massages muvaffaqiyatli o'chirildi (soft delete)",
	}, nil
}

func (c *ChatStorage) DeleteChat(in *pb.ChatId) (*pb.Message, error) {
	query := `update chat set deleted_at = date_part('epoch', current_timestamp)::INT
                  where id = $1 and deleted_at = 0`

	_, err := c.db.ExecContext(context.Background(), query, in.ChatId)
	if err != nil {
		return nil, err
	}

	return &pb.Message{
		Massage: "Massages muvaffaqiyatli o'chirildi (soft delete)",
	}, nil
}
