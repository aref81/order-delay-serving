package internal

import (
	"OrderDelayServing/api/http"
	"OrderDelayServing/internal/config"
	"github.com/sirupsen/logrus"
)

func Run() {
	appConfig, err := config.LoadConfig()
	if err != nil {
		logrus.Warnln("error in loading configs")
	}

	http.Run(appConfig)
}
