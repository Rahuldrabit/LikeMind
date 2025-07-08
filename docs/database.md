# Database Schema

LikeMind uses PostgreSQL for relational data, Redis for caching, and Qdrant as a vector database.

## PostgreSQL
Tables are created automatically from the Go models using GORM migrations. Important tables include users, chat sessions, messages and knowledge documents.

The initial schema can be found in `databases/postgresql/init.sql` and is executed automatically when running the Docker compose environment.

## Redis
Redis is used for caching chat history and other ephemeral data. The default configuration is provided in `databases/redis/redis.conf`.

## Qdrant
Qdrant stores vector embeddings for semantic search. The AI layer interacts with Qdrant directly.
