package uploaders

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/zling/zi-goproject/formats"
)

func WriteOutput(file string, data formats.CompanyDataJsonRecords) {
	// final output data
	fileHandle, _ := os.Create(file)
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	for _, value := range data {
		str := fmt.Sprintf("%d\t%s\t%s\t%s\t%s\t%d\t%s\t%s\t%s\t%d\t%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s\t%d\t%d\t%d\t%s\t%s\t%d\t%d\t%s\t%s\t%s\t%t\t%t\t%s", value.CompanyId, value.CompanyNameShort, value.CompanyNameLong, value.CompanyUrl, value.CompanyLogoSmall, value.CompanyRevenue, value.CeoName, value.CeoTitle, value.CeoPhoto, value.CeoRating, value.TotalNumberOfCeoRatings, value.CityName, value.State, value.Country, value.Industry, value.IndustryName, value.Sector1, value.Sector2, value.Sector3, value.Sector4, value.Sector5, value.NumberOfActiveUsersOnOwler, value.YearFounded, value.NumberOfEmployees, value.NumberOfCeoRatingsFromEmployees, value.NumberOfCeoRatingsFromNonEmployees, value.TwitterProfile, value.Ownership, value.NumberOfTwitterFollowers, value.FollowersOnOwler, value.CompanyStatus, value.MappedCity, strings.Join(value.MappedSectors, " | "), value.IsCityWinner, value.IsIndustryWinner, strings.Join(value.MappedSegments, " | "))
		fmt.Fprintln(writer, str)
		writer.Flush()
	}
}
