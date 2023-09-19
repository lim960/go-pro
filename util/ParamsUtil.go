package util

func HaveBlank(params ...string) bool {
	for _, item := range params {
		if len(item) == 0 {
			return true
		}
	}
	return false
}

func HaveZero(params ...uint) bool {
	for _, item := range params {
		if item <= 0 {
			return true
		}
	}
	return false
}

func AllBlank(params ...string) bool {
	for _, item := range params {
		if len(item) > 0 {
			return false
		}
	}
	return true
}
