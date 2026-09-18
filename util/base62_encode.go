package util

const base62chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Base62Encode(n int) string {
	if n == 0 {
		return "0"
	}

	result := ""

	for n > 0 {
		result = string(base62chars[n%62]) + result
		n /= 62
	}

	return result
}
