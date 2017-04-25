package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/zling/zi-goproject/data"
	"github.com/zling/zi-goproject/mappers"
)

func main() {
	records := data.GetCompanyRawRecords("files/company_list.txt")

	cityMapped, unMappedCities := mappers.MapCities(records)

	sectorMapped, unMappedSectors := mappers.MapSectors(records)

	result := mappers.CombineCityAndSectorMapped(cityMapped, sectorMapped)

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
