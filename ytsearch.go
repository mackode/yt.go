package main

import (
	"fmt"
	"log"
)

func main() {
  service, err := apiInit()
  if err != nil {
    log.Fatalf("%v", err)
  }

  query := "Vlog"
  resp, err := service.Search.List([]string{"snippet"}).Q(query).Type("channel").Do()
  if err != nil {
    log.Fatalf("%v", err)
  }

  for i := 0; i < len(resp.Items); i++ {
    snippet := resp.Items[i].Snippet
    fmt.Printf("%20s: %s\n", snippet.ChannelTitle, snippet.ChannelId)
  }
}

