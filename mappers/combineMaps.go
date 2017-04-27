package mappers

import "github.com/zling/zi-goproject/formats"

func convertToMap(list formats.MappedCompanyRecords, companyMap map[int]*formats.MappedCompanyRecord) map[int]*formats.MappedCompanyRecord {
	for _, elem := range list {
		value, ok := companyMap[elem.CompanyId]
		if !ok {
			companyMap[elem.CompanyId] = &formats.MappedCompanyRecord{
				CompanyId: elem.CompanyId,
				City:      elem.City,
				Sector:    elem.Sector,
			}
		} else {
			var city string
			var sectors []string
			if value.City == "" {
				city = elem.City
			} else {
				city = value.City
			}
			if len(value.Sector) == 0 {
				sectors = elem.Sector
			} else {
				sectors = value.Sector
			}
			companyMap[elem.CompanyId] = &formats.MappedCompanyRecord{
				CompanyId: elem.CompanyId,
				City:      city,
				Sector:    sectors,
			}
		}
	}

	return companyMap
}

func CombineCityAndSectorMapped(cityMapped formats.MappedCompanyRecords, sectorMapped formats.MappedCompanyRecords) map[int]*formats.MappedCompanyRecord {
	companyMap := make(map[int]*formats.MappedCompanyRecord)

	companyMap = convertToMap(cityMapped, companyMap)
	companyMap = convertToMap(sectorMapped, companyMap)

	return companyMap
}
