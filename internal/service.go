package internal

import (
	"OrderDelayServing/api/http"
	"OrderDelayServing/internal/config"
	"OrderDelayServing/pkg/model"
	"OrderDelayServing/pkg/repository"
	"OrderDelayServing/utils/datasources"
	"github.com/sirupsen/logrus"
)

func Run() {
	appConfig, err := config.LoadConfig()
	if err != nil {
		logrus.Warnln("error in loading configs")
	}

	db, err := datasources.InitPostgres(appConfig)
	if err != nil {
		logrus.Fatalf("failed to connect database: %v", err)
	}

	err = model.Migrate(db)
	if err != nil {
		logrus.Fatalf("failed to apply migrations: %v", err)
	}

	repos := repository.InitRepos(db)

	http.Run(appConfig, repos)
}
