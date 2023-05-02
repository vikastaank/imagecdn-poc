package constants

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	// /home/vikastaank/go/src/github.com/vikastaank/imagescdn/resources/images/
	IMAGE_DISK_CACHE_PATH = "/home/vikastaank/ImageCDN/cache/"

	// err msgs
	BASE64_DECODE_ERR          = "illegal base64 data"
	INVALID_URL_IDENTIFIER_KEY = "There is some issue with passed url identifier key, please check"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Printf(".env not found")
	} else {
		log.Printf(".env loaded")
	}

	if os.Getenv("IMAGE_DISK_CACHE_PATH") != "" {
		IMAGE_DISK_CACHE_PATH = os.Getenv("IMAGE_DISK_CACHE_PATH")
	}
}
