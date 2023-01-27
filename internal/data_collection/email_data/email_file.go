package email

import (
	"bufio"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/common/country"
	"github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/common/providers"
	"os"
	"sort"
	"strconv"
	"strings"
)

// StartEmailService - Функция запускает сервис для получения данных о состоянии системы Email из файла формата data.
// Данные считиваются и происходит сортировка. Результат выполениния - map[string][][]EmailData.
func StartEmailService() map[string][][]EmailData {
	return SortedEmailData(ReadEmailFile())
}

func ReadEmailFile() []EmailData {
	// Чтение из файла
	contents, err := os.ReadFile(viper.GetString("data.email"))
	if err != nil {
		log.Error().Err(err).Msg("Ошибка чтения файла")
		return nil
	}

	// Преобразование содержимого файла в строку
	data := string(contents)

	// Создание среза для хранения данных
	var emailData []EmailData

	// Используем новый сканер для чтения файла построчно
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		// Получение строки
		line := scanner.Text()

		// Деление строки по разделителю
		fields := strings.Split(line, ";")

		// Проверка на правильное количество полей
		if len(fields) != 3 {
			continue
		}

		// Проверка кода страны
		if !country.CountryCode(fields[0]) {
			continue
		}

		// Проверка провайдеров
		if !providers.EmailProviders(fields[1]) {
			continue
		}

		// Преобразование строк в int
		deliveryTime, err := strconv.Atoi(fields[2])
		if err != nil {
			log.Error().Err(err).Msg("Ошибка при преобразовании в int")
			continue
		}

		// Создание новой структуры и добавление ее в срез.
		emailData = append(emailData, EmailData{
			Country:      country.GetCountryName(fields[0]),
			Provider:     fields[1],
			DeliveryTime: deliveryTime,
		})
	}
	return emailData
}

// SortedEmailData - Функция сортирует данные о состоянии системы Email. На вход принимает []EmailData, результат
// выполнения -  map[string][][]EmailData.
func SortedEmailData(emailData []EmailData) map[string][][]EmailData {

	result := make(map[string][][]EmailData)
	//Сортировка по наименованию страны
	sort.SliceStable(emailData, func(i, j int) bool {
		return emailData[i].Country > emailData[j].Country
	})
	//Получить список уникальных стран во входном срезе
	uniqueCountryList := getCountry(emailData)

	for _, country := range uniqueCountryList {
		//Создание среза
		res := make([]EmailData, 0)

		for _, data := range emailData {
			//Добавление страны в срез res, если страны совпадают
			if country == data.Country {
				res = append(res, data)
			}
		}
		//Сортировка среза по полю DeliveryTime в порядке убывания
		sort.SliceStable(res, func(i, j int) bool {
			return res[i].DeliveryTime > res[j].DeliveryTime
		})
		//Получение трёх самых быстрых и трёх самых медленных провайдеров из среза res со страной в качестве ключа.
		result[country] = [][]EmailData{res[:3], res[len(res)-3:]}
	}

	return result
}

// getCountry - функция для получения списка стран из структуры emailData
func getCountry(data []EmailData) []string {
	//Создаем срез для хранения результата
	result := make([]string, 0)

	for _, item := range data {
		//Добавляем страну в срез данных
		result = append(result, item.Country)
	}
	//Удаление дубликатов через функцию uniqueCountry
	return uniqueCountry(result)
}

// uniqueCountry - функция удаляет дубликаты из среза стран
func uniqueCountry(array []string) []string {
	keys := make(map[string]bool)

	result := make([]string, 0)

	//Перебираем входной срез, содержащий наименование стран
	for _, item := range array {
		//Проверяем, есть ли текущая страна в map keys
		if _, value := keys[item]; !value {
			//Если нет, добавляем в страну в конечный срез
			keys[item] = true
			result = append(result, item)
		}
	}

	return result
}
