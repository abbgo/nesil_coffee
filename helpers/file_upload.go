package helpers

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FileUpload(fileName, path string, context *gin.Context) (string, error) {
	file, err := context.FormFile(fileName)
	if err != nil {
		return "", err
	}

	extensionFile := filepath.Ext(file.Filename)

	if extensionFile != ".jpg" && extensionFile != ".JPG" {
		return "", errors.New("the image must be .jpg format")
	}
	newFileName := uuid.New().String() + extensionFile

	_, err = os.Stat(ServerPath + "uploads/" + path)
	if err != nil {
		if err := os.MkdirAll(ServerPath+"uploads/"+path, os.ModePerm); err != nil {
			return "", err
		}
	}
	if err := context.SaveUploadedFile(file, ServerPath+"uploads/"+path+"/"+newFileName); err != nil {
		return "", err
	}

	return "uploads/" + path + "/" + newFileName, nil
}
