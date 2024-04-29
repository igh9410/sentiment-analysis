package sentiment

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SentimentHandler struct {
	SentimentService
}

func NewHandler(s SentimentService) *SentimentHandler {
	return &SentimentHandler{
		SentimentService: s,
	}
}

func (h *SentimentHandler) HandleSentimentPost(c *gin.Context) {
	var req SentimentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Prepare the request payload
	payload, err := json.Marshal(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal request payload"})
		return
	}

	// Send the request to the FastAPI model server
	resp, err := http.Post("http://fastapi-model:8000/predict", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to communicate with the model server"})
		return
	}
	defer resp.Body.Close()

	// Read the response from the model server
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response from the model server"})
		return
	}

	var sentimentResponse SentimentResponse
	if err := json.Unmarshal(body, &sentimentResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal response from the model server"})
		return
	}

	c.JSON(http.StatusOK, sentimentResponse)
}
