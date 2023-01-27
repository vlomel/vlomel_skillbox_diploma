package result

import (
	billing "github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/billing_data"
	email "github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/email_data"
	incident "github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/incident_data"
	mms "github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/mms_data"
	sms "github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/sms_data"
	voice "github.com/vlomel/vlomel_skillbox_diploma/internal/data_collection/voice_data"
)

// ResultT - Результирующая структура с данными для передачи по HTTP
type ResultT struct {
	Status bool        `json:"status"`
	Data   *ResultSetT `json:"data"`
	Error  string      `json:"error"`
}

// ResultSetT - Структура для хранения данных со всех сервисов
type ResultSetT struct {
	SMS       [][]sms.SMSData                `json:"sms"`
	MMS       [][]mms.MMSData                `json:"mms"`
	VoiceCall []voice.VoiceData              `json:"voice_call"`
	Email     map[string][][]email.EmailData `json:"email"`
	Billing   billing.BillingData            `json:"billing"`
	Support   []int                          `json:"support"`
	Incidents []incident.IncidentData        `json:"incident"`
}
