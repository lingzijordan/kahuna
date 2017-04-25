package mappers

import "github.com/zling/zi-goproject/formats"

func convertToMap(list []*formats.NewCompanyRecord, companyMap map[string]*formats.NewCompanyRecord) map[string]*formats.NewCompanyRecord {
	for _, elem := range list {
		value, ok := companyMap[elem.CompanyId]
		if !ok {
			companyMap[elem.CompanyId] = &formats.NewCompanyRecord{
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
			companyMap[elem.CompanyId] = &formats.NewCompanyRecord{
				CompanyId: elem.CompanyId,
				City:      city,
				Sector:    sectors,
			}
		}
	}

	return companyMap
}

func CombineCityAndSectorMapped(cityMapped []*formats.NewCompanyRecord, sectorMapped []*formats.NewCompanyRecord) map[string]*formats.NewCompanyRecord {
	companyMap := make(map[string]*formats.NewCompanyRecord)

	companyMap = convertToMap(cityMapped, companyMap)
	companyMap = convertToMap(sectorMapped, companyMap)

	return companyMap
}
