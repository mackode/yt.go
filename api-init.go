package main

import (
  "context"
  "flag"
  "log"
  "google.golang.org/api/option"
  "google.golang.org/api/youtube/v3"
)

func apiInit() (*youtube.Service, error) {
  apiKey := flag.String("api-key", "", "API Key")
  flag.Parse()

  if *apiKey == "" {
    log.Fatalf("Provide an API Key")
  }

  ctx := context.Background()
  service, err := youtube.NewService(ctx, option.WithAPIKey(*apiKey))
  return service, err
}
