package main

import (
  "log"
  "fmt"
  "encoding/csv"
  "os"
  "flag"
)

func main() {
  service, err := apiInit()
  if err != nil {
    log.Fatalf("%v", err)
  }

  channelId := flag.String("channel-id", "", "Channel Id")
  flag.Parse()
  if *channelId == "" {
    log.Fatalf("Provide a Channel Id")
  }

  stats, err := channelViews(service, *channelId)
  if err != nil {
    log.Fatalf("%v", err)
  }

  w := csv.NewWriter(os.Stdout)
  defer w.Flush()

  for _, stat := range stats {
    w.Write([]string{stat.vid, fmt.Sprintf("%d", stat.views), stat.title})
  }
}
