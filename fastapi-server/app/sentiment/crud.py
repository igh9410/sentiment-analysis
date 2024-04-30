# app/sentiment/crud.py
"""
The crud module contains the business logic for the sentiment analysis service. 
It defines a single function, analyze_sentiment, 
which takes a text string as input and returns the sentiment of the text as a string. 
The function sends a POST request to the sentiment analysis 
model server and returns the sentiment extracted from the response.
"""
import requests


def analyze_sentiment(text: str) -> str:
    """Analyze the sentiment of the given text using the sentiment analysis model.
    Args:
        text (str): The text to analyze.
    """
    payload = {"text": text}
    response = requests.post("http://fastapi-model:8000/predict", json=payload)

    if response.status_code == 200:
        sentiment_data = response.json()
        return sentiment_data["sentiment"]
    else:
        raise Exception("Failed to communicate with the model server")
