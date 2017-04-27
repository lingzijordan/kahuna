package data

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	log "github.com/golang/glog"
)

func GetNewSectorMapping(fileName string) map[int][]string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	mappings := make(map[int][]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "\t")
		i, err := strconv.Atoi(arr[1])
		if err != nil {
			log.Error(err)
		}

		value, ok := mappings[i]
		if !ok {
			mappings[i] = []string{arr[0]}
		} else {
			for _, sector := range value {
				if arr[0] != sector {
					value = append(value, arr[0])
					break
				}
			}
			mappings[i] = value
		}

	}

	return mappings
}
