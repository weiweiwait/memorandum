package handler

import "errors"

func PanicTfUserError(err error) {
	if err != nil {
		err = errors.New("User Service--" + err.Error())
		panic(err)
	}
}
