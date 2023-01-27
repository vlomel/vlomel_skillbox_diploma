package result

import (
	billing "github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/billing_data"
	email "github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/email_data"
	incident "github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/incident_data"
	mms "github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/mms_data"
	sms "github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/sms_data"
	support "github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/support_data"
	voice "github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/voice_data"
	"strings"
)

// GetResultData - Функция собирает данные со всех сервисов и возвращает результат в виде структуры ResultT
func GetResultData() ResultT {
	sms := sms.StartSmsService()
	mms, errMMS := mms.StartMMSService()
	voice := voice.StartVoiceService()
	email := email.StartEmailService()
	billing, errBilling := billing.StartBillingService()
	support, errSupport := support.StartSupportService()
	incident, errIncident := incident.StartIncidentService()

	if errMMS != nil || errSupport != nil || errIncident != nil || errBilling != nil {
		return ResultT{false,
			nil,
			errorToString(errMMS, errSupport, errIncident, errBilling),
		}
	}

	return ResultT{true, &ResultSetT{
		sms,
		mms,
		voice,
		email,
		billing,
		support,
		incident},
		"",
	}
}

// errorToString - функция объединяет нескольких сообщений об ошибках в одну строку.
func errorToString(err ...error) string {
	var errorString string
	for _, item := range err {
		if item != nil {
			errorString += item.Error() + ", "
		}
	}
	return strings.TrimRight(errorString, ", ")
}
