package src

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strconv"
	"strings"
)

type Pubmed struct {
	selector  string
	searchKw  string
	goalURL   string
	paperurls map[string]string
}

func PubmedInit() Pubmed {
	p := Pubmed{
		selector: "#search-results",
		searchKw: "",
		goalURL:  "https://pubmed.ncbi.nlm.nih.gov/?term=",
	}
	return p
}

func collyHandeler(p *Pubmed) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong", err)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("Visited", response.Request.URL)
	})

	c.OnHTML(p.selector, func(element *colly.HTMLElement) {
		p.paperurls = make(map[string]string, 100)
		element.ForEach("a", func(_ int, htmlElement *colly.HTMLElement) {
			thisA := htmlElement.Attr("href")
			if strings.HasPrefix(thisA, "/") && strings.HasSuffix(thisA, "/") {
				thisA = "https://pubmed.ncbi.nlm.nih.gov" + thisA
				p.paperurls[thisA] = strings.TrimSpace(htmlElement.Text)

			}
		})
	})

	c.Visit(p.goalURL)
}

func (p *Pubmed) Search(searchKW string, number int) {
	(*p).goalURL += searchKW
	(*p).goalURL += "&size=" + strconv.Itoa(number)
	collyHandeler(p)
}

func (p *Pubmed) GetSearchResult(urlChan chan string) map[string]string {
	//for k, v := range (*p).paperurls {
	//	fmt.Printf("%s\t%s\n", k, v)
	//}
	for k, _ := range p.paperurls {
		urlChan <- k
	}
	close(urlChan)
	return p.paperurls
}
