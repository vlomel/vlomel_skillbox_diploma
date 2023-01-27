package sms

import (
	"bufio"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/common/country"
	"github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/common/providers"
	"os"
	"sort"
	"strings"
)

// StartSmsService - Функция запускает сервис для получения данных о состоянии системы SMS из файла формата data.
// Данные считиваются и затем происходит сортировка. Результат выполениния - [][]SMSData.
func StartSmsService() [][]SMSData {
	return SortedSMSData(ReadSMSFile())
}

func ReadSMSFile() []SMSData {
	// Чтение из файла
	contents, err := os.ReadFile(viper.GetString("data.sms"))
	if err != nil {
		log.Error().Err(err).Msg("Ошибка чтения файла")
		return nil
	}

	// Преобразование содержимого файла в строку
	data := string(contents)

	// Создание среза для хранения данных
	var smsData []SMSData

	// Используем новый сканер для чтения файла построчно
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		// Получение строки
		line := scanner.Text()

		// Деление строки по разделителю
		fields := strings.Split(line, ";")

		// Проверка на правильное количество полей
		if len(fields) != 4 {
			continue
		}

		// Проверка кода страны и его подмена на полное название страны
		if !country.CountryCode(fields[0]) {
			continue
		}

		// Проверка провайдеров
		if !providers.SMSMMSProviders(fields[3]) {
			continue
		}

		// Создание новой структуры и добавление ее в срез.
		smsData = append(smsData, SMSData{
			Country:      fields[0],
			Bandwidth:    fields[1],
			ResponseTime: fields[2],
			Provider:     fields[3],
		})
	}
	return smsData
}

// SortedSMSData - Функция сортирует данные о состоянии системы SMS. На вход принимает []SMSData, результат
// выполнения -  срез [][]SMSData. Первый список отсортирован по названию провайдера от A до Z.
// Второй список отсортирован по названию страны от A до Z.
func SortedSMSData(sms []SMSData) [][]SMSData {
	result := make([][]SMSData, 0)
	smsDataSortedByCountryName := make([]SMSData, 0)
	smsDataSortedByProviderName := make([]SMSData, 0)

	for _, item := range sms {
		item.Country = country.GetCountryName(item.Country)
		smsDataSortedByCountryName = append(smsDataSortedByCountryName, item)
		smsDataSortedByProviderName = append(smsDataSortedByProviderName, item)
	}

	sort.SliceStable(smsDataSortedByCountryName, func(i, j int) bool {
		return smsDataSortedByCountryName[i].Country < smsDataSortedByCountryName[j].Country
	})

	sort.SliceStable(smsDataSortedByProviderName, func(i, j int) bool {
		return smsDataSortedByProviderName[i].Provider < smsDataSortedByProviderName[j].Provider
	})

	result = append(result, smsDataSortedByCountryName)
	result = append(result, smsDataSortedByProviderName)

	return result
}
