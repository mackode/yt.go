package main

import (
  "google.golang.org/api/youtube/v3"
  "log"
)

type chStats struct {
  vid string
  title string
  views uint64
}

func channelViews(service *youtube.Service, id string) ([]chStats, error) {
  stats := []chStats{}
  resp, err := service.Channels.List([]string{"contentDetails"}).Id(id).Do()
  if err != nil {
    log.Fatalf("%v", err)
  }

  if len(resp.Items) == 0 {
    log.Fatalf("Channel not found")
  }

  plid := resp.Items[0].ContentDetails.RelatedPlaylists.Uploads

  pageToken := ""
  for {
    plResp, err := service.PlaylistItems.List([]string{"snippet"}).PlaylistId(plid).MaxResults(50).PageToken(pageToken).Do()
    if err != nil {
      log.Fatalf("%v", err)
    }

    for _, item := range plResp.Items {
      videoID := item.Snippet.ResourceId.VideoId
      videoResp, err := service.Videos.List([]string{"statistics"}).Id(videoID).Do()
      if err != nil {
        log.Fatalf("%v", err)
      }

      video := videoResp.Items[0]
      viewCount := video.Statistics.ViewCount
      stats = append(stats, chStats{vid: videoID, views: viewCount, title: item.Snippet.Title})
    }

    pageToken = plResp.NextPageToken

    if pageToken == "" {
      break
    }
  }

  return stats, nil
}
