package ltog

import (
	"fmt"

	"example.com/m/service/logx"
	"github.com/sirupsen/logrus"
)

func Infoln(v ...any) {

	if logx.GetLambdaActive() {
		fmt.Println(v...)
		return
	}

	logx.LoadInit()
	logrus.Infoln(v...)

}
