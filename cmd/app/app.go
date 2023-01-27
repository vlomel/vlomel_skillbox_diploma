package main

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	service "github.com/vlomel/vlomel_skillbox_diploma/internal"
	"github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/common/logging"
)

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func main() {
	logging.InitLogging()
	log.Info().Msg("Запуск приложения")

	if err := initConfig(); err != nil {
		log.Err(err).Msg("Файл конфигурации не найден")
	}

	service.RunService()
}
