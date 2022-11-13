package main

import (
	"flag"
	"fmt"
	"os"
	"pubmedCrawler/src"
	"sort"
	"strings"
	"time"
)

var cliKW = flag.String("kw", "", "Input Your Search Key Words in PubMed\ne.g: R2-R3 myb domain")
var cliSS = flag.String("ss", "small", "Search Size choices (small, mid, big, verybig)")
var cliBaidutrans = flag.Bool("trans", false, "Using Baidu Translate API to translate Abstract to Chinese")
var cliBaiduID = flag.String("tid", "", "APP ID of Baidu Translation API\ndetails: http://api.fanyi.baidu.com/")
var cliBaiduSK = flag.String("tsk", "", "Secret Key of Your APP ID\ndetails: http://api.fanyi.baidu.com/")
var outPath = flag.String("out", "./default.docx", "Output File")

func main() {
	flag.Parse()
	var needToTranslate = true
	if *cliKW == "" {
		fmt.Printf("%v -h for help\n\n", os.Args[0])
		fmt.Printf("Input Your Search Key Words in PubMed\n")
		os.Exit(1)
	}
	if !src.IsTruePath(outPath) {

		fmt.Printf("Invalid out file path\n")
		os.Exit(1)

	}
	if *cliBaidutrans {
		if *cliBaiduID == "" || *cliBaiduSK == "" {
			fmt.Printf("-tid and -tsk parameters have to be set\n")
			os.Exit(1)
		} else {
			if str := src.CheckValidation(*cliBaiduID, *cliBaiduSK); str != "PARAM_FROM_TO_OR_Q_EMPTY" {
				fmt.Println(str)
				fmt.Println("ERROR: UNAUTHORIZED USER in BaiduTranslate")
				os.Exit(1)
			}
		}
	} else {
		if *cliBaiduID != "" || *cliBaiduSK != "" {
			fmt.Printf("-trans has to be included\n")
			os.Exit(1)
		}
		needToTranslate = false
	}
	arr := strings.Split(*cliKW, " ")
	searchKW := strings.Join(arr, "+")
	var searchSize int

	switch *cliSS {
	case "small":
		searchSize = 10
	case "mid":
		searchSize = 50
	case "big":
		searchSize = 100
	case "verybig":
		searchSize = 200
	default:
		fmt.Printf("Seach size not correct! (small|mid|big|verybig) is need, but %v input\n", *cliSS)
		os.Exit(1)
	}

	var paper src.PaperInfo
	var papers src.PaperInfos
	urlChan := make(chan string, 100)
	exitChan := make(chan bool, 4)
	papersChan := make(chan src.PaperInfo, 100)
	start := time.Now()
	p := src.PubmedInit()

	p.Search(searchKW, searchSize)
	go p.GetSearchResult(urlChan)
	for i := 0; i < 4; i++ {
		go paper.PaperPageParse(urlChan, papersChan, exitChan)
	}
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(exitChan)
		close(papersChan)
	}()
	for {
		v, ok := <-papersChan
		if !ok {
			break
		}
		if needToTranslate {
			v.BaiduTranslate(*cliBaiduID, *cliBaiduSK)
		}
		papers = append(papers, v)
	}
	sort.Sort(&papers)
	end := time.Since(start)
	src.SaveAsWord(papers, *outPath)

	fmt.Println(end)
}
