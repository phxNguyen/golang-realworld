package ginupload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-realworld/common"
	"golang-realworld/component"
	"golang-realworld/modules/upload/uploaderrors"
	_ "image/jpeg"
	_ "image/png"
)

func Upload(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		//db := appCtx.GetMainDBConnection()

		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.SaveUploadedFile(fileHeader, fmt.Sprintf("./static/%s", fileHeader.Filename)); err != nil {
			panic(uploaderrors.ErrCannotSaveFile(err))

		}
		//folder := c.DefaultPostForm("folder", "img")
		//
		//file, err := fileHeader.Open()
		//if err != nil {
		//	panic(common.ErrInvalidRequest(err))
		//}
		//
		//defer file.Close() // we can close here
		//
		//dataBytes := make([]byte, fileHeader.Size)
		//if _, err := file.Read(dataBytes); err != nil {
		//	panic(common.ErrInvalidRequest(err))
		//}
		//
		////imgStore := uploadstorage.NewSQLStore(db)
		//service := uploadservice.NewUploadImageService(appCtx.UploadProvider(), nil)
		//img, err := service.UpLoadImage(c.Request.Context(), dataBytes, folder, fileHeader.Filename)
		//if err != nil {
		//	panic(err)
		//}
		c.JSON(200, common.NewShortSuccessResponse(true))
	}
}
