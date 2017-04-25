package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/zling/zi-goproject/data"
	"github.com/zling/zi-goproject/formats"
	"github.com/zling/zi-goproject/mappers"
)

func getCeoRatingAndNumberOfVotes(records []*formats.Record, resultMap map[string]*formats.MappedCompanyRecord) formats.MappedCompanyRecords {
	var ceoList formats.MappedCompanyRecords

	for _, record := range records {
		companyId := record.CompanyId
		rating, err := strconv.Atoi(record.CeoRating)
		if err != nil {
			panic(err)
		}
		numberOfVotes, err := strconv.Atoi(record.TotalNumberOfCeoRatings)
		if err != nil {
			panic(err)
		}
		value, ok := resultMap[companyId]
		if !ok {
			continue
		} else {
			value.Rating = rating
			value.NumberOfVotes = numberOfVotes
			ceoList = append(ceoList, value)
		}
	}

	return ceoList
}

func main() {
	records := data.GetCompanyRawRecords("files/company_list.txt")

	cityMapped, unMappedCities := mappers.MapCities(records)

	sectorMapped, unMappedSectors := mappers.MapSectors(records)

	result := mappers.CombineCityAndSectorMapped(cityMapped, sectorMapped)

	ceoList := getCeoRatingAndNumberOfVotes(records, result)

	sort.Sort(ceoList)

	for _, ceo := range ceoList {
		fmt.Println("%d : %d", ceo.Rating, ceo.NumberOfVotes)
	}

	fmt.Println(len(cityMapped))

	fmt.Println(len(sectorMapped))

	fmt.Println(len(result))

	err := ioutil.WriteFile("files/unMappedCities.txt", []byte(strings.Join(unMappedCities, ",")), 0644)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("files/unMappedSectors.txt", []byte(strings.Join(unMappedSectors, "\t")), 0644)
	if err != nil {
		panic(err)
	}

	fileHandle, _ := os.Create("files/output.txt")
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	for _, value := range result {
		str := fmt.Sprintf("%s\t%s\t%s", value.CompanyId, value.City, strings.Join(value.Sector, "|"))
		fmt.Fprintln(writer, str)
		writer.Flush()
	}
}
