package selectors

import (
	"github.com/zling/zi-goproject/data"
	"github.com/zling/zi-goproject/formats"
)

func TagWinners(ceoList formats.MappedCompanyRecords) formats.MappedCompanyRecords {
	cityList, industryList := selectCeoAwards(ceoList)
	cityWinners, industryWinners := getWinnerCompanyIds(cityList, industryList)

	for _, record := range ceoList {
		companyId := record.CompanyId
		_, ok := cityWinners[companyId]
		if !ok {
			continue
		} else {
			record.IsCityWinner = true
		}
	}

	for _, record := range ceoList {
		companyId := record.CompanyId
		_, ok := industryWinners[companyId]
		if !ok {
			continue
		} else {
			record.IsIndustryWinner = true
		}
	}

	return ceoList
}

func getWinnerCompanyIds(cityList map[string][]int, industryList map[string][]int) (map[int]bool, map[int]bool) {
	cityResults := make(map[int]bool)
	industryResults := make(map[int]bool)

	for _, value := range cityList {
		for _, elem := range value {
			cityResults[elem] = true
		}
	}

	for _, value := range industryList {
		for _, elem := range value {
			industryResults[elem] = true
		}
	}

	return cityResults, industryResults
}

func selectCeoAwards(sortedList formats.MappedCompanyRecords) (map[string][]int, map[string][]int) {
	cityList := make(map[string][]int)
	industryList := make(map[string][]int)
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
			cityList[city] = []int{companyId}
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
				industryList[sector] = []int{companyId}
			} else {
				if len(value) < count2 && record.Rating >= 80 {
					value = append(value, companyId)
				}
				industryList[sector] = value
			}
		}
	}

	return cityList, industryList
}

func SelectWinners(ceoList formats.CompanyDataJsonRecords) formats.CompanyDataJsonRecords {
	cityAwardCount := data.GetCityAwardCount()
	sectorAwardCount := data.GetIndustryAwardCount()
	cityCounter := make(map[string]int)
	sectorCounter := make(map[string]int)

	for _, ceo := range ceoList {
		city := ceo.MappedCity
		sectors := ceo.MappedSectors

		numberOfAwards, ok := cityAwardCount[city]
		if !ok {
			ceo.IsCityWinner = false
			continue
		}
		cityCount, ok := cityCounter[city]
		if !ok && ceo.CeoRating >= 80 {
			ceo.IsCityWinner = true
			cityCounter[city] = 1
			continue
		} else if cityCount < numberOfAwards && ceo.CeoRating >= 80 {
			ceo.IsCityWinner = true
			cityCounter[city] = cityCount + 1
		}

		for _, sector := range sectors {
			numberOfAwards, ok := sectorAwardCount[sector]
			if !ok {
				ceo.IsIndustryWinner = false
				continue
			}
			sectorCount, ok := sectorCounter[sector]
			if !ok && ceo.CeoRating >= 80 {
				ceo.IsIndustryWinner = true
				sectorCounter[sector] = 1
				continue
			} else if sectorCount < numberOfAwards && ceo.CeoRating >= 80 {
				ceo.IsIndustryWinner = true
				sectorCounter[sector] = sectorCount + 1
			}
		}
	}

	return ceoList
}
