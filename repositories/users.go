package repositories

import (
    "database/sql"

    models "corkboard/internal/pkg/models"
)

type UsersRepository struct {
    DB *sql.DB
}

func (repo *UsersRepository) GetUserByID(id int) (models.User, error) {
    user := &models.User{}

    row := repo.DB.QueryRow(`
        SELECT
            id, username, email, account_status, profile_photo_name
        FROM users
        WHERE id = ?
    `, id)

    err := row.Scan(
        &user.ID,
        &user.Username,
        &user.Email,
        &user.AccountStatus,
        &user.ProfilePhotoName
    )
    if err == sql.ErrNoRows {
        return nil, nil
    } else if err != nil {
        return nil, err
    }

    return user, nil
}
