package data

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func GetNewSectorMapping(fileName string) map[string][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mappings := make(map[string][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "\t")

		value, ok := mappings[arr[1]]
		if !ok {
			mappings[arr[1]] = []string{arr[0]}
		} else {
			for _, sector := range value {
				if arr[0] != sector {
					value = append(value, arr[0])
					break
				}
			}
			mappings[arr[1]] = value
		}

	}

	return mappings
}
