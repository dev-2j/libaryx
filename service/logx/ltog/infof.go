package ltog

import (
	"fmt"

	"example.com/m/service/config"
	"example.com/m/service/logx"
	"github.com/sirupsen/logrus"
)

func Infof(format string, v ...any) {

	if logx.GetLambdaActive() {
		fmt.Println(fmt.Sprintf(format, v...))
		return
	}

	if config.IsShowLogs() {
		logx.LoadInit()
		logrus.Infof(format, v...)
	}
}
