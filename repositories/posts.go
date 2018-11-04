package repositories

import (
	"database/sql"

	models "corkboard/internal/pkg/models"
)

type PostsRespository struct {
	DB *sql.DB
}

func (repo *PostsRespository) GetPostByID(id int) (*models.Post, error) {
	post := &models.Post{}

	row := repo.DB.QueryRow(`
		SELECT
			id, user_id, body, photo_name
		FROM posts
		WHERE id = ?
	`, id)

	err := row.Scan(&post.ID, &post.UserID, &post.Body, &post.PhotoName)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return post, nil
}