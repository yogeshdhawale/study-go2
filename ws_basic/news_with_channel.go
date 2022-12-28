package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type NewsMap3 struct {
	Keyword  string
	Location string
}

type NewsAggPage3 struct {
	Title string
	News  map[string]NewsMap3
}

type Sitemapindex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News3 struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

func indexHandlerC(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

var wg_news3 sync.WaitGroup

func newsRoutine(c chan News3, Location string) {

	defer wg_news3.Done()
	fmt.Println("Starting go routine for: ", Location)

	if false {
		var n News3
		n.Titles = []string{"1"}
		n.Keywords = []string{"1"}
		n.Locations = []string{Location}
		c <- n
	} else {
		var n News3
		resp, err := http.Get(Location)
		if err == nil {
			bytes, _ := ioutil.ReadAll(resp.Body)
			xml.Unmarshal(bytes, &n)
			resp.Body.Close()
			c <- n
		} else {
			fmt.Println("Error is:", err)
		}
	}
}

var wspostXML3 = []byte(`
<sitemapindex>
	<sitemap>
		<loc>http://www.washingtonpost.com/news-politics-sitemap.xml</loc>
	</sitemap>
	<sitemap>
		<loc>http://www.washingtonpost.com/news-blogs-politics-sitemap.xml</loc>
	</sitemap>
	<sitemap>
		<loc>http://www.washingtonpost.com/news-opinions-sitemap.xml</loc>
	</sitemap>
</sitemapindex>`)

func newsAggHandlerC(w http.ResponseWriter, r *http.Request) {

	var s SitemapIndex
	if false {
		resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &s)
		resp.Body.Close()
	} else {
		bytes := wspostXML3
		xml.Unmarshal(bytes, &s)
	}
	news_map := make(map[string]NewsMap3)

	cqueue := make(chan News3, 10)

	fmt.Println("Start time: ", time.Now())
	for _, Location := range s.Locations {
		wg_news3.Add(1)
		go newsRoutine(cqueue, Location)
	}

	wg_news3.Wait()
	close(cqueue)

	fmt.Println("End time: ", time.Now())

	max := 15
	fmt.Println("Post processing data: ", len(cqueue))
	for x := range cqueue {
		for idx, _ := range x.Keywords {
			news_map[x.Titles[idx]] = NewsMap3{x.Keywords[idx], x.Locations[idx]}
		}
		max--
		if max <= 0 {
			break
		}
	}

	/*for x, y := range news_map {
		fmt.Println(x)
		fmt.Println(y.Keyword)
		fmt.Println(y.Location)
		fmt.Println("...")
	}*/

	p := NewsAggPage3{Title: "Amazing News Aggregator", News: news_map}

	//t, _ := template.ParseFiles("template_basic_agg.html")
	t, _ := template.ParseFiles("template_basic_agg5.html")
	t.Execute(w, p)
}

func news_with_channel() {
	fmt.Println("Reading news with go routines")
	http.HandleFunc("/", indexHandlerC)
	http.HandleFunc("/agg/", newsAggHandlerC)
	http.ListenAndServe(":8000", nil)
}
