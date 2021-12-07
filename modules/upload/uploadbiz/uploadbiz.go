package uploadbiz

import (
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"lift-tracker-api/common"
	"lift-tracker-api/component/uploadprovider"
	"lift-tracker-api/modules/upload/uploadmodel"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStore interface {
	CreateImage(ctx context.Context, data *common.Image) error
}

type uploadbiz struct {
	provider uploadprovider.UploadProvider
	imgStore CreateImageStore
}

func NewUploadBiz(provider uploadprovider.UploadProvider, imgStore CreateImageStore) *uploadbiz {
	return &uploadbiz{provider: provider, imgStore: imgStore}
}

func (biz *uploadbiz) Upload(
	ctx context.Context,
	data []byte,
	folder,
	fileName string,
) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)
	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)                                // "image.jpg" -> ".jpg"
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt) // 11413435463.jpg

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))
	if err != nil {
		return nil, err
	}

	img.Width = w
	img.Height = h
	img.CloudName = "s3" // should be set in the provider
	img.Extension = fileExt

	// TODO: not sure we need image module?
	// if err := biz.imgStore.CreateImage(ctx, img); err != nil {
	// 	return nil, uploadmodel.ErrCannotSaveFile(err)
	// }

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err when get image dimension:", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
