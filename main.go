package main

import (
	"flag"
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"strings"
	"sync"
)

type Champion struct {
	position string
	winrate  string
}

func main() {
	var data = getChampionsData()

	var c string

	flag.StringVar(&c, "champion", "nidalee", "champion to gather data")

	flag.Parse()

	if champion, ok := data[strings.ToLower(c)]; ok {
		fmt.Printf("champion selected: %s \n", c)
		fmt.Printf("winrate: %s \n", champion.winrate)
		fmt.Printf("position: %s \n", champion.position)
		return
	}

	fmt.Printf("champion: %s, not found. \n", c)
}

func getChampionsData() map[string]Champion {
	c := colly.NewCollector()

	data := make(map[string]Champion)

	var wg sync.WaitGroup

	wg.Add(1)

	c.OnHTML("table", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			td1 := row.ChildText("td:nth-child(1)")
			td2 := row.ChildText("td:nth-child(2)")
			td5 := row.ChildText("td:nth-child(5)")

			data[strings.ToLower(td2)] = Champion{position: td1, winrate: td5}
		})
	})

	c.OnScraped(func(r *colly.Response) {
		wg.Done()
	})

	err := c.Visit("https://www.op.gg/champions?tier=all")

	if err != nil {
		log.Fatalf("Error while visiting URL: %v", err)
	}

	wg.Wait()

	return data
}
