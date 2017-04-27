package mappers

import (
	"strconv"

	log "github.com/golang/glog"
)

const (
	ENTERPRISES = "Enterprises"
	SMBS        = "SMBs"
	EARLY_STAGE = "Early Stage Companies"
	NEW_KIDS    = "New Kids on the Block"
	BILLION     = 1000000000
	MILLION     = 1000000
)

func MapSegments(yearFounded string, revenue int) []string {
	year, err := strconv.Atoi(yearFounded)
	if err != nil {
		log.Error("year conversion failed.")
	}

	var segments []string

	if revenue > BILLION {
		segments = append(segments, ENTERPRISES)
	}
	if year >= 2010 && revenue < BILLION {
		segments = append(segments, SMBS)
	}
	if (year >= 2010 && year <= 2015) && revenue < 50*MILLION {
		segments = append(segments, EARLY_STAGE)
	}
	if year >= 2015 {
		segments = append(segments, NEW_KIDS)
	}

	return segments
}
