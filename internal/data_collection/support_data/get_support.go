package support

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

// StartSupportService - Функция запускает сервис для получения данных о системе Support. Результат выполения функциия -
// []SupportData, либо ошибка.
func StartSupportService() ([]int, error) {
	data, err := GetSupportData()
	if err != nil {
		return []int{0, 0}, err
	}
	return validSupportData(data), nil
}

// GetSupportData - функция отправляет запрос, возвращает данные, в случае ошибки, возвращает пустую строку
func GetSupportData() ([]SupportData, error) {
	resp, err := http.Get(viper.GetString("request.support"))
	if err != nil {
		log.Error().Err(err).Msg("Ошибка при отправке GET-запроса к /support")
		return []SupportData{}, nil
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Error().Int("statusCode", resp.StatusCode).Msg("Получен код состояния, отличный от 200")
		return []SupportData{}, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка чтения ответа")
		return []SupportData{}, nil
	}

	var supportData []SupportData
	err = json.Unmarshal(body, &supportData)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка демаршалинга JSON в структуру SupportData")
		return []SupportData{}, nil
	}

	return supportData, nil
}

// validSupportData - Функция валидирует данные о состоянии системы Support. На вход принимаем []SupportData, результат
// выполнения - []int. Срез из двух int, первый из которых показывает загруженность службы поддержки (1–3),
// а второй — среднее время ожидания ответа.
func validSupportData(data []SupportData) []int {
	result := make([]int, 0)
	var totalTopic, load, averageTime int

	for _, item := range data {
		totalTopic += item.ActiveTickets
	}

	switch {
	case totalTopic < 9:
		load = 1
	case totalTopic <= 16:
		load = 2
	default:
		load = 3
	}

	averageTime = int((float64(60) / float64(18)) * float64(totalTopic))
	result = append(result, load)
	result = append(result, averageTime)

	return result

}
