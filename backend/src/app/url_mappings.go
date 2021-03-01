package app

import (
	regristrationRouter "twistagram/src/modules/regristration/routing"
	userRouter "twistagram/src/modules/user/routing"
)

func MapURLs() {
	regristrationRouter.MapURLs(router)
	userRouter.MapURLs(router)
}
