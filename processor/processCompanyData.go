package processor

import (
	"bufio"
	"encoding/json"
	"os"
	"sync"

	log "github.com/golang/glog"
	"github.com/zling/zi-goproject/formats"
	"github.com/zling/zi-goproject/mappers"
)

func ReadCompanyRawDataJsonFile(path string) formats.CompanyDataJsonRecords {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	companyRecords := make(chan string)
	results := make(chan *formats.CompanyDataJson)

	wg := new(sync.WaitGroup)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go marshalAndMapToCityAndIndustry(companyRecords, results, wg)
	}

	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			companyRecords <- scanner.Text()
		}
		close(companyRecords)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var companyRecordArr formats.CompanyDataJsonRecords
	for v := range results {
		companyRecordArr = append(companyRecordArr, v)
	}

	return companyRecordArr
}

func addWhenNonEmpty(result []string, elem string) []string {
	if elem != "" {
		result = append(result, elem)
		return result
	}

	return result
}

func marshalAndMapToCityAndIndustry(companyRecords <-chan string, results chan<- *formats.CompanyDataJson, wg *sync.WaitGroup) {

	defer wg.Done()

	// eventually I want to have a []string channel to work on a chunk of lines not just one line of text
	for record := range companyRecords {

		var companyData formats.CompanyDataJson
		err := json.Unmarshal([]byte(record), &companyData)
		if err != nil {
			log.Error(err)
		}

		var sectors []string
		sectors = addWhenNonEmpty(sectors, companyData.IndustryName)
		sectors = addWhenNonEmpty(sectors, companyData.Sector1)
		sectors = addWhenNonEmpty(sectors, companyData.Sector2)
		sectors = addWhenNonEmpty(sectors, companyData.Sector3)
		sectors = addWhenNonEmpty(sectors, companyData.Sector4)
		sectors = addWhenNonEmpty(sectors, companyData.Sector5)
		companyData.MappedSectors = mappers.MapSector(sectors, companyData.CompanyId)
		companyData.MappedSegments = mappers.MapSegments(companyData.YearFounded, companyData.CompanyRevenue)

		if companyData.CityName != "" {
			companyData.MappedCity, err = mappers.MapCity(companyData.CityName)
			if err != nil {
				log.Warning(err)
			}
		}

		results <- &companyData
	}
}
