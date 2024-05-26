package v1routes

import (
	uploadController "projectsprintw4/src/http/controllers/upload"
	middleware "projectsprintw4/src/http/middlewares"
)

func (i *V1Routes) MountUpload() {
	g := i.Echo.Group("")

	upload := uploadController.New(&uploadController.V1Upload{
		DB: i.DB,
	})
	g.Use(middleware.Authentication())
	g.POST("/image", upload.UploadImage)

}
