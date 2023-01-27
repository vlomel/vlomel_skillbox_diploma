package voice

import (
	"bufio"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/common/country"
	"github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/common/providers"
	"os"
	"strconv"
	"strings"
)

// StartVoiceService - Функция запускает сервис для получения данных о состоянии системы VoiceCall из файла формата data.
// Данные считиваются и результат выполениния - []VoiceData.
func StartVoiceService() []VoiceData {
	return ReadVoiceFile()
}

func ReadVoiceFile() []VoiceData {
	// Чтение из файла
	contents, err := os.ReadFile(viper.GetString("data.voice"))
	if err != nil {
		log.Error().Err(err).Msg("Ошибка чтения файла")
		return nil
	}

	// Преобразование содержимого файла в строку
	data := string(contents)

	// Создание среза для хранения данных
	var voiceData []VoiceData

	// Используем новый сканер для чтения файла построчно
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		// Получение строки
		line := scanner.Text()

		// Деление строки по разделителю
		fields := strings.Split(line, ";")

		// Проверка на правильное количество полей
		if len(fields) != 8 {
			continue
		}

		// Проверка кода страны и его подмена на полное название страны
		if !country.CountryCode(fields[0]) {
			continue
		}

		// Проверка провайдеров
		if !providers.VoiceProviders(fields[3]) {
			continue
		}

		connectionStability, err := strconv.ParseFloat(fields[4], 64)
		if err != nil {
			log.Error().Err(err).Msg("Ошибка при преобразовании во float64")
			continue
		}

		// Преобразование строк в int
		bandwidth, err := strconv.Atoi(fields[1])
		if err != nil {
			log.Error().Err(err).Msg("Ошибка при преобразовании в int")
			continue
		}

		responseTime, err := strconv.Atoi(fields[2])
		if err != nil {
			log.Error().Err(err).Msg("Ошибка при преобразовании в int")
			continue
		}

		ttfb, err := strconv.Atoi(fields[5])
		if err != nil {
			log.Error().Err(err).Msg("Ошибка при преобразовании в int")
			continue
		}

		voicePurity, err := strconv.Atoi(fields[6])
		if err != nil {
			log.Error().Err(err).Msg("Ошибка при преобразовании в int")
			continue
		}

		medianOfCallsDuration, err := strconv.Atoi(fields[7])
		if err != nil {
			log.Error().Err(err).Msg("Ошибка при преобразовании в int")
			continue
		}

		// Создание новой структуры и добавление ее в срез.
		voiceData = append(voiceData, VoiceData{
			Country:               country.GetCountryName(fields[0]),
			Bandwidth:             bandwidth,
			ResponseTime:          responseTime,
			Provider:              fields[3],
			ConnectionStability:   connectionStability,
			TTFB:                  ttfb,
			VoicePurity:           voicePurity,
			MedianOfCallsDuration: medianOfCallsDuration,
		})
	}
	return voiceData
}
