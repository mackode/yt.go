package main

import (
  "encoding/csv"
  "fmt"
  "os"
  "sort"
  "strconv"
)

func main() {
  file, err := os.Open("ytfetch.csv")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  type entry struct {
    title string
    views int64
  }

  entries := []entry{}
  r := csv.NewReader(file)
  for {
    e, err := r.Read()
    if err != nil {
      break
    }
    v, _ := strconv.ParseInt(e[1], 10, 64)
    entries = append(entries, entry{title: e[2], views: v})
  }

  sort.Slice(entries, func(i, j int) bool {
    return entries[i].views > entries[j].views
  })

  entries = entries[0:5]
  for _, e := range entries {
    fmt.Printf("%5d %s\n", e.views, e.title)
  }
}
