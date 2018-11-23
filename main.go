package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gocolly/colly"
)

type Question struct {
	ID          string `json:"id"`
	Summary     string `json:"summary"`
	Explanation string `json:"explanation"`
}

var Pairs map[string]*Question

func main() {
	Pairs = make(map[string]*Question)

	c := colly.NewCollector()

	c.OnHTML("tr[id]", func(e *colly.HTMLElement) {
		split := strings.Split(e.Attr("id"), "_")

		qType := split[1]
		ID := split[2]

		if qType == "summary" {

			_, ok := Pairs[ID]

			if !ok {
				Pairs[ID] = &Question{
					ID: ID,
				}
			}

			Pairs[ID].Summary = e.ChildText("td.border_gray:nth-child(2) > div")
		} else {
			_, ok := Pairs[ID]

			if !ok {
				Pairs[ID] = &Question{
					ID: ID,
				}
			}

			Pairs[ID].Explanation = e.ChildText("div[id=question_explanation]")
		}
	})

	c.OnScraped(func(r *colly.Response) {
		b, err := json.Marshal(Pairs)

		if err != nil {
			fmt.Println(err)
		}

		ioutil.WriteFile("data/out.json", b, 777)
	})

	c.Visit("http://localhost/index.html")
}
