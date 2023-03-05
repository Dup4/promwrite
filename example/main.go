package main

import (
	"context"
	"log"
	"time"

	"github.com/castai/promwrite"
)

func main() {
	client := promwrite.NewClient("http://127.0.0.1:1234/receive")

	_, err := client.Write(context.Background(), &promwrite.WriteRequest{
		TimeSeries: []promwrite.TimeSeries{
			{
				Labels: []promwrite.Label{
					{
						Name:  "__name__",
						Value: "my_metric_name",
					},
					{
						Name:  "dd",
						Value: "ddd",
					},
				},
				Sample: promwrite.Sample{
					Time:  time.Now(),
					Value: 123,
				},
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
