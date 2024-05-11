package random

import "crypto/rand"

func generateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	for i := range randomBytes {
		randomBytes[i] = charset[int(randomBytes[i])%len(charset)]
	}
	return string(randomBytes), nil
}

func GenerateHex64() (string, error) {
	return generateRandomString(64)
}
