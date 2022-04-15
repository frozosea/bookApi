package utils

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"runtime"
	"strings"
)

func LoadEnv() {
	projectPath, exc := os.Getwd()
	if exc != nil {
		log.Fatalf(`get path err: %s`, exc)
	}
	if runtime.GOOS == "windows" {
		workPath := strings.Split(projectPath, "\\")
		envPath := strings.Join(workPath[:len(workPath)-3], "\\")
		if err := godotenv.Load(fmt.Sprintf(`%s/.env`, envPath)); err != nil {
			log.Fatal(fmt.Sprintf(`read from .env file err:%s`, err.Error()))
		}
	} else {
		workPath := strings.Split(projectPath, "/")
		envPath := strings.Join(workPath[:len(workPath)-3], "/")
		if err := godotenv.Load(fmt.Sprintf(`%s/.env`, envPath)); err != nil {
			log.Fatal(fmt.Sprintf(`read from .env file err:%s`, err.Error()))
		}
	}

}
