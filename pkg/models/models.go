package models

type Book struct {
	BookId        int     `json:"bookId"`
	BookAuthor    string  `json:"bookAuthor"`
	BookTitle     string  `json:"bookTitle"`
	YearOfRelease int     `json:"yearOfRelease"`
	CoverUrl      string  `json:"coverUrl"`
	Description   string  `json:"description"`
	Rating        float64 `json:"rating"`
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
}

type UserReview struct {
	User
	Review   string `json:"review"`
	ReviewId int    `json:"reviewId"`
}

