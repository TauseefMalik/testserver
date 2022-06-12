package main

import (
	"encoding/json"
	_ "errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/hashicorp/go-retryablehttp"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"sync"
)

//https://raw.githubusercontent.com/assignment132/assignment/main/duckduckgo.json
//https://raw.githubusercontent.com/assignment132/assignment/main/google.json
//https://raw.githubusercontent.com/assignment132/assignment/main/wikipedia.json

//Write a Golang based HTTP server which accepts GET requests with input parameter as “sortKey”
//and “limit”. The server queries three URLs mentioned below, combines the results from all three
//URLs, sorts them by the sortKey and returns the response. The Server should also limit the number
//of items in the API response to input parameter “limit”.


//sortKey String relevanceScore or views
//limit Integer Greater than 1, less than 200

//
//RESPONSE FORMAT
//{
//"data": [
//{
//"url": "www.yahoo.com/abc6",
//"views": 6000,
//"relevanceScore": 0.6
//},
//...
//“count”: <number of items> in the result,
//}

//1. The server should query URLs concurrently. (done)
//2. Server should have re-try mechanism and error handling on failures while querying URLs.
//3. Code should have a decent amount of Test Coverage.
//4. Provide README for testing and deployment. (done)

const (
	YAHOO = "https://raw.githubusercontent.com/assignment132/assignment/main/duckduckgo.json"
	WIKI = "https://raw.githubusercontent.com/assignment132/assignment/main/wikipedia.json"
	GOOGLE = "https://raw.githubusercontent.com/assignment132/assignment/main/google.json"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.DEBUG)
	// Routes
	e.GET("/info", GetInfo)

	// Start server
	e.Logger.Fatal(e.Start(":5000"))
}

var wg sync.WaitGroup
func GetInfo(c echo.Context) (err error) {
	c.Logger().Info("GetInfo invoked")
	var allSiteInfo []Info
	key := "views"
	limit := 1
	key = c.QueryParam("sortKey")
	limit, err = strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid query params")
	}
	if limit < 0 || limit >= 200{
		// default to 1 if out of range
		limit = 1
	}
	yahooCh := make(chan Info)
	wikiCh := make(chan Info)
	googleCh := make(chan Info)
	wg.Add(3)
	go DoGet(YAHOO, yahooCh)
	go DoGet(WIKI, wikiCh)
	go DoGet(GOOGLE, googleCh)
	for item := range yahooCh {
		allSiteInfo = append(allSiteInfo, item)
	}
	for item := range wikiCh {
		allSiteInfo = append(allSiteInfo, item)
	}
	for item := range googleCh {
		allSiteInfo = append(allSiteInfo, item)
	}
	wg.Wait()

	SortWebSites(c, key, allSiteInfo)
	//for _, d := range allSiteInfo {
	//	c.Logger().Debug(fmt.Sprintf("siteinfo %+v", d))
	//}
	c.Logger().Debug(fmt.Sprintf("len %d \n", len(allSiteInfo)))

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Server Error")
	}
	// handle limit greater than size of slice
	if limit > len(allSiteInfo) {
		limit = len(allSiteInfo)
	}
	return c.JSON(http.StatusOK, Response{Data: allSiteInfo[:limit], Count: limit})
}

func SortWebSites(c echo.Context, key string, allSiteInfo []Info) {
	if key != "" && key == "views" {
		c.Logger().Info(fmt.Sprintf("sort by views"))
		sort.Sort(InfoByViews(allSiteInfo))
	}
	if key != "" && key == "relevanceScore" {
		c.Logger().Info(fmt.Sprintf("sort by score"))
		sort.Sort(InfoByScore(allSiteInfo))
	}
}

func DoGet(url string, ch chan Info) {
	defer wg.Done()
	c := retryablehttp.NewClient()
	c.RetryMax = 2 //max retries
	resp, err := c.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var r map[string]interface{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		log.Fatal(err)
	}
	d := r["data"]
	ch = ParseData(d, ch)
	close(ch)
}

func ParseData(d interface{}, ch chan Info) chan Info {
	var s Info
	for _, m := range d.([]interface{}) {
		for k, v := range m.(map[string]interface{}) {
			switch k {
			case "url":
				s.Url = v.(string)
			case "views":
				s.Views = v.(float64)
			case "relevanceScore":
				s.RelevanceScore = v.(float64)
			}
		}
		ch <- s
	}
	return ch
}

