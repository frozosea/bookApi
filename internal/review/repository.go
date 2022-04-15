package review

import (
	"books/pkg/models"
	"database/sql"
	"log"
)

type Review interface {
	WriteReview
	UpdateReview
	DeleteReview
}
type WriteReview interface {
	WriteReview(userId, bookId int, review string) (int, error)
}
type UpdateReview interface {
	UpdateReview(reviewId int, newReview string) bool
}
type GetReviews interface {
	GetReviews(bookID int) ([]models.UserReview, error)
}
type DeleteReview interface {
	DeleteReview(reviewId int) bool
}

type Repository struct {
	Db *sql.DB
}

func (s *Repository) WriteReview(userId, bookId int, review string) (int, error) {
	var reviewId int
	row := s.Db.QueryRow(`INSERT INTO "Review" (book_id, user_id, Review) VALUES ($1,$2,$3) RETURNING id;`, bookId, userId, review)
	if err := row.Scan(&reviewId); err != nil {
		log.Fatalf(`create review err: %s`, err)
		return reviewId, err
	}
	return reviewId, nil
}

func (s *Repository) UpdateReview(reviewId int, newReview string) bool {
	_, _ = s.Db.Exec(`UPDATE "Review" AS r SET Review = $1 WHERE r.id = $2;`, newReview, reviewId)
	return true
}
func (s *Repository) GetReviews(bookId int) ([]models.UserReview, error) {
	var reviews []models.UserReview
	rows, err := s.Db.Query(`SELECT
       r.id,
       r.Review, 
       u.id, 
       u.first_name, 
       u.last_name, 
       u.username 
	FROM "Review" AS r 
	    LEFT JOIN "User" as U on U.id = r.user_id WHERE r.book_id = $1;`, bookId)
	if err != nil {
		log.Fatal(err)
		return reviews, err
	}
	defer rows.Close()
	for rows.Next() {
		var UserReview models.UserReview
		if exceptions := rows.Scan(
			&UserReview.ReviewId,
			&UserReview.Review,
			&UserReview.Id,
			&UserReview.FirstName,
			&UserReview.LastName,
			&UserReview.Username); exceptions != nil {
			log.Fatal(exceptions)
			return reviews, exceptions
		}
		reviews = append(reviews, UserReview)
	}
	return reviews, nil
}
func (s *Repository) DeleteReview(reviewId int) bool {
	//defer s.Db.Close()
	_, _ = s.Db.Exec(`DELETE FROM "Review" AS r WHERE r.id = $1`, reviewId)
	return true
}
