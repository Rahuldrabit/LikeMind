# AI Integration Guide

The AI layer is built with FastAPI and integrates with OpenAI for language models and Qdrant for vector search.

## Running Locally
1. Install dependencies:
   ```bash
   cd ai-layer
   pip install -r requirements.txt
   ```
2. Start the service:
   ```bash
   uvicorn main:app --reload
   ```

The service listens on `http://localhost:8000` and exposes a `/health` endpoint for checks.

## Environment Variables
- `OPENAI_API_KEY` – your OpenAI key
- `QDRANT_URL` – URL of the Qdrant instance
- `REDIS_URL` – Redis connection string
