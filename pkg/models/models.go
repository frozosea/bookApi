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

//func Setup(repo *sqlRepo.BaseSqlRepo) {
//	db, err := repo.Connect()
//	if err != nil {
//		log.Fatalln(fmt.Sprintf(`connect to database err: %s`, err))
//	}
//	defer db.Close()
//	const CreateUserTableRow = `CREATE TABLE IF NOT EXISTS "User" (
//			id serial not null unique,
//			first_name varchar(255),
//			last_name varchar(255),
//    		username varchar(255) unique ,
//			password varchar(255));`
//	const CreateBookTableQueryRow = `
//	CREATE TABLE IF NOT EXISTS "Books" (
//	book_id serial not null primary key unique ,
//	book_author varchar(255),
//	book_title varchar(255),
//	year_of_release int,
//	book_cover_url varchar,
//	description varchar(255)
//	);
//	`
//	const CreateRatingTableRow = `
//		CREATE TABLE IF NOT EXISTS "Rating" (
//			id serial not null unique ,
//			book_id int,
//			user_id int,
//			rating int not null,
//			foreign key (book_id) references "Books"(book_id),
//			foreign key (user_id) references "User"(id)
//
//	);`
//	const CreateReviewTableRow = `
//	CREATE TABLE IF NOT EXISTS "Review" (
//		id serial not null unique,
//		book_id int,
//		user_id int,
//		review varchar(500),
//		foreign key (book_id) references "Books"(book_id),
//		foreign key (user_id) references "User"(id)
//	);`
//	_, err = db.Exec(CreateUserTableRow)
//	if err != nil {
//		log.Fatal(fmt.Sprintf(`create user table err: %s`, err))
//	}
//	_, err = db.Exec(CreateBookTableQueryRow)
//	if err != nil {
//		log.Fatal(fmt.Sprintf(`create book table err: %s`, err))
//	}
//	_, err = db.Exec(CreateRatingTableRow)
//	if err != nil {
//		log.Fatal(fmt.Sprintf(`create Rating table err: %s`, err))
//	}
//	_, err = db.Exec(CreateReviewTableRow)
//	if err != nil {
//		log.Fatal(fmt.Sprintf(`create Review table err: %s`, err))
//	}
//
//}
