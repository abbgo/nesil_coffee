package helpers

import (
	"errors"
	"image"
	"image/jpeg"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/bbrks/go-blurhash"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FileUpload(fileName, path string, context *gin.Context) (string, error) {
	file, err := context.FormFile(fileName)
	if err != nil {
		return "", err
	}

	extensionFile := filepath.Ext(file.Filename)

	if fileName == "image" {
		if extensionFile != ".jpg" && extensionFile != ".JPG" && extensionFile != ".JPEG" && extensionFile != ".jpeg" && extensionFile != ".png" && extensionFile != ".PNG" {
			return "", errors.New("the image must be .jpg or .jpeg format")
		}
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

func BlurHashFileUpload(fileName, path string, context *gin.Context) (string, string, error) {
	file, err := context.FormFile(fileName)
	if err != nil {
		return "", "", err
	}

	extensionFile := filepath.Ext(file.Filename)

	if fileName == "image" {
		if extensionFile != ".jpg" && extensionFile != ".JPG" && extensionFile != ".JPEG" && extensionFile != ".jpeg" && extensionFile != ".png" && extensionFile != ".PNG" {
			return "", "", errors.New("the image must be .jpg or .jpeg format")
		}
	}
	newFileName := uuid.New().String() + extensionFile

	_, err = os.Stat(ServerPath + "uploads/" + path)
	if err != nil {
		if err := os.MkdirAll(ServerPath+"uploads/"+path, os.ModePerm); err != nil {
			return "", "", err
		}
	}
	if err := context.SaveUploadedFile(file, ServerPath+"uploads/"+path+"/"+newFileName); err != nil {
		return "", "", err
	}

	blurHashOfImage, err := GenerateBlurHashImage(file)
	if err != nil {
		return "", "", err
	}

	return "uploads/" + path + "/" + newFileName, blurHashOfImage, nil
}

func GenerateBlurHashImage(file *multipart.FileHeader) (string, error) {
	// Dosyayı aç
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Dosyayı geçici olarak kaydet
	tempFileName := "temp_image.jpg"
	out, err := os.Create(tempFileName)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = out.ReadFrom(src)
	if err != nil {
		return "", err
	}

	// JPEG dosyasını aç ve BlurHash için hazırla
	imgFile, err := os.Open(tempFileName)
	if err != nil {
		return "", err
	}
	defer imgFile.Close()

	img, err := jpeg.Decode(imgFile)
	if err != nil {
		return "", err
	}

	// Görüntüyü RGBA formatına çevir
	rgbaImg := image.NewRGBA(img.Bounds())
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			rgbaImg.Set(x, y, img.At(x, y))
		}
	}

	// BlurHash'i oluştur
	hash, err := blurhash.Encode(4, 3, rgbaImg)
	if err != nil {
		return "", err
	}

	return hash, nil
}
