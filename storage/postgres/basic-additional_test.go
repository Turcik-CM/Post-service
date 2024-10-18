package postgres

import (
	"log"
	"testing"

	"github.com/google/uuid"
	pb "post-servic/genproto/post"
)

func TestGetUserRecommendation(t *testing.T) {
	db, err := TestConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userID := uuid.NewString()

	postID := uuid.New().String()
	_, err = db.Exec(`INSERT INTO posts(id, user_id, title, content, country, location, hashtag, description)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		postID, userID, "Test Post", "This is a test post.", "Turkey", "New York", "testhashtag", "Test description")
	if err != nil {
		log.Fatal("--- 1 ---", err)
	}

	//_, err = db.Exec(`INSERT INTO likes(user_id, post_id) VALUES ($1, $2)`, userID, postID)
	//if err != nil {
	//	log.Fatal(err)
	//}

	repo := NewBasicAdditional(db)

	userRes, err := repo.GetUserRecommendation(&pb.Username{Username: userID})
	if err != nil {
		log.Fatal("--- 2 ---", err)
	}

	log.Println(userRes)

	if len(userRes.Post) == 0 {
		log.Fatal("Expected posts not found in recommendations")
	}

	//_, err = db.Exec(`DELETE FROM likes WHERE user_id = $1 AND post_id = $2`, userID, postID)
	//if err != nil {
	//	log.Fatal(err)
	//}
	_, err = db.Exec(`DELETE FROM posts WHERE id = $1`, postID)
	if err != nil {
		log.Fatal("--- 3 ---", err)
	}
}

func TestSearchPost(t *testing.T) {
	db, err := TestConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userID := uuid.NewString()
	postID := uuid.New().String()

	_, err = db.Exec(`INSERT INTO posts(id, user_id, title, content, country, location, hashtag, description)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		postID, userID, "Search Post Test", "This is a test search post.", "Turkey", "Istanbul", "testhashtag", "Search description")
	if err != nil {
		log.Fatal(err)
	}

	repo := NewBasicAdditional(db)

	searchReq := &pb.Search{Action: "Search"}
	searchRes, err := repo.SearchPost(searchReq)
	if err != nil {
		log.Fatal(err)
	}

	if len(searchRes.Post) == 0 {
		t.Fatal("Expected post not found in search results")
	}

	searchReqEmpty := &pb.Search{Action: "NonExistentKeyword"}
	searchResEmpty, err := repo.SearchPost(searchReqEmpty)
	if err != nil {
		log.Fatal(err)
	}

	if len(searchResEmpty.Post) > 0 {
		t.Fatal("No posts should be returned for a non-matching search")
	}

	_, err = db.Exec(`DELETE FROM posts WHERE id = $1`, postID)
	if err != nil {
		log.Fatal(err)
	}
}

func TestGetPostsByUsername(t *testing.T) {
	db, err := TestConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userID := uuid.NewString()
	postID1 := uuid.NewString()
	postID2 := uuid.NewString()

	_, err = db.Exec(`INSERT INTO posts(id, user_id, title, content, country, location, hashtag, description)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		postID1, userID, "User Post 1", "Content of user post 1", "Turkey", "Istanbul", "testhashtag", "Description 1")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`INSERT INTO posts(id, user_id, title, content, country, location, hashtag, description)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		postID2, userID, "User Post 2", "Content of user post 2", "Turkey", "Ankara", "testhashtag", "Description 2")
	if err != nil {
		log.Fatal(err)
	}

	repo := NewBasicAdditional(db)

	postsRes, err := repo.GetPostsByUsername(&pb.Username{Username: userID})
	if err != nil {
		log.Fatal(err)
	}

	if len(postsRes.Post) != 2 {
		t.Fatalf("Expected 2 posts, got %d", len(postsRes.Post))
	}

	if postsRes.Post[0].Title != "User Post 2" || postsRes.Post[1].Title != "User Post 1" {
		t.Fatal("Posts are not in the correct order by creation date")
	}
	_, err = db.Exec(`DELETE FROM posts WHERE user_id = $1`, userID)
	if err != nil {
		log.Fatal(err)
	}
}
