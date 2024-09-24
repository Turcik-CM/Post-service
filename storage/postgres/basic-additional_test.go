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

	// Создаем пост
	postID := uuid.New().String()
	_, err = db.Exec(`INSERT INTO posts(id, user_id, title, content, country, location, hashtag, description)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		postID, userID, "Test Post", "This is a test post.", "Turkey", "New York", "testhashtag", "Test description")
	if err != nil {
		log.Fatal(err)
	}

	// Лайкаем пост
	_, err = db.Exec(`INSERT INTO likes(user_id, post_id) VALUES ($1, $2)`, userID, postID)
	if err != nil {
		log.Fatal(err)
	}

	// Инициализируем репозиторий
	repo := NewBasicAdditional(db)

	// Выполняем метод GetUserRecommendation
	userRes, err := repo.GetUserRecommendation(&pb.Username{Username: userID})
	if err != nil {
		log.Fatal(err)
	}

	// Проверяем, что пост возвращается
	if len(userRes.Post) == 0 {
		log.Fatal("Expected posts not found in recommendations")
	}

	// Чистим созданные записи
	_, err = db.Exec(`DELETE FROM likes WHERE user_id = $1 AND post_id = $2`, userID, postID)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`DELETE FROM posts WHERE id = $1`, postID)
	if err != nil {
		log.Fatal(err)
	}
}

func TestSearchPost(t *testing.T) {
	db, err := TestConnect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create test data
	userID := uuid.NewString()
	postID := uuid.New().String()

	// Insert a post into the 'posts' table
	_, err = db.Exec(`INSERT INTO posts(id, user_id, title, content, country, location, hashtag, description)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		postID, userID, "Search Post Test", "This is a test search post.", "Turkey", "Istanbul", "testhashtag", "Search description")
	if err != nil {
		log.Fatal(err)
	}

	// Initialize repository
	repo := NewBasicAdditional(db)

	// Test SearchPost with a valid action (title, description, or hashtag)
	searchReq := &pb.Search{Action: "Search"}
	searchRes, err := repo.SearchPost(searchReq)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the post is found
	if len(searchRes.Post) == 0 {
		t.Fatal("Expected post not found in search results")
	}

	// Test SearchPost with an action that doesn't match any post
	searchReqEmpty := &pb.Search{Action: "NonExistentKeyword"}
	searchResEmpty, err := repo.SearchPost(searchReqEmpty)
	if err != nil {
		log.Fatal(err)
	}

	// Ensure no posts are returned for unmatched keyword
	if len(searchResEmpty.Post) > 0 {
		t.Fatal("No posts should be returned for a non-matching search")
	}

	// Clean up test data
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

	// Create a test user and posts
	userID := uuid.NewString()
	postID1 := uuid.NewString()
	postID2 := uuid.NewString()

	// Insert two posts for the same user
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

	// Initialize repository
	repo := NewBasicAdditional(db)

	// Execute the GetPostsByUsername method
	postsRes, err := repo.GetPostsByUsername(&pb.Username{Username: userID})
	if err != nil {
		log.Fatal(err)
	}

	// Check if two posts are returned
	if len(postsRes.Post) != 2 {
		t.Fatalf("Expected 2 posts, got %d", len(postsRes.Post))
	}

	// Check if posts are ordered by creation date (assuming postID2 was inserted later and should come first)
	if postsRes.Post[0].Title != "User Post 2" || postsRes.Post[1].Title != "User Post 1" {
		t.Fatal("Posts are not in the correct order by creation date")
	}

	// Cleanup test data
	_, err = db.Exec(`DELETE FROM posts WHERE user_id = $1`, userID)
	if err != nil {
		log.Fatal(err)
	}
}
