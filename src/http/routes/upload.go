package v1routes

import (
	uploadController "projectsprintw4/src/http/controllers/upload"
	middleware "projectsprintw4/src/http/middlewares"
)

func (i *V1Routes) MountUpload() {

	upload := uploadController.New(&uploadController.V1Upload{
		DB: i.DB,
	})
	i.Echo.POST("/image", upload.UploadImage, middleware.Authentication())

}
