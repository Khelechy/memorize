package services

import (
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	"github.com/khelechy/memorize/helpers"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

const (
	publicStaticMediaRouteUrl = "https://memorize-c59r.onrender.com/media/uploads"
	publicStaticMediaRoute    = "media/uploads"
	fileUploadPath            = "./uploads"
	staticSitePath            = "./views"
	baseUrl                   = "https://memorize-c59r.onrender.com/memory"
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

func SaveUserQr(userId string) (string, string, error) {

	fileName := fmt.Sprintf("%s-qr.png", userId)
	staticSitePublicUrl := fmt.Sprintf("%s/%s/index.html", baseUrl, userId)

	bucketPath := fmt.Sprintf("%s/%s", fileUploadPath, userId)
	destination := fmt.Sprintf("%s/%s", bucketPath, fileName)

	if _, err := os.Stat(bucketPath); err == nil {
		return fmt.Sprintf("%s/%s/%s", publicStaticMediaRouteUrl, userId, fileName), staticSitePublicUrl, nil
	}

	qrCode, _ := qrcode.New(staticSitePublicUrl, qrcode.Medium)

	//Create bucket if not exist
	_ = os.MkdirAll(bucketPath, os.ModePerm)

	err := qrCode.WriteFile(256, destination)
	if err != nil {
		return "", "", err
	}

	qrUrl := fmt.Sprintf("%s/%s/%s", publicStaticMediaRouteUrl, userId, fileName)

	// Generate static web page
	sitePath := fmt.Sprintf("%s/%s", staticSitePath, userId)
	staticSiteDestination := fmt.Sprintf("%s/%s", sitePath, "index.html")

	//Create static site path if not exist
	_ = os.MkdirAll(sitePath, os.ModePerm)

	//Generate static file content
	htmlContent := helpers.GenerateStaticSite(qrUrl, userId)

	err = os.WriteFile(staticSiteDestination, []byte(htmlContent), 0644) //create a new file
	if err != nil {
		return "", "", err
	}

	return qrUrl, staticSitePublicUrl, nil

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
