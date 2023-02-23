package uploadservice

import (
	"bytes"
	"context"
	"fmt"
	"golang-realworld/common"
	"golang-realworld/component/uploadprovider"
	"golang-realworld/modules/upload/uploaderrors"
	"image"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStore interface {
	CreateImage(context context.Context, data *common.Image) error
}

type createImageService struct {
	store    CreateImageStore
	provider uploadprovider.UploadProvider
}

func NewUploadImageService(store CreateImageStore) *createImageService {
	return &createImageService{
		store: store,
	}
}
func (service createImageService) UpLoadImage(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {

	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploaderrors.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                                // "img.jpg" => ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 9129324893248.jpg

	img, err := service.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	img.Width = w
	img.Height = h

	img.Extension = fileExt
	return img, nil
}
func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
