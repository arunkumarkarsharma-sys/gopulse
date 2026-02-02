package input

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func ReadURLsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			urls = append(urls, line)
		}
	}

	if len(urls) == 0 {
		return nil, errors.New("no URLs found in file")
	}

	return urls, nil
}
