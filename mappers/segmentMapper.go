package mappers

import (
	"strconv"

	log "github.com/golang/glog"
)

const (
	OVERALL     = "Overall"
	ENTERPRISES = "Enterprises"
	SMBS        = "SMBs"
	EARLY_STAGE = "Early Stage Companies"
	NEW_KIDS    = "New Kids on the Block"
	USOVERALL   = "US Overall"
	BILLION     = 1000000000
	MILLION     = 1000000
)

func MapSegments(yearFounded string, revenue int, city string) []string {
	year, err := strconv.Atoi(yearFounded)
	nonUsCities := map[string]bool{
		"London":    true,
		"Toronto":   true,
		"New Delhi": true,
		"Mumbai":    true,
		"Vancouver": true,
		"Montreal":  true,
		"Melbourne": true,
		"Sydney":    true,
	}
	if err != nil {
		log.Error("year conversion failed.")
	}

	var segments []string

	if revenue > BILLION {
		segments = append(segments, ENTERPRISES)
	}
	if year <= 2008 && revenue < BILLION {
		segments = append(segments, SMBS)
	}
	if (year >= 2009 && year <= 2014) && revenue < 50*MILLION {
		segments = append(segments, EARLY_STAGE)
	}
	if year >= 2015 {
		segments = append(segments, NEW_KIDS)
	}
	_, ok := nonUsCities[city]
	if !ok {
		segments = append(segments, USOVERALL)
	}

	segments = append(segments, OVERALL)

	return segments
}
