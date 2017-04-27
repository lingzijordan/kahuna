package mappers

import (
	"errors"
	"fmt"

	"github.com/zling/zi-goproject/data"
	"github.com/zling/zi-goproject/formats"
)

func MapCities(records []*formats.Record) (formats.MappedCompanyRecords, []string) {
	var results formats.MappedCompanyRecords
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

		result := &formats.MappedCompanyRecord{
			CompanyId: companyId,
			City:      city,
		}
		results = append(results, result)
	}

	return results, unMappedCities
}

func MapCity(city string) (string, error) {
	cityMap := data.GetCityMapping()
	var mappedCity string

	value, ok := cityMap[city]
	if !ok {
		return "", errors.New(fmt.Sprintf("can't map to city ", city))
	} else {
		mappedCity = value
	}

	return mappedCity, nil
}
