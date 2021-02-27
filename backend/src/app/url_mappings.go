package app

import userRouter "twistagram/src/modules/regristration/routing"

func MapURLs() {
	userRouter.MapURLs(router)
}
