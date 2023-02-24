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

type uploadImageService struct {
	store    CreateImageStore
	provider uploadprovider.UploadProvider
}

func NewUploadImageService(provider uploadprovider.UploadProvider, store CreateImageStore) *uploadImageService {
	return &uploadImageService{
		provider: provider,
		store:    store,
	}
}
func (service uploadImageService) UpLoadImage(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {

	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploaderrors.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                                // read extension file e.g: "img.jpg" => ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // uniquely name the file by getting the current time in nanosecond

	// s3 provider
	img, err := service.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	img.Width = w
	img.Height = h
	img.Extension = fileExt
	
	//if err := biz.imgStore.CreateImage(ctx, img); err != nil {
	//	// delete img on S3
	//	return nil, uploadmodel.ErrCannotSaveFile(err)
	//}
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
