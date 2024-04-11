package services

import (
	"mime/multipart"
	"fmt"
	"strings"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func HandleMediaUpload(userId string, file *multipart.FileHeader, c fiber.Ctx) error{

	uniqueId := uuid.New()
 
  	filename := strings.Replace(uniqueId.String(), "-", "", -1)
	fileExt := strings.Split(file.Filename, ".")[1]
	newFile := fmt.Sprintf("%s.%s", filename, fileExt)

	// Create file bucket if not exist
	bucketPath := fmt.Sprintf("./uploads/%s", userId)
	
	_= os.MkdirAll(bucketPath, os.ModePerm)

	destination := fmt.Sprintf("%s/%s", bucketPath, newFile)

	if err := c.SaveFile(file, destination); err != nil {
		return err
	}

	return nil
}

func FetchUserUploads(userId string) ([]string, error){

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