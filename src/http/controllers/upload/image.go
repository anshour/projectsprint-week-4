package uploadController

import (
	"fmt"
	"net/http"
	"path/filepath"
	"projectsprintw4/src/config"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UploadImageResponse struct {
	ImageUrl string `json:"imageUrl"`
}

func (dbase *V1Upload) UploadImage(c echo.Context) (err error) {
	const minSize int64 = 10 * 1024       // 10KB
	const maxSize int64 = 2 * 1024 * 1024 // 2MB

	contentType := c.Request().Header.Get("Content-Type")
	if !strings.Contains(contentType, "multipart/form-data") {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "Content type must be multipart/form-data",
		})
	}

	// Read form file
	file, err := c.FormFile("file")
	if err != nil {
		if err.Error() == "http: no such file" {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "File is empty",
			})
		}

		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	if file.Size < minSize {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "file size minimum 10KB",
		})
	}
	if file.Size > maxSize {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "file size maximum 2MB",
		})
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: "file must be a .jpg or .jpeg",
		})
	}

	// Source
	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AWS_REGION),
		Credentials: credentials.NewStaticCredentials(
			config.AWS_ACCESS_KEY_ID,
			config.AWS_SECRET_ACCESS_KEY,
			"",
		),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	uploader := s3manager.NewUploader(sess)
	newFilename := uuid.New().String() + ext

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.AWS_BUCKET),
		Key:    aws.String(newFilename),
		Body:   src,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", config.AWS_BUCKET, config.AWS_REGION, newFilename)

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "File uploaded sucessfully",
		Data: UploadImageResponse{
			ImageUrl: fileURL,
		},
	})

}
