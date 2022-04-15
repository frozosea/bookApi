package bookslogger

import (
	"books/pkg/log"
	"github.com/sirupsen/logrus"
)

type IBookLogger interface {
	logger.IBaseExceptionLogger
	AuthorUpdatedLog(bookId int, newAuthor string)
	DescriptionUpdateLog(bookId int, newDescription string)
	YearOfReleaseUpdateLog(bookId, newYearOfRelease int)
	CoverUrlUpdateLog(bookId int, newCoverUrl string)
	CreateBookLog(bookId, userId int)
}

type BookLogger struct {
	Logger *logrus.Logger
}

func (b *BookLogger) ExceptionLog(err error) {
	b.Logger.Error(err.Error())
}
func (b *BookLogger) AuthorUpdatedLog(bookId int, newAuthor string) {
	b.Logger.Infof(`book id: %d new author : %s`, bookId, newAuthor)
}
func (b *BookLogger) DescriptionUpdateLog(bookId int, newDescription string) {
	b.Logger.Infof(`book id: %d new description: %s`, bookId, newDescription)
}
func (b *BookLogger) YearOfReleaseUpdateLog(bookId, newYearOfRelease int) {
	b.Logger.Infof(`book id: %d new year of release: %d`, bookId, newYearOfRelease)

}
func (b *BookLogger) CoverUrlUpdateLog(bookId int, newCoverUrl string) {
	b.Logger.Infof(`book id: %d new cover url : %s`, bookId, newCoverUrl)

}
func (b *BookLogger) CreateBookLog(bookId, userId int) {
	b.Logger.Infof(`book id: %d user id: %d was created`, bookId, userId)
}
