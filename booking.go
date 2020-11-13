package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	// On every a element which has href attribute call callback
	c.OnResponse(func(r *colly.Response) {

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		if err != nil {
			return
		}
		//var price = doc.Find("#hprt-table > tbody > tr.hprt-table-cheapest-block.hprt-table-cheapest-block-fix.js-hprt-table-cheapest-block > td.d_pd_hp_price_left_align.hprt-table-cell.hprt-table-cell-price > div > div > div:nth-child(1) > div:nth-child(2) > span")
		//var price = doc.Find("#hprt-table > tbody > tr.hprt-table-cheapest-block.hprt-table-cheapest-block-fix.js-hprt-table-cheapest-block > td.d_pd_hp_price_left_align.hprt-table-cell.hprt-table-cell-price > div > div > div.bui-panel.bui-u-hidden.prco-price-area-popover > div > div > div.bui-grid__column-4 > div")
		hidenBody := doc.Find("#hprt-table > tbody > tr.hprt-table-cheapest-block.hprt-table-cheapest-block-fix.js-hprt-table-cheapest-block > td.d_pd_hp_price_left_align.hprt-table-cell.hprt-table-cell-price > div > div > div:nth-child(1) > div:nth-child(2) > div.bui-price-display__value.prco-text-nowrap-helper.prco-inline-block-maker-helper.prco-font16-helper")
		id, _ := hidenBody.Attr("data-popover-content-id")

		pathid := "#" + id + "> div > div.per-night-tt-table-wrapper.prco-inline-block-maker-helper.prco-font16-helper > table > tbody > tr > td.per-night-tt-table-cell-value > span"
		price := doc.Find(pathid)
		var children = price.Text()
		fmt.Println(children)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	var url = "https://www.booking.com/hotel/no/spitsbergen.en-gb.html?no_rooms=1&checkin=2021-02-03&checkout=2021-02-07&group_adults=2&group_children=0&req_adults=2&req_children=0&dist=0&type=total&selected_currency=NOK"
	c.Visit(url)
	url = "https://www.booking.com/hotel/es/sol-pelicanos-ocas.es.html?no_rooms=1&checkin=2021-02-03&checkout=2021-02-07&group_adults=2&group_children=0&req_adults=2&req_children=0&dist=0&type=total&selected_currency=EUR"
	c.Visit(url)
}
