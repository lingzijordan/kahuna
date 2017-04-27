package formats

type CompanyDataJson struct {
	CompanyId                          int      `json:"company_id"`
	CompanyNameShort                   string   `json:"company_name_short,omitempty"`
	CompanyNameLong                    string   `json:"company_name_long,omitempty"`
	CompanyUrl                         string   `json:"CP_URL,omitempty"`
	CompanyLogoSmall                   string   `json:"Company_Logo_Small,omitempty"`
	CompanyRevenue                     int      `json:"Company_Revenue,omitempty"`
	CeoName                            string   `json:"CEO_name"`
	CeoTitle                           string   `json:"Title,omitempty"`
	CeoPhoto                           string   `json:"CEO_photo,omitempty"`
	CeoRating                          int      `json:"CEO_Rating"`
	TotalNumberOfCeoRatings            int      `json:"Total_#_of_CEO_Ratings"`
	CityName                           string   `json:"City_Name,omitempty"`
	State                              string   `json:"state,omitempty"`
	Country                            string   `json:"country,omitempty"`
	Industry                           string   `json:"industry,omitempty"`
	IndustryName                       string   `json:"Industry_Name,omitempty"`
	Sector1                            string   `json:"sector_1,omitempty"`
	Sector2                            string   `json:"sector_2,omitempty"`
	Sector3                            string   `json:"sector_3,omitempty"`
	Sector4                            string   `json:"sector_4,omitempty"`
	Sector5                            string   `json:"sector_5,omitempty"`
	NumberOfActiveUsersOnOwler         int      `json:"#_of_Active_Users_on_Owler,omitempty"`
	YearFounded                        string   `json:"Year_Founded,omitempty"`
	NumberOfEmployees                  int      `json:"Number_of_Employees,omitempty"`
	NumberOfCeoRatingsFromEmployees    int      `json:"#_of_CEO_Ratings_from_Employees,omitempty"`
	NumberOfCeoRatingsFromNonEmployees int      `json:"#_of_CEO_Ratings_from_Non-Employees,omitempty"`
	TwitterProfile                     string   `json:"Twitter_Profile,omitempty"`
	Ownership                          string   `json:"ownership,omitempty"`
	NumberOfTwitterFollowers           int      `json:"#_of_Twitter_Followers,omitempty"`
	FollowersOnOwler                   int      `json:"Followers_on_Owler,omitempty"`
	CompanyStatus                      string   `json:"Company_Status,omitempty"`
	MappedCity                         string   `json:"mapped_city,omitempty"`
	MappedSectors                      []string `json:"mapped_sectors,omitempty"`
	IsCityWinner                       bool     `json:"is_city_winner",omitempty`
	IsIndustryWinner                   bool     `json:"is_industry_winner,omitempty"`
	MappedSegments                     []string `json:"mapped_segments"`
}

type CompanyDataJsonRecords []*CompanyDataJson

func (slice CompanyDataJsonRecords) Len() int {
	return len(slice)
}

func (slice CompanyDataJsonRecords) Less(i, j int) bool {
	if slice[i].CeoRating > slice[j].CeoRating {
		return true
	}

	if slice[i].CeoRating < slice[j].CeoRating {
		return false
	}

	if slice[i].TotalNumberOfCeoRatings > slice[j].TotalNumberOfCeoRatings {
		return true
	}

	if slice[i].TotalNumberOfCeoRatings < slice[j].TotalNumberOfCeoRatings {
		return false
	}

	return true
}

func (slice CompanyDataJsonRecords) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
