package main

import (
	"books/internal/auth"
	"books/internal/books"
	"books/internal/healthCheck"
	"books/internal/middleware"
	"books/internal/rating"
	"books/internal/review"
	"books/internal/user"
	logger "books/pkg/log"
	accessLogger "books/pkg/log/access"
	booksLogger "books/pkg/log/books"
	userLogger "books/pkg/log/user"
	"books/settings"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func SetupAuth(db *sql.DB, secretKey string, AccessTokenExpiration, RefreshTokenExpiration int) *auth.Api {
	repository := auth.Repository{
		Db:                     db,
		SecretKey:              secretKey,
		AccessTokenExpiration:  AccessTokenExpiration,
		RefreshTokenExpiration: RefreshTokenExpiration}
	service := auth.Service{Repo: &repository}
	appHandler := auth.Api{Service: &service}
	return &appHandler
}
func SetupUserApi(db *sql.DB, jwtSecret string) *user.Api {
	loggerEngine := logger.New(`userService`, 84600, 604800)
	userLogger := userLogger.UserLogger{Logger: loggerEngine}
	repository := user.Repository{Db: db}
	service := user.Service{Repository: &repository, Logger: &userLogger}
	appHandler := user.Api{Service: &service, SecretKey: jwtSecret}
	return &appHandler
}
func SetupBooksApi(db *sql.DB, secretKey string) *books.Api {
	loggerEngine := logger.New(`booksService`, 84600, 604800)
	booksLogger := booksLogger.BookLogger{Logger: loggerEngine}
	repository := books.Repository{Db: db}
	service := books.Service{Repository: &repository, Logger: &booksLogger}
	appHandler := books.Api{Service: &service, SecretKey: secretKey}
	return &appHandler
}
func SetupRatingApi(db *sql.DB) *rating.Api {
	repository := rating.Repository{Db: db}
	service := rating.Service{Repository: &repository}
	appHandler := rating.Api{Service: &service}
	return &appHandler
}
func SetupReviewApi(db *sql.DB, secretKey string) *review.Api {
	repository := review.Repository{Db: db}
	service := review.Service{Repository: repository}
	appHandler := review.Api{Service: &service, SecretKey: secretKey}
	return &appHandler
}

func InitApp(db *sql.DB, jwtSecret string, AccessTokenExpiration, RefreshTokenExpiration int) *gin.Engine {
	settings.CreateTables(db)
	authApi := SetupAuth(db, jwtSecret, AccessTokenExpiration, RefreshTokenExpiration)
	userApi := SetupUserApi(db, jwtSecret)
	booksApi := SetupBooksApi(db, jwtSecret)
	ratingApi := SetupRatingApi(db)
	reviewApi := SetupReviewApi(db, jwtSecret)
	healthCheckApi := healthCheck.HealthCheck{}
	authMiddleware := middleware.AuthMiddleware{SecretKey: jwtSecret}
	r := gin.New()
	accessLogger.SetupApiLogger(r)
	r.GET("/", healthCheckApi.HealthCheck)
	r.POST("/auth", authApi.Login)
	r.POST("/register", authApi.Register)
	r.GET("/books/GetBookByID", booksApi.GetBookByID)
	r.GET("/books/GetAllBooks", booksApi.GetAllBooks)
	r.GET("/books/author", booksApi.GetBookByAuthor)
	r.GET("/books/random", booksApi.GetRandomBook)
	r.GET("/review/GetReviewsByBookId", reviewApi.GetAllReviewsByBookID)
	r.POST("/rating/setRating", ratingApi.SetRating)
	r.GET("/user/GetInfoAboutUser", userApi.GetInfoAboutUser)
	protectedApi := r.Group("/v1")
	protectedApi.Use(authMiddleware.Auth)
	{
		protectedApi.PUT("/user/UpdateUsername", userApi.UpdateUsername)
		protectedApi.PUT("/user/UpdateFirstName", userApi.UpdateFirstName)
		protectedApi.PUT("/user/UpdateLastName", userApi.UpdateLastname)
		protectedApi.PUT("/user/UpdatePassword", userApi.UpdatePassword)
		protectedApi.POST("/books/CreateBook", booksApi.CreateNewBook)
		protectedApi.PUT("/books/updateYearOfRelease", booksApi.UpdateYearOfRelease)
		protectedApi.PUT("/books/updateAuthor", booksApi.UpdateAuthor)
		protectedApi.PUT("/books/updateDescription", booksApi.UpdateDescription)
		protectedApi.PUT("/books/updateCover", booksApi.UpdateCover)
		protectedApi.PUT("/books/updateTitle", booksApi.UpdateTitle)
		protectedApi.POST("/review/WriteReview", reviewApi.WriteNewReview)
		protectedApi.PUT("/review/UpdateReview", reviewApi.UpdateWriteReview)
		protectedApi.DELETE("/review/DeleteReview", reviewApi.DeleteReview)
	}
	return r
}
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(fmt.Sprintf(`read from .env file err:%s`, err.Error()))
	}
	db, err := settings.GetDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}
	appSettings := settings.GetTokenSettings()
	r := InitApp(db, appSettings.JwtSecretKey, appSettings.AccessTokenExpiration, appSettings.RefreshTokenExpiration)
	if err := r.Run(`:8080`); err != nil {
		log.Fatalf(`run server err: %s`, err.Error())
	}
}
