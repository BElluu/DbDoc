package config

import (
	"bufio"
	"os"
	"strings"
)

func LoadConfig(filename string) (map[string]string, error) {

	//CODE MODE
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//ON REALISE MODE
	// exePath, err := os.Executable()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return nil, err
	// }

	// file, err := os.Open(filepath.Dir(exePath) + "\\" + filename)
	// if err != nil {
	// 	return nil, err
	// }
	// defer file.Close()

	config := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // Skip lines that don't have key=value format
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		config[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}
