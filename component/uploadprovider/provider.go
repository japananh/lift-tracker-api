package uploadprovider

import (
	"context"
	"lift-tracker-api/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
