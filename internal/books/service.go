package books

import (
	bookslogger "books/pkg/log/books"
	"books/pkg/models"
)

type BookServiceInterface interface {
	CreateBook(BookAuthor, BookTitle string, YearOfRelease int, CoverUrl string, Description string) (*models.Book, error)
	DeleteBook(bookId int) (bool, error)
	GetBookById(bookId int) (*models.Book, error)
	GetAllBooks() ([]*models.Book, error)
	UpdateAuthor(bookId int, newAuthor string) (bool, error)
	UpdateYearOfRelease(bookId, newYearOfRelease int) (bool, error)
	UpdateBookTitle(bookId int, newTitle string) (bool, error)
	UpdateCover(bookId int, newCoverUrl string) (bool, error)
	UpdateDescription(bookId int, newDescription string) (bool, error)
	GetAllBooksByAuthor(author string) ([]*models.Book, error)
	GetRandomBook() (*models.Book, error)
}
type Service struct {
	Repository BookRepo
	Logger     bookslogger.IBookLogger
}

//CreateBook insert new book into database
func (b *Service) CreateBook(BookAuthor, BookTitle string, YearOfRelease int, CoverUrl string, Description string) (*models.Book, error) {
	book, err := b.Repository.CreateBook(BookAuthor, BookTitle, YearOfRelease, CoverUrl, Description)
	if err != nil {
		go b.Logger.ExceptionLog(err)
		return book, err
	}
	return book, err
}

//DeleteBook delete book by id
func (b *Service) DeleteBook(bookId int) (bool, error) {
	ok, err := b.Repository.DeleteBook(bookId)
	if err != nil {
		go b.Logger.ExceptionLog(err)
		return ok, err
	}
	return ok, err
}

//GetBookById ...
func (b *Service) GetBookById(bookId int) (*models.Book, error) {
	book, err := b.Repository.GetBookById(bookId)
	if err != nil {
		go b.Logger.ExceptionLog(err)
		return book, err
	}
	return book, err
}

//GetAllBooks ...
func (b *Service) GetAllBooks() ([]*models.Book, error) {
	books, err := b.Repository.GetAllBooks()
	if err != nil {
		go b.Logger.ExceptionLog(err)
		return books, err
	}
	return books, err
}

//UpdateAuthor ...
func (b *Service) UpdateAuthor(bookId int, newAuthor string) (bool, error) {
	ok, err := b.Repository.UpdateAuthor(bookId, newAuthor)
	if err != nil {
		go b.Logger.ExceptionLog(err)
		return ok, err
	}
	go b.Logger.AuthorUpdatedLog(bookId, newAuthor)
	return ok, err
}

//UpdateYearOfRelease ...
func (b *Service) UpdateYearOfRelease(bookId, newYearOfRelease int) (bool, error) {
	ok, err := b.Repository.UpdateYearOfRelease(bookId, newYearOfRelease)
	if err != nil {
		go b.Logger.ExceptionLog(err)
		return ok, err
	}
	go b.Logger.YearOfReleaseUpdateLog(bookId, newYearOfRelease)
	return ok, err
}

//UpdateBookTitle ...
func (b *Service) UpdateBookTitle(bookId int, newTitle string) (bool, error) {
	ok, err := b.Repository.UpdateBookTitle(bookId, newTitle)
	if err != nil {
		go b.Logger.ExceptionLog(err)
		return ok, err
	}
	return ok, err
}

//UpdateCover ...
func (b *Service) UpdateCover(bookId int, newCoverUrl string) (bool, error) {
	ok, err := b.Repository.UpdateCover(bookId, newCoverUrl)
	if err != nil {
		go b.Logger.ExceptionLog(err)
		return ok, err
	}
	go b.Logger.CoverUrlUpdateLog(bookId, newCoverUrl)
	return ok, err
}

//UpdateDescription ...
func (b *Service) UpdateDescription(bookId int, newDescription string) (bool, error) {
	ok, err := b.Repository.UpdateDescription(bookId, newDescription)
	if err != nil {
		go b.Logger.ExceptionLog(err)
		return ok, err
	}
	go b.Logger.DescriptionUpdateLog(bookId, newDescription)
	return ok, err
}

//GetAllBooksByAuthor ...
func (b *Service) GetAllBooksByAuthor(author string) ([]*models.Book, error) {
	books, err := b.Repository.GetAllBooksByAuthor(author)
	if err != nil {
		go b.Logger.ExceptionLog(err)
		return books, err
	}
	return books, err
}

//GetRandomBook ...
func (b *Service) GetRandomBook() (*models.Book, error) {
	book, err := b.Repository.GetRandomBook()
	if err != nil {
		go b.Logger.ExceptionLog(err)
		return book, err
	}
	return book, err
}
