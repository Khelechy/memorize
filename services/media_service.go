package services

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

const (
	publicStaticMediaRouteUrl = "http://localhost:3000/media/uploads"
	publicStaticMediaRoute = "media/uploads"
	fileUploadPath         = "./uploads"
	baseUrl                = "http://memorize.com/memory"
)

func HandleMediaUpload(userId string, file *multipart.FileHeader, c fiber.Ctx) error {

	uniqueId := uuid.New()

	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	newFile := fmt.Sprintf("%s.%s", filename, fileExt)

	bucketPath := fmt.Sprintf("%s/%s", fileUploadPath, userId)
	destination := fmt.Sprintf("%s/%s", bucketPath, newFile)

	//Create bucket if not exist
	_ = os.MkdirAll(bucketPath, os.ModePerm)

	if err := c.SaveFile(file, destination); err != nil {
		return err
	}

	return nil
}

func SaveUserQr(userId string) (string, error) {

	fileName := fmt.Sprintf("%s-qr.png", userId)

	bucketPath := fmt.Sprintf("%s/%s", fileUploadPath, userId)
	destination := fmt.Sprintf("%s/%s", bucketPath, fileName)

	if _, err := os.Stat(bucketPath); err == nil {
		return fmt.Sprintf("%s/%s/%s", publicStaticMediaRouteUrl, userId, fileName), nil
	}

	url := fmt.Sprintf("%s/%s", baseUrl, userId)
	qrCode, _ := qrcode.New(url, qrcode.Medium)

	//Create bucket if not exist
	_ = os.MkdirAll(bucketPath, os.ModePerm)

	err := qrCode.WriteFile(256, destination)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s/%s", publicStaticMediaRouteUrl, userId, fileName), nil

}

func FetchUserUploads(userId string) ([]string, error) {

	bucketPath := fmt.Sprintf("%s/%s", fileUploadPath, userId)

	files, err := os.ReadDir(bucketPath)
	if err != nil {
		return nil, err
	}

	var mediaLocation []string

	for _, file := range files {
		mediaLocation = append(mediaLocation, fmt.Sprintf("%s/%s/%s", publicStaticMediaRouteUrl, userId, file.Name()))
	}

	return mediaLocation, nil
}
