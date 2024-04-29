package sentiment

type SentimentRequest struct {
	Text string `json:"text"`
}

type SentimentResponse struct {
	Sentiment string `json:"sentiment"`
}
