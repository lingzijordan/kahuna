package data

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/zling/zi-goproject/formats"
)

func GetCompanyRawRecords(fileName string) []*formats.Record {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var records []*formats.Record

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "\t")

		if len(arr) < 21 {
			//fmt.Println(line)
			continue
		}

		record := &formats.Record{
			CompanyId:                  arr[0],
			CompanyUrl:                 arr[1],
			CompanyNameLong:            arr[2],
			CompanyNameShort:           arr[3],
			CompanyLogoSmall:           arr[4],
			CeoName:                    arr[5],
			CeoPhoto:                   arr[6],
			TotalNumberOfCeoRatings:    arr[7],
			CeoRating:                  arr[8],
			FollowersOnOwler:           arr[9],
			NumberOfActiveUsersOnOwler: arr[10],
			TwitterProfile:             arr[11],
			NumberOfTwitterFollowers:   arr[12],
			CityName:                   arr[13],
			IndustryName:               arr[14],
			YearFounded:                arr[15],
			CompanyStatus:              arr[16],
			CountryName:                arr[17],
			Sector1:                    arr[18],
			Sector2:                    arr[19],
			Sector3:                    arr[20],
		}
		records = append(records, record)
	}

	return records
}
