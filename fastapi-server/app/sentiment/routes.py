# app/sentiment/routes.py
from fastapi import APIRouter, HTTPException
from app.sentiment.models import SentimentRequest, SentimentResponse
from app.sentiment.crud import analyze_sentiment

router = APIRouter()


@router.post("/", response_model=SentimentResponse)
async def sentiment_analysis(request: SentimentRequest):
    try:
        sentiment = analyze_sentiment(request.text)
        return SentimentResponse(sentiment=sentiment)
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))
