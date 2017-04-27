package aggregators

import (
	"fmt"
	"os"
	"strconv"

	"github.com/zling/zi-goproject/formats"
)

func GetCeoRatingAndNumberOfVotes(records []*formats.Record, resultMap map[string]*formats.MappedCompanyRecord) formats.MappedCompanyRecords {
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

func JoinCompanyData(records []*formats.Record, ceoList formats.MappedCompanyRecords) []*formats.CompanyData {
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
				IsCityWinner:     company.IsCityWinner,
				IsIndustryWinner: company.IsIndustryWinner,
			}

			results = append(results, entry)
		}
	}

	return results
}
