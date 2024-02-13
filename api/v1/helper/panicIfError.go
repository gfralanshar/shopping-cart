package helper

import "github.com/sirupsen/logrus"

func PanicIfError(err error) {
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}
}
