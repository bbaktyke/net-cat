package server

func isValidtext(text string) bool {
	if text == "" {
		return false
	}

	for _, simbol := range text {
		if simbol < 32 || simbol > 127 {
			return false
		}
	}
	return true
}
