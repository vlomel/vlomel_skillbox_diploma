package mms

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/common/country"
	"github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/common/providers"
	"io"
	"net/http"
	"sort"
)

// StartMMSService - Функция запускает сервис для получения данных о системе MMS. Результат выполения функциия -
// [][]MMSData, либо ошибка
func StartMMSService() ([][]MMSData, error) {
	data, err := GetMMSData()

	if err != nil {
		var res [][]MMSData
		return res, err
	}

	return SortedMMSData(data), nil
}

// GetMMSData - функция отправляет запрос, разбирает полученный ответ и возвращает отфильтрованные данные
func GetMMSData() ([]MMSData, error) {
	// Отправляем GET-запрос
	resp, err := http.Get(viper.GetString("request.mms"))
	if err != nil {
		log.Error().Err(err).Msg("Ошибка при отправке GET-запроса к /mms")
		return nil, err
	}
	defer resp.Body.Close()

	// Проверка кода ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ошибка: %d", resp.StatusCode)
	}

	// Чтение ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка чтения ответа")
		return nil, err
	}

	// Разделение ответа на срез структуры MMSData
	var data []MMSData
	if err := json.Unmarshal(body, &data); err != nil {
		log.Error().Err(err).Msg("Ошибка демаршалинга JSON")
		return nil, err
	}

	filteredData := make([]MMSData, 0)
	for _, d := range data {
		if country.CountryCode(d.Country) && providers.SMSMMSProviders(d.Provider) {
			filteredData = append(filteredData, d)
		}
	}

	return filteredData, nil
}

// SortedMMSData - Функция сортирует данные о состоянии системы MMS. На вход принимаем []MMSData, результат
// выполнения -  срез [][]MMSData. Первый список отсортирован по названию страны от A до Z.
// Второй список отсортирован по названию провайдера от A до Z.
func SortedMMSData(mms []MMSData) [][]MMSData {
	result := make([][]MMSData, 0)
	mmsDataSortedByCountryName := make([]MMSData, 0)
	mmsDataSortedByProviderName := make([]MMSData, 0)

	for _, item := range mms {
		item.Country = country.GetCountryName(item.Country)
		mmsDataSortedByCountryName = append(mmsDataSortedByCountryName, item)
		mmsDataSortedByProviderName = append(mmsDataSortedByProviderName, item)
	}

	sort.SliceStable(mmsDataSortedByCountryName, func(i, j int) bool {
		return mmsDataSortedByCountryName[i].Country < mmsDataSortedByCountryName[j].Country
	})

	sort.SliceStable(mmsDataSortedByProviderName, func(i, j int) bool {
		return mmsDataSortedByProviderName[i].Provider < mmsDataSortedByProviderName[j].Provider
	})

	result = append(result, mmsDataSortedByCountryName)
	result = append(result, mmsDataSortedByProviderName)

	return result
}
