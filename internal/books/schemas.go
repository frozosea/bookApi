package books

type CreateBookSchema struct {
	BookAuthor    string `json:"book_author" form:"Author"`
	BookTitle     string `json:"book_title" form:"title"`
	YearOfRelease int    `json:"year_of_release" form:"yearOfRelease"`
	CoverUrl      string `json:"cover_url" form:"coverUrl"`
	Description   string `json:"description" form:"description"`
}
type BaseResponse struct {
	Success bool `json:"success"`
}
type BookIdSchema struct {
	BookId int `form:"bookId" json:"book_id"`
}
type UpdateAuthorSchema struct {
	BookIdSchema
	NewAuthor string `json:"new_author" binding:"required" form:"NewAuthor"`
}
type UpdateYearOfReleaseSchema struct {
	BookIdSchema
	NewYearOfRelease int `json:"new_year_of_release" binding:"required" form:"NewYearOfRelease"`
}
type UpdateBookTitleSchema struct {
	BookIdSchema
	NewBookTitle string `json:"new_book_title" binding:"required" form:"NewBookTitle"`
}
type UpdateCoverSchema struct {
	BookIdSchema
	NewCoverUrl string `json:"new_cover_url" binding:"required" form:"NewCoverUrl"`
}
type UpdateDescriptionSchema struct {
	BookIdSchema
	NewDescription string `json:"new_description" binding:"required" form:"NewDescription"`
}
type GetAllBooksByAuthorSchema struct {
	Author string `json:"author" form:"Author"`
}
