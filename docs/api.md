# LikeMind API Documentation

This document describes the main API endpoints provided by the backend service.

## Authentication
- `POST /api/v1/auth/register` – register a new user
- `POST /api/v1/auth/login` – log in and receive a JWT token
- `POST /api/v1/auth/refresh` – refresh an authentication token
- `POST /api/v1/auth/logout` – invalidate the current token

## Chat
- `GET /api/v1/chat/sessions` – list chat sessions for the current user
- `POST /api/v1/chat/sessions` – create a new chat session
- `GET /api/v1/chat/sessions/:id/messages` – fetch messages in a session
- `POST /api/v1/chat/sessions/:id/messages` – send a message to the AI

## Search
- `POST /api/v1/search/semantic` – perform a semantic search of the knowledge base
- `GET /api/v1/search/history` – view recent searches
- `POST /api/v1/search/index` – index documents for search

## Knowledge Base
- `GET /api/v1/knowledge/documents` – list uploaded documents
- `POST /api/v1/knowledge/documents` – upload a new document
- `DELETE /api/v1/knowledge/documents/:id` – remove a document

WebSocket endpoints for real time chat and notifications are available under `/api/v1/ws/*`.
