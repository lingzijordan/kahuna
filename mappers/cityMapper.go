package mappers

import (
	"github.com/zling/zi-goproject/data"
	"github.com/zling/zi-goproject/formats"
)

func MapCities(records []*formats.Record) ([]*formats.NewCompanyRecord, []string) {
	var results []*formats.NewCompanyRecord
	var unMappedCities []string
	cityMap := data.GetCityMapping()
	for _, record := range records {
		companyId := record.CompanyId
		var city string
		value, ok := cityMap[record.CityName]
		if !ok {
			unMappedCities = append(unMappedCities, record.CityName)
			continue
		} else {
			city = value
		}

		result := &formats.NewCompanyRecord{
			CompanyId: companyId,
			City:      city,
		}
		results = append(results, result)
	}

	return results, unMappedCities
}
