package mappers

import (
	"errors"

	log "github.com/golang/glog"

	"github.com/zling/zi-goproject/data"
	"github.com/zling/zi-goproject/formats"
)

var newSectorCompanyMapping map[int][]string

func init() {
	newSectorCompanyMapping = data.GetNewSectorMapping("../../files/newSector.txt")
}

func industryMapper(industry string, maps map[string]string) (string, error) {
	if industry != "" {
		value, ok := maps[industry]
		if !ok {
			err := errors.New("industry that can't be mapped.")
			return "", err
		} else {
			return value, nil
		}
	}

	return "", errors.New("industry field is empty.")
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for _, v := range elements {
		if encountered[v] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[v] = true
			// Append to result slice.
			result = append(result, v)
		}
	}
	// Return the new slice.
	return result
}

func MapSector(sectors []string, companyId int) []string {
	sectorMap := data.GetIndustryMapping()
	var mappedSectors []string

	for _, sector := range sectors {
		if sector == "" {
			continue
		}
		mappedSector, err := industryMapper(sector, sectorMap)
		if err != nil {
			log.Warning("industry can't be mapped ", sector)
		} else {
			mappedSectors = append(mappedSectors, mappedSector)
		}
	}

	values, ok := newSectorCompanyMapping[companyId]
	if !ok {
		log.Warning("company can't be mapped to a new sector ", companyId)
		return removeDuplicates(mappedSectors)
	}
	mappedSectors = append(mappedSectors, values...)

	return removeDuplicates(mappedSectors)
}

func addToIndustries(industry string, mappedIndustries []string) []string {
	for _, elem := range mappedIndustries {
		if industry == elem {
			return mappedIndustries
		}
	}

	mappedIndustries = append(mappedIndustries, industry)
	return mappedIndustries
}

func MapSectors(records []*formats.Record) (formats.MappedCompanyRecords, []string) {
	var results formats.MappedCompanyRecords
	var unMappedSectors []string
	sectorMap := data.GetIndustryMapping()
	//newSectorCompanyMapping := data.GetNewSectorMapping("files/newSector.txt")

	for _, record := range records {
		var mappedIndustries []string

		companyId := record.CompanyId
		industry := record.IndustryName
		sector1 := record.Sector1
		sector2 := record.Sector2
		sector3 := record.Sector3

		newIndustry, err := industryMapper(industry, sectorMap)
		if err == nil {
			mappedIndustries = addToIndustries(newIndustry, mappedIndustries)
		}
		newIndustry, err = industryMapper(sector1, sectorMap)
		if err == nil {
			mappedIndustries = addToIndustries(newIndustry, mappedIndustries)
		}
		newIndustry, err = industryMapper(sector2, sectorMap)
		if err == nil {
			mappedIndustries = addToIndustries(newIndustry, mappedIndustries)
		}
		newIndustry, err = industryMapper(sector3, sectorMap)
		if err == nil {
			mappedIndustries = addToIndustries(newIndustry, mappedIndustries)
		}

		value, ok := newSectorCompanyMapping[companyId]
		if ok {
			mappedIndustries = append(mappedIndustries, value...)
		}

		if len(mappedIndustries) == 0 {
			unMappedSectors = append(unMappedSectors, []string{industry, sector1, sector2, sector3}...)
			continue
		}

		result := &formats.MappedCompanyRecord{
			CompanyId: companyId,
			Sector:    mappedIndustries,
		}

		results = append(results, result)

	}

	return results, unMappedSectors
}
