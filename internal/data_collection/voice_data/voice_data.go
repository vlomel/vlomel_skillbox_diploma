package voice

type VoiceData struct {
	Country               string  `json:"country"`
	Bandwidth             int     `json:"bandwidth"`
	ResponseTime          int     `json:"response_time"`
	Provider              string  `json:"provider"`
	ConnectionStability   float64 `json:"connection_stability"`
	TTFB                  int     `json:"ttfb"`
	VoicePurity           int     `json:"voice_purity"`
	MedianOfCallsDuration int     `json:"median_of_calls_time"`
}
