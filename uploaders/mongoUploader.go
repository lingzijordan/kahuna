package uploaders

import (
	"fmt"
	"sync"

	"github.com/zling/zi-goproject/formats"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func InsertToMongo(data formats.CompanyDataJsonRecords) {
	//password := "d3vdb2@$!"
	url := fmt.Sprintf("mongodb://devtest:devtest@devdb.owler.com:37117/owler")
	session, err := mgo.Dial(url)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("owler").C("ceo_rating1")

	err = session.Ping()
	if err != nil {
		fmt.Println("didn't connect")
	}

	sourceInChan := make(chan *formats.CompanyDataJson, 10000)
	for _, elem := range data {
		sourceInChan <- elem
	}
	close(sourceInChan)

	fmt.Println(len(sourceInChan))

	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(inChan chan *formats.CompanyDataJson, c *mgo.Collection) {
			defer wg.Done()
			for elem := range inChan {
				document := &formats.CompanyDataMongo{
					CompanyId:        elem.CompanyId,
					CompanyUrl:       elem.CompanyUrl,
					CompanyNameLong:  elem.CompanyNameLong,
					CompanyNameShort: elem.CompanyNameShort,
					CompanyLogoSmall: elem.CompanyLogoSmall,
					CeoName:          elem.CeoName,
					CeoPhoto:         elem.CeoPhoto,
					NumberOfVotes:    elem.TotalNumberOfCeoRatings,
					Rating:           elem.CeoRating,
					City:             elem.MappedCity,
					Sectors:          elem.MappedSectors,
					IsCityWinner:     elem.IsCityWinner,
					IsIndustryWinner: elem.IsIndustryWinner,
					MappedSegments:   elem.MappedSegments,
				}

				upsertdata := bson.M{"$set": document}
				_, err := c.UpsertId(document.CompanyId, upsertdata)
				//err := c.Insert(document)
				if err != nil {
					fmt.Println("what is error here")
				}
			}
		}(sourceInChan, c)
	}

	index := mgo.Index{
		Key: []string{"$text:ceoname", "$text:companynamelong"},
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	wg.Wait()
}
