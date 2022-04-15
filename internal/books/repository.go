package books

import (
	"books/pkg/models"
	"database/sql"
	"fmt"
	"log"
)

type BookCreate interface {
	CreateBook(
		BookAuthor,
		BookTitle string,
		YearOfRelease int,
		CoverUrl string,
		Description string) (*models.Book, error)
}

type BookDelete interface {
	DeleteBook(bookId int) (bool, error)
}

type UpdateBook interface {
	UpdateAuthor(bookId int, newAuthor string) (bool, error)
	UpdateYearOfRelease(bookId, newYearOfRelease int) (bool, error)
	UpdateBookTitle(bookId int, newTitle string) (bool, error)
	UpdateCover(bookId int, newCoverUrl string) (bool, error)
	UpdateDescription(bookId int, newDescription string) (bool, error)
}

type GetBooks interface {
	GetBookById(bookId int) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	GetAllBooksByAuthor(author string) ([]*models.Book, error)
	GetRandomBook() (*models.Book, error)
}
type BookRepo interface {
	GetBooks
	UpdateBook
	BookDelete
	BookCreate
}

type Repository struct {
	Db *sql.DB
}

func (s *Repository) updateParam(bookId int, column, newValue string) (bool, error) {
	var ExecRow = fmt.Sprintf(`UPDATE "Books" SET %s = '%s' WHERE book_id = %d`, column, newValue, bookId)
	_, exc := s.Db.Exec(ExecRow)
	if exc != nil {
		log.Fatal(fmt.Sprintf(`update %s with book id: %d err: %s`, column, bookId, exc))
		return false, exc
	}
	return true, nil
}
func (s *Repository) GetBookById(bookId int) (*models.Book, error) {
	var book models.Book
	//defer s.Db.Close()
	var null sql.NullFloat64
	row := s.Db.QueryRow(`SELECT 
       b.book_id,
       b.book_author,
       b.book_title,
       b.year_of_release,
       b.book_cover_url,
       b.description, 
       avg(r.rating) 
		FROM "Rating" AS r 
		    RIGHT JOIN "Books" AS b 
		        ON b.book_id = r.book_id 
		WHERE b.book_id = $1 
		GROUP BY b.book_id, b.book_author, b.book_title, b.year_of_release, b.book_cover_url, b.description,r.rating`, bookId)
	if err := row.Scan(
		&book.BookId,
		&book.BookAuthor,
		&book.BookTitle,
		&book.YearOfRelease,
		&book.CoverUrl,
		&book.Description,
		&null); err != nil {
		return &book, err
	}
	if null.Valid {
		book.Rating = null.Float64
	} else {
		book.Rating = 0
	}
	return &book, nil
}
func (s *Repository) CreateBook(BookAuthor, BookTitle string, YearOfRelease int, CoverUrl string, Description string) (*models.Book, error) {
	var id int
	if err := s.Db.QueryRow(`INSERT INTO "Books" 
    		(book_author,
    		 book_title,
    		 year_of_release,
    		 book_cover_url,
    		 description) 
    		VALUES ($1,$2,$3,$4,$5) 
    		RETURNING book_id`,
		BookAuthor,
		BookTitle,
		YearOfRelease,
		CoverUrl,
		Description).Scan(&id); err != nil {
		log.Fatal(fmt.Sprintf(`create book err: %s`, err))
		return &models.Book{}, err
	}
	return &models.Book{BookId: id, BookAuthor: BookAuthor, BookTitle: BookTitle, YearOfRelease: YearOfRelease,
		CoverUrl: CoverUrl, Description: Description, Rating: float64(0)}, nil
}
func (s *Repository) DeleteBook(bookId int) (bool, error) {
	_, err := s.Db.Exec(`DELETE FROM "Books" AS b WHERE b.book_id = $1`, bookId)
	if err != nil {
		log.Fatal(fmt.Sprintf(`delete book_id: %d err: %s`, bookId, err))
		return false, err
	}
	return true, nil
}
func (s *Repository) UpdateAuthor(bookId int, newAuthor string) (bool, error) {
	return s.updateParam(bookId, `book_author`, newAuthor)
}
func (s *Repository) UpdateYearOfRelease(bookId, newYearOfRelease int) (bool, error) {
	stringYear := fmt.Sprint(newYearOfRelease)
	return s.updateParam(bookId, `year_of_release`, stringYear)
}
func (s *Repository) UpdateDescription(bookId int, newDescription string) (bool, error) {
	return s.updateParam(bookId, `description`, newDescription)

}
func (s *Repository) UpdateCover(bookId int, newCoverUrl string) (bool, error) {
	return s.updateParam(bookId, `book_cover_url`, newCoverUrl)
}
func (s *Repository) UpdateBookTitle(bookId int, newTitle string) (bool, error) {
	return s.updateParam(bookId, `book_title`, newTitle)

}

func (s *Repository) GetAllBooks() ([]*models.Book, error) {
	var ListOfBooks []*models.Book
	var book models.Book
	var null sql.NullFloat64
	rows, errs := s.Db.Query(`select 
       b.book_id,
       b.book_author,
       b.book_title,
       b.year_of_release,
       b.book_cover_url,
       b.description, 
       avg(r.rating) 
		from "Rating" as r 
		    right join "Books" as b 
		        on b.book_id = r.book_id  
		GROUP BY b.book_id, b.book_author, b.book_title, b.year_of_release, b.book_cover_url, b.description,r.rating`)
	defer rows.Close()
	if errs != nil {
		log.Fatal(errs)
		return ListOfBooks, errs
	}
	for rows.Next() {
		if err := rows.Scan(&book.BookId, &book.BookAuthor, &book.BookTitle, &book.YearOfRelease, &book.CoverUrl, &book.Description, &null); err != nil {
			log.Fatal(err)
			return ListOfBooks, err
		}
		if null.Valid {
			book.Rating = null.Float64
		} else {
			book.Rating = 0
		}
		ListOfBooks = append(ListOfBooks, &book)
	}
	return ListOfBooks, nil
}
func (s *Repository) GetAllBooksByAuthor(author string) ([]*models.Book, error) {
	var AuthorBooks []*models.Book
	var null sql.NullFloat64
	rows, exc := s.Db.Query(`SELECT 
       b.book_id,
       b.book_author,
       b.book_title,
       b.year_of_release,
       b.book_cover_url,
       b.description, 
       avg(r.rating) 
		FROM "Rating" AS r 
		    RIGHT JOIN "Books" AS b 
		        ON b.book_id = r.book_id 
	WHERE b.book_author = $1	
	GROUP BY b.book_id, b.book_author, b.book_title, b.year_of_release, b.book_cover_url, b.description,r.rating`, author)
	defer rows.Close()
	if exc != nil {
		log.Fatal(exc)
		return AuthorBooks, exc
	}
	for rows.Next() {
		var book models.Book
		if errs := rows.Scan(&book.BookId, &book.BookAuthor, &book.BookTitle, &book.YearOfRelease, &book.CoverUrl, &book.Description, &null); errs != nil {
			log.Fatal(errs)
			return AuthorBooks, errs
		}
		if null.Valid {
			book.Rating = null.Float64
		} else {
			book.Rating = float64(0)
		}
		AuthorBooks = append(AuthorBooks, &book)
	}
	return AuthorBooks, nil
}
func (s *Repository) GetRandomBook() (*models.Book, error) {
	var book models.Book
	var null sql.NullFloat64
	var SelectRow = `SELECT b.book_id,
       b.book_author,
       b.book_title,
       b.year_of_release,
       b.book_cover_url,
       b.description, 
       avg(r.rating) 
		FROM "Rating" AS r 
		    RIGHT JOIN "Books" as b 
		        ON b.book_id = r.book_id  
		GROUP BY b.book_id, b.book_author, b.book_title, b.year_of_release, b.book_cover_url, b.description,r.rating
		ORDER BY random() LIMIT 1;`
	row := s.Db.QueryRow(SelectRow)
	if exc := row.Scan(&book.BookId, &book.BookAuthor, &book.BookTitle, &book.YearOfRelease, &book.CoverUrl, &book.Description, &null); exc != nil {
		log.Fatal(exc)
		return &book, exc
	}
	if null.Valid {
		book.Rating = null.Float64
	} else {
		book.Rating = float64(0)
	}
	return &book, nil
}
