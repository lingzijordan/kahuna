package formats

type MappedCompanyRecord struct {
	CompanyId     string
	City          string
	Sector        []string
	Rating        int
	NumberOfVotes int
	IsWinner      bool
}

type MappedCompanyRecords []*MappedCompanyRecord

func (slice MappedCompanyRecords) Len() int {
	return len(slice)
}

func (slice MappedCompanyRecords) Less(i, j int) bool {
	if slice[i].Rating > slice[j].Rating {
		return true
	}

	if slice[i].Rating < slice[j].Rating {
		return false
	}

	if slice[i].NumberOfVotes > slice[j].NumberOfVotes {
		return true
	}

	if slice[i].NumberOfVotes < slice[j].NumberOfVotes {
		return false
	}

	return true
}

func (slice MappedCompanyRecords) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
