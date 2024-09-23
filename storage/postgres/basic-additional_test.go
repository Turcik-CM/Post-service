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
