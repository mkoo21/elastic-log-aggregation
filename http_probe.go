package main

import (
  "fmt"
  "net/http"
  "time"
)

func main () {
  const RETRIES = 60
  es_up, kibana_up := false, false
  for i := 0; i < RETRIES; i++ {
    time.Sleep(5000 * time.Millisecond)
    if !es_up {
      fmt.Println("Checking if elasticsearch is up...")
      resp, _ := http.Get("http://elasticsearch:9200")
      if resp != nil && resp.StatusCode == 200 {
        fmt.Println("Got success response from elasticsearch")
        es_up = true
      }
    }
    if !kibana_up {
      fmt.Println("Checking if kibana is up...")
      resp, _ := http.Get("http://kibana:5601")
      if resp != nil && resp.StatusCode == 200 {
        fmt.Println("Got success response from kibana")
        kibana_up = true
      }
    }
    if es_up && kibana_up {
      fmt.Println("Service check successful. Continuing with setup.")
      break
    }
      
  }
  // Timeout TODO stderr (use log package)
  fmt.Println("Service check timed out.")
}
