package handlers

import (
	"fmt"
	"imagescdn/constants"
	"imagescdn/logger"
	"imagescdn/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var Logger = logger.GetLogger()

func GetImageData(c *gin.Context) {
	urlKey := c.DefaultQuery("url", "")
	fmt.Println(urlKey)
	if urlKey == "" {
		Logger.Error("image identifier key not received in query string params")
		c.JSON(http.StatusBadRequest, gin.H{"msg": constants.BAD_REQUEST_ERR})
		return
	}

	contextLogger := Logger.WithFields(logrus.Fields{
		"ImgEncodedKey": urlKey,
	})

	decodedImgKey, err := services.DecodeImgUrlKey(urlKey, contextLogger)
	if err != nil {
		contextLogger.Errorf("error occurred while decoding the img key: %s", err.Error())
		if strings.Contains(err.Error(), constants.BASE64_DECODE_ERR) {
			// illegal base64 data at input byte
			c.JSON(http.StatusBadRequest, gin.H{"msg": constants.INVALID_URL_IDENTIFIER_KEY})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"msg": constants.INTERNAL_SERVER_ERR})
		return
	}
	contextLogger.Infof("Image key received in query string params decoded as : %s", string(decodedImgKey))

	var imgRawData []byte
	imgRawData = services.CheckInMemoryCache(urlKey)
	if imgRawData != nil {
		contextLogger.Info("returning data from in-memory cache")
		c.Data(http.StatusOK, "image/jpeg", imgRawData)
	} else {
		exists := services.CheckInDiskCache(string(decodedImgKey), contextLogger)
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"msg": constants.IMG_NOT_FOUND_ERR})
			return
		}

		imgRawData, err = services.GetImgRawDataFromDiskCache(string(decodedImgKey), contextLogger)
		if err != nil {
			contextLogger.Errorf("error occurred while getting raw data from disk-cache: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"msg": constants.INTERNAL_SERVER_ERR})
		} else {
			contextLogger.Info("returning data from disk cache")
			services.AddInMemoryCache(urlKey, imgRawData, contextLogger)
			c.Data(http.StatusOK, "image/jpeg", imgRawData)
		}
	}
	return
}

