package main

import (
  "encoding/csv"
  "os"
  "strconv"
  "github.com/wcharczuk/go-chrt/v2"
)

func main() {
  file, err := os.Open("ytfetch.csv")
  if err != nil {
    panic(err)
  }
  defer file.Close()

  reader := csv.NewReader(file)
  records, err := reader.ReadAll()
  if err != nil {
    panic(err)
  }

  values := []chart.Value{}
  for _, record := range records {
    v, err := strconv.ParseFloat(record[1], 64)
    if err != nil {
      panic(err)
    }
    values = append(values, chart.Value{Value: v})
  }

  graph := chart.BarChart{
    Title: "YouTube Views By Video",
    Background: chart.Style{
      Padding: chart.Box{
        Top: 40,
      },
    },
    Height: 512,
    BarWidth: 60,
    Bars: values,
  }

  f, _ := os.Create("views.png")
  defer f.Close()
  graph.Render(chart.PNG, f)
}
