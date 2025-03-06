package util

import (
	"crypto/rand"
	"embed"
	"fmt"
	"math/big"
	"strings"
)

//go:embed english.txt
var f embed.FS

var words []string

func loadWords(wordlist string) error {
	data, err := f.ReadFile(wordlist)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	words = strings.Split(strings.ReplaceAll(string(data), "\r", ""), "\n")
	return nil
}

func GenerateMnemonic(length int) (string, error) {
	// If we need to, load words
	if len(words) == 0 {
		err := loadWords("english.txt") // TODO: Get file name from .env to allow for different languages
		if err != nil {
			return "", err
		}
	}

	// Fetch random words
	mnemonic := make([]string, length)
	for i := 0; i < length; i++ {
		// Generate random number
		rand_num, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random number: %v", err)
		}

		mnemonic[i] = words[rand_num.Int64()]
	}

	return strings.Join(mnemonic, " "), nil
}
