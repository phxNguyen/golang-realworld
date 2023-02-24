package ginupload

import (
	"github.com/gin-gonic/gin"
	"golang-realworld/common"
	"golang-realworld/component"
	"golang-realworld/modules/upload/uploadservice"
	_ "image/jpeg"
	_ "image/png"
)

func Upload(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		//db := appCtx.GetMainDbConnection()

		// Get image from header
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err)) //
		}

		folder := c.DefaultPostForm("folder", "img")

		// open file
		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close() //defer để make sure là close file

		dataBytes := make([]byte, fileHeader.Size)      // tạo slice
		if _, err := file.Read(dataBytes); err != nil { // rồi đọc cái slice vừa rồi
			panic(common.ErrInvalidRequest(err))
		}

		//imgStore := uploadstore.UploadStore{}
		service := uploadservice.NewUploadImageService(appCtx.UploadProvider(), nil)
		img, err := service.UpLoadImage(c.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			panic(err)
		}
		c.JSON(200, common.NewShortSuccessResponse(img))
	}
}
