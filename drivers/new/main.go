package main

import (
	"sort"

	"github.com/zling/zi-goproject/processor"
	"github.com/zling/zi-goproject/selectors"
	"github.com/zling/zi-goproject/uploaders"
)

func main() {
	fileName := "/Users/owler-zi/Documents/ceo_rating/da-641-2017-04-27.txt"

	results := processor.ReadCompanyRawDataJsonFile(fileName)
	sort.Sort(results)

	finalResults := selectors.SelectWinners(results)
	uploaders.WriteOutput("../../files/outputData.txt", finalResults)

	uploaders.InsertToMongo(finalResults)

}
