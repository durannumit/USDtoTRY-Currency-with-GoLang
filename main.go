package main

import (
  "fmt"
  "log"
  "github.com/PuerkitoBio/goquery"
)

func USDtoTRY() {
  doc, err := goquery.NewDocument("http://www.exchangerates.org.uk/USD-TRY-exchange-rate-history.html") 
  if err != nil {
    log.Fatal(err)
  }

  // Find the review items
  doc.Find("div#hd-maintable tr.colone").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
    currency  := s.Find("td").Text()
    day := s.Find("a").Text()
    fmt.Printf("Last 89 Days %d: %s - %s\n", i, currency, day)

  })
}

func liveCurrency() {

  doc, err := goquery.NewDocument("https://www.google.com/finance?q=USDTRY") 
  if err != nil {
    log.Fatal(err)
  }

  // Find the review items
  doc.Find("div#currency_value.sfe-section .sfe-break-bottom-4").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
    currency  := s.Find("span.pr").Text()
    day := s.Find("").Text()
    fmt.Printf("Live Currency %d: %s - %s\n", i, currency, day)
  })
}
  



func main() {
  
  liveCurrency()
  USDtoTRY()
  
}