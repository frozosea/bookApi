package rating

import (
	"database/sql"
	"log"
)

type Rating interface {
	SetRating(userId, bookId, rating int) (bool, error)
}

type Repository struct {
	Db *sql.DB
}

func (s *Repository) SetRating(bookId, rating int) (bool, error) {
	_, exc := s.Db.Exec(`INSERT INTO "Rating" (book_id, Rating) VALUES ($1,$2)`,
		bookId, rating)
	if exc != nil {
		log.Fatalf(`set rating err: %s`, exc.Error())
		return false, exc
	}
	return true, exc
}
