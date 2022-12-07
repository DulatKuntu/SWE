package utils

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	salt             = "hjqrhjqw124617ajfhajs"
	signingKey       = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL         = 12 * time.Hour
	boymanLocalToken = "ba2228a00d21e19c23e4f210a5b8a300"
)

type tokenClaims struct {
	jwt.StandardClaims
	ID   string `json:"id"`
	Role string `json:"role"`
}

func CreateItemImage(c *gin.Context) (string, error) {
	// Image handling
	c.Request.ParseMultipartForm(10 << 20)

	fileName := ""
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, _, err := c.Request.FormFile("image")
	if file != nil {
		defer file.Close()
	}

	if file != nil {
		defer file.Close()
	}

	locationImage, exists := os.LookupEnv("LocationItemDocker")

	if !exists {
		return "", errors.New("enviroment variable is not set")
	}

	if err != nil {
		if err.Error() != "http: no such file" {

			return "", err
		}
	} else {
		// Create a temporary file within our temp-images directory that follows
		tempFile, err := ioutil.TempFile(locationImage, "upload-*.jpeg")
		if err != nil {
			return "", err
		}
		defer tempFile.Close()
		// read all of the contents of our uploaded file into a
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			return "", err
		}
		tempFile.Write(fileBytes)

		fileName = filepath.Base(tempFile.Name())
	}
	c.Request.ParseMultipartForm(0)

	return fileName, nil
}

func GenerateToken(id string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
		},
		id,
		role,
	})
	log.Print(token)
	return token.SignedString([]byte(signingKey))
}

func ParseToken(accessToken string) (string, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	log.Print(token)
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", "", errors.New("token claims are not of type *tokenClaims")
	}
	log.Print(claims)

	return claims.ID, claims.Role, nil
}
