package src

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"log"
	"time"
)

type PaperInfo struct {
	Title     string
	Author    string
	DOI       string
	Abstract  string
	Time      string
	Journal   string
	Content   string
	Translate string
}
type PaperInfos []PaperInfo

func (this *PaperInfo) PaperPageParse(urlChan chan string, papersChan chan PaperInfo, exitChan chan bool) {
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	err := c.Limit(&colly.LimitRule{
		DomainRegexp: `.\.gov`,
		RandomDelay:  3 * time.Second,
		Parallelism:  12,
	})
	if err != nil {
		log.Println(err)
	}
	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visiting: %v\n", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong ", err)
	})

	c.OnResponse(func(r *colly.Response) {

	})

	c.OnHTML("body", func(element *colly.HTMLElement) {
		this.Author = ""
		element.ForEach("#article-details .authors-list .full-name", func(_ int, htmlElement *colly.HTMLElement) {
			this.Author += fmt.Sprintf("%s, ", htmlElement.Text)
		})
		this.Content = element.ChildText(".abstract-content>p")
		this.Time = element.ChildText(".cit")
		this.Journal = element.ChildText("#full-view-journal-trigger")
		this.DOI = element.ChildText(".full-view .identifiers .doi .id-link")
		this.Title = element.ChildText(".full-view .heading-title")
	})
	for {
		v, ok := <-urlChan
		if !ok {
			break
		}
		c.Visit(v)
		papersChan <- *this
	}

	exitChan <- true
}

func (this *PaperInfos) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

func (this *PaperInfos) Less(i, j int) bool {
	return (*this)[i].Time > (*this)[j].Time
}

func (this *PaperInfos) Len() int {
	return len(*this)
}
