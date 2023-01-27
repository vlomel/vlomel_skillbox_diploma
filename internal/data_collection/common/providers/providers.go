package providers

func SMSMMSProviders(code string) bool {
	// Map допустимых провайдеров
	providers := map[string]bool{
		"Topolo": true,
		"Rond":   true,
		"Kildy":  true,
	}

	if _, ok := providers[code]; ok {
		return true
	}
	return false
}

func VoiceProviders(code string) bool {
	// Map допустимых провайдеров
	providers := map[string]bool{
		"TransparentCalls": true,
		"E-Voice":          true,
		"JustPhone":        true,
	}

	if _, ok := providers[code]; ok {
		return true
	}
	return false
}

func EmailProviders(code string) bool {
	// Map допустимых провайдеров
	providers := map[string]bool{
		"Gmail":       true,
		"Yahoo":       true,
		"Hotmail":     true,
		"MSN":         true,
		"Orange":      true,
		"Comcast":     true,
		"AOL":         true,
		"Live":        true,
		"RediffMail":  true,
		"GMX":         true,
		"Proton Mail": true,
		"Yandex":      true,
		"Mail.ru":     true,
	}

	if _, ok := providers[code]; ok {
		return true
	}
	return false
}
