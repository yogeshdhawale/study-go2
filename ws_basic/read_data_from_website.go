package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

var wspostXML = []byte(`
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

/*
type Location struct {
	Loc string `xml:"loc"`
}

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}
*/

type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

func read_data_from_website() {
	fmt.Println("Getting data - sitemap ")
	//resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	//bytes, _ := ioutil.ReadAll(resp.Body)
	//str_body := string(bytes)
	//fmt.Println(str_body)
	//resp.Body.Close()

	bytes := wspostXML
	var s SitemapIndex
	xml.Unmarshal(bytes, &s)
	//fmt.Println("Values:", s.Locations)

	fmt.Println("Data is:")
	for _, x := range s.Locations {
		fmt.Println("=>", x)
	}
	fmt.Println()

}

var wspostNewsXML = []byte(`
<url>
	<news>
		<title>title 1</title>
		<keywords>keywords 1</keywords>
	</news>
	<news>
		<title>title 2</title>
		<keywords>keywords 2</keywords>
	</news>
	<news>
		<title>title 3</title>
		<keywords>keywords 3</keywords>
	</news>
	<locations>xyz</locations>
</url>`)

var wspoliticalNewsXML = []byte(`
<url>
	<loc>https://www.washingtonpost.com/politics/2022/10/24/florida-governor-debate-desantis-crist/</loc>
	<news>
	<news:title>
		<![CDATA[ DeSantis dodges questions on 2024, abortion at Florida gubernatorial debate ]]>
	</news:title>
	<news:publication_date>2022-10-25T02:50:39.328Z</news:publication_date>
	</news>
</url>
<url>
	<loc>https://www.washingtonpost.com/politics/2022/10/24/biden-ukraine-liberals/</loc>
	<news>
		<news:title>
		<![CDATA[ Liberals urge Biden to rethink Ukraine strategy ]]>
		</news:title>
	</news>
	<news:publication_date>2022-10-25T00:33:17.400Z</news:publication_date>
</url>
<url>
	<loc>https://www.washingtonpost.com/politics/2022/10/24/biden-democratic-national-committee-midterms/</loc>
	<news>
		<news:title>
		<![CDATA[ Post Politics Now: Biden warns that Republican majority could force government shutdown ]]>
		</news:title>
	</news>
	<news:publication_date>2022-10-25T00:29:09.430Z</news:publication_date>
</url>
`)

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywards  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>locations"`
}

type PoliticsNews struct {
	Titles    []string `xml:"url>news>title"`
	PubDate   []string `xml:"url>news>publication_date"`
	Locations []string `xml:"url>loc"`
}

type PnMap struct {
	Title    string
	Location string
	PubDate  string
}

func dump_political_news_from_website() {

	fmt.Println("Getting data - political news ")
	resp, _ := http.Get("https://www.washingtonpost.com/news-politics-sitemap.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var n PoliticsNews
	xml.Unmarshal(bytes, &n)

	/*fmt.Println("Data is:")
	for _, x := range n.Locations {
		fmt.Println("=>", x)
	}
	fmt.Println()
	*/
	pnmap := make(map[string]PnMap)
	max := 5
	for _, x := range n.Locations {
		fmt.Println("Processing location: ", x)
		//resp, _ := http.Get(x)
		//bytes, _ := ioutil.ReadAll(resp.Body)
		//xml.Unmarshal(bytes, &n)
		for idx := range n.Titles {
			pnmap[n.Titles[idx]] = PnMap{n.Titles[idx], n.Locations[idx], n.PubDate[idx]}
		}
		max--
		if max <= 0 {
			break
		}
	}

	fmt.Println("From map of size:", len(pnmap))
	for idx, data := range pnmap {
		fmt.Print("title:", idx)
		fmt.Print(", ", data.Location)
		fmt.Println(", ", data.PubDate)
	}
	fmt.Println()
}
