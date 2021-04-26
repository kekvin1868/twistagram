package app

import (
	bookmarkRouter "twistagram/src/modules/bookmark/routing"
	commentRouter "twistagram/src/modules/comment/routing"
	followRouter "twistagram/src/modules/follow/routing"
	likeRouter "twistagram/src/modules/like/routing"
	photoRouter "twistagram/src/modules/photo/routing"
	postRouter "twistagram/src/modules/post/routing"
	regristrationRouter "twistagram/src/modules/regristration/routing"
	userRouter "twistagram/src/modules/user/routing"
)

func MapURLs() {
	regristrationRouter.MapURLs(router)
	userRouter.MapURLs(router)
	postRouter.MapURLs(router)
	commentRouter.MapURLs(router)
	likeRouter.MapURLs(router)
	bookmarkRouter.MapURLs(router)
	followRouter.MapURLs(router)
	photoRouter.MapURLs(router)
}
