package main

import (
	"flag"
	"sort"

	"github.com/zling/zi-goproject/data"
	"github.com/zling/zi-goproject/formats"
	"github.com/zling/zi-goproject/processor"
	"github.com/zling/zi-goproject/selectors"
	"github.com/zling/zi-goproject/uploaders"
)

var (
	inputFile               string
	outputFile              string
	newSectorCompanyMapping map[int][]string
	newSectorFile           string
)

func init() {
	flag.StringVar(&inputFile, "inputFile", "/Users/owler-zi/Documents/ceo_rating/da-641-2017-04-27.txt", "a path string")
	flag.StringVar(&outputFile, "outputFile", "../../files/outputData.txt", "a path string")
	flag.StringVar(&newSectorFile, "newSector", "../../files/newSector.txt", "a path string")

	flag.Parse()
}

func rankCeos(ceos formats.CompanyDataJsonRecords) {
	cityRankingMap := make(map[string][]int)
	sectorRankingMap := make(map[string][]int)
	for _, ceo := range ceos {
		companyId := ceo.CompanyId
		city := ceo.MappedCity
		sectors := ceo.MappedSectors

		companies, ok := cityRankingMap[city]
		if !ok {
			cityRankingMap[city] = []int{companyId}
		} else {
			companies = append(companies, companyId)
			cityRankingMap[city] = companies
		}

		for _, sector := range sectors {
			companies, ok := sectorRankingMap[sector]
			if !ok {
				sectorRankingMap[sector] = []int{companyId}
			} else {
				companies = append(companies, companyId)
				sectorRankingMap[sector] = companies
			}
		}
	}
}

func main() {
	newSectorCompanyMapping = data.GetNewSectorMapping(newSectorFile)
	results := processor.ReadCompanyRawDataJsonFile(inputFile, newSectorCompanyMapping)
	sort.Sort(results)

	finalResults := selectors.SelectWinners(results)
	uploaders.WriteOutput(outputFile, finalResults)

	uploaders.InsertToMongo(finalResults)

}
