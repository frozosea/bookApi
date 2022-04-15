package settings

import (
	"database/sql"
	"fmt"
	"github.com/go-ini/ini"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"os"
)

type DataBase struct {
	DatabaseUser     string
	DatabasePassword string
	Database         string
}

var DbSettings = &DataBase{}

type App struct {
	AccessTokenExpiration  int
	RefreshTokenExpiration int
	JwtSecretKey           string
}

type Server struct {
	HttpPort int
}

func SetupDatabaseConfig() *DataBase {
	DbSettings.DatabaseUser = os.Getenv(`DATABASE_USER`)
	DbSettings.DatabasePassword = os.Getenv(`DATABASE_PASSWORD`)
	DbSettings.Database = os.Getenv(`DATABASE_NAME`)
	return DbSettings
}

func getSqlUrl(dbConf *DataBase) string {
	return fmt.Sprintf("host=postgres port=%d user=%s password=%s dbname=%s sslmode=disable",
		5432,
		dbConf.DatabaseUser,
		dbConf.DatabasePassword,
		dbConf.Database)
}

//GetDatabase get sql database from environment variables
func GetDatabase() (*sql.DB, error) {
	dbConf := SetupDatabaseConfig()
	db, err := sql.Open(`pgx`, fmt.Sprintf("host=postgres-db port=%d user=%s password=%s dbname=%s sslmode=disable",
		5432,
		dbConf.DatabaseUser,
		dbConf.DatabasePassword,
		dbConf.Database))
	if err != nil {
		log.Fatalf(`open database err:%s`, err.Error())
		return db, err
	}
	if exc := db.Ping(); exc != nil {
		return db, exc

	}
	return db, nil
}

//GetJwtSecret from environment
func GetJwtSecret() string {
	secretKey := os.Getenv(`JWT_SECRET_KEY`)
	return secretKey
}

//GetTokenSettings from conf/cfg.ini (access token expiration, refresh token expiration)
func GetTokenSettings() *App {
	var app App
	cfg, err := ini.Load(`conf/cfg.ini`)
	tokenSection := cfg.Section(`TOKEN`)
	if err != nil {
		log.Fatalf(`read config from ini file err:%s`, err)
	}
	if exc := tokenSection.MapTo(&app); exc != nil {
		log.Fatalf(`read from conf.ini error: %s`, exc.Error())
	}
	tokenSection.MapTo(&app)
	app.JwtSecretKey = GetJwtSecret()
	return &app
}

//GetServerSetting from conf/cfg.ini file
func GetServerSetting() *Server {
	var serve Server
	cfg, err := ini.Load(`conf/cfg.ini`)
	if err != nil {
		log.Fatalf(`read config from ini file err:%s`, err)
	}
	serverSection := cfg.Section(`SERVER`)
	if exc := serverSection.MapTo(&serve); exc != nil {
		log.Fatalf(`read from conf.ini error: %s`, exc.Error())
	}
	serverSection.MapTo(&serve)
	return &serve
}
func CreateTables(db *sql.DB) {
	_, exc := db.Exec(`CREATE TABLE IF NOT EXISTS "User"
(
    id         serial       not null unique,
    first_name varchar(255) not null,
    last_name  varchar(255) not null,
    username   varchar(255) not null unique,
    password   varchar      not null
);
CREATE TABLE IF NOT EXISTS "Books"
(
    book_id         serial        not null unique,
    book_author     varchar(255)  not null,
    book_title      varchar(255)  not null,
    year_of_release int           not null,
    book_cover_url  varchar(255)  not null,
    description     varchar(1000) not null
);
CREATE TABLE IF NOT EXISTS "Rating"
(
    id      serial not null unique,
    book_id int,
    user_id int,
    rating  int    not null,
    FOREIGN KEY (book_id) REFERENCES "Books" (book_id),
    FOREIGN KEY (user_id) REFERENCES "User" (id)
);
CREATE TABLE IF NOT EXISTS "Review"
(
    id      serial        not null unique,
    book_id int,
    user_id int,
    review  varchar(1000) not null,
    FOREIGN KEY (book_id) REFERENCES "Books" (book_id),
    FOREIGN KEY (user_id) REFERENCES "User" (id)
);`)
	if exc != nil {
		log.Fatalf(`create tables err: %s`, exc.Error())
	}
	fmt.Println(`tables was created`)
}
