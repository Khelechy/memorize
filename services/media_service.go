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

func HandleMediaUpload(userId string, file *multipart.FileHeader, c fiber.Ctx) error {

	uniqueId := uuid.New()

	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	newFile := fmt.Sprintf("%s.%s", filename, fileExt)

	bucketPath := fmt.Sprintf("./uploads/%s", userId)
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

	bucketPath := fmt.Sprintf("./uploads/%s", userId)
	destination := fmt.Sprintf("%s/%s", bucketPath, fileName)

	if _, err := os.Stat("sample.txt"); err == nil {
		return fmt.Sprintf("http://localhost:3000/media/uploads/%s/%s", userId, fileName), nil
	}

	url := fmt.Sprintf("http://memorize.com/memory/%s", userId)
	qrCode, _ := qrcode.New(url, qrcode.Medium)

	//Create bucket if not exist
	_ = os.MkdirAll(bucketPath, os.ModePerm)

	err := qrCode.WriteFile(256, destination)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://localhost:3000/media/uploads/%s/%s", userId, fileName), nil

}

func FetchUserUploads(userId string) ([]string, error) {

	bucketPath := fmt.Sprintf("./uploads/%s", userId)

	files, err := os.ReadDir(bucketPath)
	if err != nil {
		return nil, err
	}

	var mediaLocation []string

	for _, file := range files {
		mediaLocation = append(mediaLocation, fmt.Sprintf("media/uploads/%s/%s", userId, file.Name()))
	}

	return mediaLocation, nil
}
