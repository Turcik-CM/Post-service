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

	postSR := service.NewPostService(postSt, likeSt, commentSt, logger)

	listen, err := net.Listen("tcp", cfg.POST_SERVICE)
	fmt.Println("Listening on " + cfg.POST_SERVICE)
	if err != nil {
		logger.Error("Error listening on port " + cfg.POST_SERVICE)
		log.Fatal(err)
	}

	server := grpc.NewServer()
	post.RegisterPostServiceServer(server, postSR)

	if err := server.Serve(listen); err != nil {
		logger.Error("Error serving on port " + cfg.POST_SERVICE)
		log.Fatal(err)
	}
}
