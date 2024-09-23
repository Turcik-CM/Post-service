package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"post-servic/genproto/post"
	"post-servic/pkg/config"
	"post-servic/pkg/logger"
	"post-servic/service"
	"post-servic/storage/postgres"
)

func main() {
	logger := logger.InitLogger()
	cfg := config.Load()

	db, err := postgres.ConnectPostgres(cfg)
	if err != nil {
		logger.Error("Error connecting to database")
		log.Fatal(err)
	}

	postSt := postgres.NewPostStorage(db)
	likeSt := postgres.NewLikeStorage(db)
	commentSt := postgres.NewCommentStorage(db)
	chatST := postgres.NewChatStorage(db)

	postSR := service.NewPostService(chatST, postSt, likeSt, commentSt, logger)

	listen, err := net.Listen("tcp", cfg.POST_PORT)
	fmt.Println("Listening on " + cfg.POST_PORT)
	if err != nil {
		logger.Error("Error listening on port " + cfg.POST_PORT)
		log.Fatal(err)
	}

	server := grpc.NewServer()
	post.RegisterPostServiceServer(server, postSR)

	if err := server.Serve(listen); err != nil {
		logger.Error("Error serving on port " + cfg.POST_PORT)
		log.Fatal(err)
	}
}
