package formats

type CompanyData struct {
	CompanyId        string
	CompanyUrl       string
	CompanyNameLong  string
	CompanyNameShort string
	CompanyLogoSmall string
	CeoName          string
	CeoPhoto         string
	NumberOfVotes    int
	Rating           int
	City             string
	Sector           []string
	IsWinner         bool
}
