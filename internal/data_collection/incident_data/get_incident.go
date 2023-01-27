package incident

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"sort"
)

// StartIncidentService - функция отправляет запрос, возвращает данные, в случае ошибки, возвращает пустую строку
func StartIncidentService() ([]IncidentData, error) {
	resp, err := http.Get(viper.GetString("request.incident"))
	if err != nil {
		log.Error().Err(err).Msg("Ошибка при отправке GET-запроса к /accendent")
		return []IncidentData{}, nil
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Error().Int("statusCode", resp.StatusCode).Msg("Получен код состояния, отличный от 200")
		return []IncidentData{}, nil
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка чтения ответа")
		return []IncidentData{}, nil
	}

	var incidentData []IncidentData
	err = json.Unmarshal(body, &incidentData)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка демаршалинга JSON в структуру IncidentData")
		return []IncidentData{}, nil
	}

	sort.SliceStable(incidentData, func(i, j int) bool {
		return incidentData[i].Status < incidentData[j].Status
	})

	return incidentData, nil
}
