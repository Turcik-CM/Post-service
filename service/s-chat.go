package service

import (
	"context"
	pb "post-servic/genproto/post"
)

func (c *PostService) StartMessaging(ctx context.Context, in *pb.CreateChat) (*pb.ChatResponse, error) {
	res, err := c.chat.StartMessaging(in)
	if err != nil {
		c.logger.Error("failed to post cHAT", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) SendMessage(ctx context.Context, in *pb.CreateMassage) (*pb.MassageResponse, error) {
	res, err := c.chat.SendMessage(in)
	if err != nil {
		c.logger.Error("failed to post cHAT", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) GetChatMessages(ctx context.Context, in *pb.List) (*pb.MassageResponseList, error) {
	res, err := c.chat.GetChatMessages(in)
	if err != nil {
		c.logger.Error("failed to post cHAT", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) MessageMarcTrue(ctx context.Context, in *pb.MassageTrue) (*pb.Message, error) {
	res, err := c.chat.MessageMarcTrue(in)
	if err != nil {
		c.logger.Error("failed to post cHAT", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) GetUserChats(ctx context.Context, in *pb.Username) (*pb.ChatResponseList, error) {
	res, err := c.chat.GetUserChats(in)
	if err != nil {
		c.logger.Error("failed to post cHAT", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) GetUnreadMessages(ctx context.Context, in *pb.ChatId) (*pb.MassageResponseList, error) {
	res, err := c.chat.GetUnreadMessages(in)
	if err != nil {
		c.logger.Error("failed to post cHAT", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) UpdateMessage(ctx context.Context, in *pb.UpdateMs) (*pb.MassageResponse, error) {
	res, err := c.chat.UpdateMessage(in)
	if err != nil {
		c.logger.Error("failed to post cHAT", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) GetTodayMessages(ctx context.Context, in *pb.ChatId) (*pb.MassageResponseList, error) {
	res, err := c.chat.GetTodayMessages(in)
	if err != nil {
		c.logger.Error("failed to post cHAT", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) DeleteMessage(ctx context.Context, in *pb.MassageId) (*pb.Message, error) {
	res, err := c.chat.DeleteMessage(in)
	if err != nil {
		c.logger.Error("failed to post cHAT", err)
		return nil, err
	}
	return res, nil
}

func (c *PostService) DeleteChat(ctx context.Context, in *pb.ChatId) (*pb.Message, error) {
	res, err := c.chat.DeleteChat(in)
	if err != nil {
		c.logger.Error("failed to post cHAT", err)
		return nil, err
	}
	return res, nil
}
