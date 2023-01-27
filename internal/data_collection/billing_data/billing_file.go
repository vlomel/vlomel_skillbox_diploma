package billing

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

// StartBillingService - Функция запускает сервис для получения данных о системе Billing из файла формата data.
func StartBillingService() (BillingData, error) {
	data, err := os.ReadFile(viper.GetString("data.billing"))
	if err != nil {
		log.Error().Err(err).Msg("Ошибка чтения файла:")
		return BillingData{}, err
	}

	// Преобразование данных в строку с последующим ее делением на срез байтов
	str := string(data)
	bytes := []byte(str)

	// Чтение с правого бита и получение числа
	var mask uint8
	for i, b := range bytes {
		// Проверка на равенство текущего байта 49 (код ASCII для «1»)
		if b == 49 {
			mask += 1 << uint8(i)
		}
	}

	// Проверка каждого бита с помощью логической операции
	return BillingData{
		CreateCustomer: mask&1 == 1,
		Purchase:       mask&2 == 2,
		Payout:         mask&4 == 4,
		Recurring:      mask&8 == 8,
		FraudControl:   mask&16 == 16,
		CheckoutPage:   mask&32 == 32,
	}, nil
}
