package handler

import (
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}
}
func GetEnv(EnvName string) string {
	LoadedEnv := os.Getenv(EnvName)
	return LoadedEnv
}
func PackageSourceValidator(PackageSource string) bool {
	re := regexp.MustCompile(`^[0-9]{24}$`)
	return re.MatchString(PackageSource)
}
