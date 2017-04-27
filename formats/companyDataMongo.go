package formats

type CompanyDataMongo struct {
	CompanyId        int `bson:"_id,omitempty"`
	CompanyUrl       string
	CompanyNameLong  string
	CompanyNameShort string
	CompanyLogoSmall string
	CeoName          string
	CeoPhoto         string
	NumberOfVotes    int
	Rating           int
	City             string
	Sectors          []string
	IsCityWinner     bool
	IsIndustryWinner bool
	MappedSegments   []string
}
