package selectors

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zling/zi-goproject/data"
	"github.com/zling/zi-goproject/formats"
)

func TagWinners(ceoList formats.MappedCompanyRecords) formats.MappedCompanyRecords {
	cityList, industryList := selectCeoAwards(ceoList)
	winners := getWinnerCompanyIds(cityList, industryList)

	for _, record := range ceoList {
		companyId := record.CompanyId
		_, ok := winners[companyId]
		if !ok {
			continue
		} else {
			record.IsWinner = true
		}
	}

	return ceoList
}

func printOutList(list map[string][]string, fileName string) {
	fileHandle, _ := os.Create(fileName)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	for key, value := range list {
		str := fmt.Sprintf("%s : %s", key, strings.Join(value, " | "))
		fmt.Fprintln(writer, str)
		writer.Flush()
	}
}

func getWinnerCompanyIds(cityList map[string][]string, industryList map[string][]string) map[string]bool {
	results := make(map[string]bool)

	for _, value := range cityList {
		for _, elem := range value {
			results[elem] = true
		}
	}

	for _, value := range industryList {
		for _, elem := range value {
			results[elem] = true
		}
	}

	return results
}

func selectCeoAwards(sortedList formats.MappedCompanyRecords) (map[string][]string, map[string][]string) {
	cityList := make(map[string][]string)
	industryList := make(map[string][]string)
	cityAwardCount := data.GetCityAwardCount()
	sectorAwardCount := data.GetIndustryAwardCount()

	for _, record := range sortedList {
		companyId := record.CompanyId
		city := record.City
		sectors := record.Sector

		//fmt.Println("%d\t%d\t%s\t%s\t%v", record.Rating, record.NumberOfVotes, companyId, city, sectors)

		count1, _ := cityAwardCount[city]
		value, ok := cityList[city]
		if !ok && record.Rating >= 80 {
			cityList[city] = []string{companyId}
		} else {
			if len(value) < count1 && record.Rating >= 80 {
				value = append(value, companyId)
			}
			cityList[city] = value
		}

		for _, sector := range sectors {
			count2, _ := sectorAwardCount[sector]
			value, ok := industryList[sector]
			if !ok && record.Rating >= 80 {
				industryList[sector] = []string{companyId}
			} else {
				if len(value) < count2 && record.Rating >= 80 {
					value = append(value, companyId)
				}
				industryList[sector] = value
			}
		}
	}

	printOutList(cityList, "files/cityAwardList.txt")
	printOutList(industryList, "files/sectorAwardList.txt")

	return cityList, industryList
}
