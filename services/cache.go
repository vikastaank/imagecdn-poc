package services

import (
	"encoding/base64"
	"fmt"
	"imagescdn/constants"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
)

var InMemoryImgCache map[string][]byte

func init() {
	InMemoryImgCache = make(map[string][]byte)
}

func DecodeImgUrlKey(encodedStr string, contextLogger *logrus.Entry) ([]byte, error) {
	rawDecodedText, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		contextLogger.Errorf("DecodeImgUrlKey :: error occurred - %s", err.Error())
		return nil, err
	}
	return rawDecodedText, nil
}

func CheckInMemoryCache(imgKey string) []byte {
	return InMemoryImgCache[imgKey]
}

func AddInMemoryCache(imgKey string, rawData []byte, contextLogger *logrus.Entry) {
	contextLogger.Info("AddInMemoryCache:: adding in memory cache for img key")
	InMemoryImgCache[imgKey] = rawData
}

func CheckInDiskCache(imgName string, contextLogger *logrus.Entry) bool {
	imageFullPath := fmt.Sprintf("%s%s", constants.IMAGE_DISK_CACHE_PATH, imgName)
	if _, err := os.Stat(imageFullPath); err != nil {
		contextLogger.Infof("CheckInDiskCache :: File does not exist in disk cache : %s", imgName)
		return false
	}
	return true
}

func GetImgRawDataFromDiskCache(imgName string, contextLogger *logrus.Entry) ([]byte, error) {
	imageFullPath := fmt.Sprintf("%s%s", constants.IMAGE_DISK_CACHE_PATH, imgName)
	fileBytes, err := ioutil.ReadFile(imageFullPath)
	if err != nil {
		contextLogger.Errorf("GetImgRawDataFromDiskCache :: error occurred - %s", err.Error())
		return nil, err
	}
	return fileBytes, nil
}
