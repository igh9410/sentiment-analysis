package main

import (
	"sentiment-api/internal/sentiment"
	"sentiment-api/router"
)

func main() {

	sentimentSvc := sentiment.NewService()
	sentimentHandler := sentiment.NewHandler(sentimentSvc)

	routerConfig := &router.RouterConfig{
		// ...
		SentimentHandler: sentimentHandler,
		// ...
	}

	router.InitRouter(routerConfig)

}
