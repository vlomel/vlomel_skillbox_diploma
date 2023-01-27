package internal

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/vlomel/vlomel_skillbox_diploma/internal/server"
	"os"
	"os/signal"
	"syscall"
)

func RunService() {
	log.Info().Msg("Запуск сервера...")
	server := new(server.Server)
	go func() {
		if err := server.RunServer(viper.GetString("server.port")); err != nil {
			log.Err(err).Msg("Сервер не запущен")
		}
	}()
	log.Info().Msg("Сервер запущен")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	log.Warn().Msg("Остановка сервера...")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Err(err).Msg("Некорректное закрытие сервера")
	}

	log.Info().Msg("Удачи!")
}
