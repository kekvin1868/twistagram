package app

import (
	commentRouter "twistagram/src/modules/comment/routing"
	likeRouter "twistagram/src/modules/like/routing"
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
}
