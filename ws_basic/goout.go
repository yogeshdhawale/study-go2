package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type PoliticsNews2 struct {
	Titles    []string `xml:"url>news>title"`
	PubDate   []string `xml:"url>news>publication_date"`
	Locations []string `xml:"url>loc"`
}

type PnMap2 struct {
	Title    string
	Location string
	PubDate  string
	Data     string
}

type NewsAgg struct {
	Title string
	News  map[string]PnMap2
}

func indexHanderAgg(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<h1>Text is here ...</h1>
	<p>This is para</p>
	`)
}

func getAggData(pnmap map[string]PnMap2) {
	//pnmap := make(map[string]PnMap2)

	if true {

		fmt.Println("Getting data - political news ")
		resp, _ := http.Get("https://www.washingtonpost.com/news-politics-sitemap.xml")
		bytes, _ := ioutil.ReadAll(resp.Body)
		var n PoliticsNews
		xml.Unmarshal(bytes, &n)
		resp.Body.Close()

		max := 15

		for _, x := range n.Locations {
			fmt.Println("Processing location: ", x)

			for idx := range n.Titles {
				pnmap[n.Titles[idx]] = PnMap2{n.Titles[idx], n.Locations[idx], n.PubDate[idx], ""}
				max--

				if max <= 0 {
					break
				}
			}
		}
	} else {
		pnmap["1"] = PnMap2{"1", "http://1", "111", ""}
		pnmap["2"] = PnMap2{"2", "http://2", "222", ""}
		pnmap["3"] = PnMap2{"3", "http://3", "333", ""}
	}
}
func newsAggHandler(w http.ResponseWriter, r *http.Request) {

	pnmap := make(map[string]PnMap2)
	getAggData(pnmap)
	fmt.Println("News loaded ... items are: ", len(pnmap))

	for i := range pnmap {
		fmt.Println("Example row:", pnmap[i])
		break
	}

	data := NewsAgg{Title: "News Articles", News: pnmap}
	fmt.Println("Example Data:", data.Title, data.News)

	//t, _ := template.ParseFiles("template_basic_agg2.html")
	//t, _ := template.ParseFiles("template_basic_agg3.html")
	t, _ := template.ParseFiles("template_basic_agg4.html")
	err := t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func newsAggPage() {
	http.HandleFunc("/", indexHanderAgg)
	http.HandleFunc("/agg/", newsAggHandler)
	http.ListenAndServe(":8000", nil)
}

func main() {
	//basic_http()

	// map eg
	//mapeg()

	//routine_eg()

	// channel eg
	//channel_eg()

	//basic_agg_eg()
	//dump_political_news_from_website()
	//read_data_from_website()

	//newsAggPage()
	news_with_channel()
}
