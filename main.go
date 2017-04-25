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
	"github.com/zling/zi-goproject/selectors"
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

func convertToMap(records []*formats.Record) map[string]*formats.Record {
	result := make(map[string]*formats.Record)
	for _, record := range records {
		result[record.CompanyId] = record
	}

	return result
}

func joinCompanyData(records []*formats.Record, ceoList formats.MappedCompanyRecords) []*formats.CompanyData {
	var results []*formats.CompanyData
	companyMap := convertToMap(records)

	for _, company := range ceoList {
		companyId := company.CompanyId
		value, ok := companyMap[companyId]
		if !ok {
			fmt.Fprintf(os.Stderr, "company is not on the list.")
			os.Exit(1)
		} else {
			entry := &formats.CompanyData{
				CompanyId:        companyId,
				CompanyUrl:       value.CompanyUrl,
				CompanyNameLong:  value.CompanyNameLong,
				CompanyNameShort: value.CompanyNameShort,
				CompanyLogoSmall: value.CompanyLogoSmall,
				CeoName:          value.CeoName,
				CeoPhoto:         value.CeoPhoto,
				NumberOfVotes:    company.NumberOfVotes,
				Rating:           company.Rating,
				City:             company.City,
				Sector:           company.Sector,
				IsWinner:         company.IsWinner,
			}

			results = append(results, entry)
		}
	}

	return results
}

func main() {
	records := data.GetCompanyRawRecords("files/company_list.txt")
	cityMapped, unMappedCities := mappers.MapCities(records)
	sectorMapped, unMappedSectors := mappers.MapSectors(records)
	result := mappers.CombineCityAndSectorMapped(cityMapped, sectorMapped)
	ceoList := getCeoRatingAndNumberOfVotes(records, result)
	sort.Sort(ceoList)
	ceoList = selectors.TagWinners(ceoList)
	finalOutput := joinCompanyData(records, ceoList)

	// final output data
	fileHandle, _ := os.Create("files/finalOutputData.txt")
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	for _, value := range finalOutput {
		str := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\t%d\t%s\t%s\t%t", value.CompanyId, value.CompanyUrl, value.CompanyNameLong, value.CompanyNameShort, value.CompanyLogoSmall, value.CeoName, value.CeoPhoto, value.NumberOfVotes, value.Rating, value.City, strings.Join(value.Sector, "|"), value.IsWinner)
		fmt.Fprintln(writer, str)
		writer.Flush()
	}

	// output data
	fileHandle, _ = os.Create("files/taggedCompanies.txt")
	writer = bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	for _, value := range ceoList {
		str := fmt.Sprintf("%s\t%t\t%d\t%d\t%s\t%s", value.CompanyId, value.IsWinner, value.Rating, value.NumberOfVotes, value.City, strings.Join(value.Sector, "|"))
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

	// fileHandle, _ := os.Create("files/output.txt")
	// writer := bufio.NewWriter(fileHandle)
	// defer fileHandle.Close()

	// for _, value := range result {
	// 	str := fmt.Sprintf("%s\t%s\t%s", value.CompanyId, value.City, strings.Join(value.Sector, "|"))
	// 	fmt.Fprintln(writer, str)
	// 	writer.Flush()
	// }

	fmt.Println(len(cityMapped))
	fmt.Println(len(sectorMapped))
	fmt.Println(len(result))
}
