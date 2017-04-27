package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/zling/zi-goproject/aggregators"
	"github.com/zling/zi-goproject/data"
	"github.com/zling/zi-goproject/mappers"
	"github.com/zling/zi-goproject/selectors"
	"github.com/zling/zi-goproject/uploaders"
)

func main() {
	records := data.GetCompanyRawRecords("files/company_list.txt")
	cityMapped, unMappedCities := mappers.MapCities(records)
	sectorMapped, unMappedSectors := mappers.MapSectors(records)
	result := mappers.CombineCityAndSectorMapped(cityMapped, sectorMapped)
	ceoList := aggregators.GetCeoRatingAndNumberOfVotes(records, result)
	sort.Sort(ceoList)
	ceoList = selectors.TagWinners(ceoList)
	finalOutput := aggregators.JoinCompanyData(records, ceoList)

	uploaders.InsertToMongo(finalOutput)

	// final output data
	fileHandle, _ := os.Create("files/finalOutputData.txt")
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	for _, value := range finalOutput {
		str := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\t%d\t%s\t%s\t%t\t%t", value.CompanyId, value.CompanyUrl, value.CompanyNameLong, value.CompanyNameShort, value.CompanyLogoSmall, value.CeoName, value.CeoPhoto, value.NumberOfVotes, value.Rating, value.City, strings.Join(value.Sector, "|"), value.IsCityWinner, value.IsIndustryWinner)
		fmt.Fprintln(writer, str)
		writer.Flush()
	}

	// output data
	fileHandle, _ = os.Create("files/taggedCompanies.txt")
	writer = bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	for _, value := range ceoList {
		str := fmt.Sprintf("%s\t%t\t%t\t%d\t%d\t%s\t%s", value.CompanyId, value.IsCityWinner, value.IsIndustryWinner, value.Rating, value.NumberOfVotes, value.City, strings.Join(value.Sector, "|"))
		fmt.Fprintln(writer, str)
		writer.Flush()
	}

	err := ioutil.WriteFile("files/unMappedCities.txt", []byte(strings.Join(unMappedCities, ",")), 0644)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("files/unMappedSectors.txt", []byte(strings.Join(unMappedSectors, "\t")), 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(cityMapped))
	fmt.Println(len(sectorMapped))
	fmt.Println(len(result))
}
