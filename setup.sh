#!/bin/bash

# LikeMind Setup Script
# This script helps you get started with the LikeMind platform

set -e

echo "🚀 LikeMind Platform Setup"
echo "=========================="

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed. Please install Docker first."
    exit 1
fi

# Check if Docker Compose is installed
if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose is not installed. Please install Docker Compose first."
    exit 1
fi

# Create .env file if it doesn't exist
if [ ! -f .env ]; then
    echo "📝 Creating .env file from template..."
    cp .env.example .env
    echo "⚠️  Please edit .env file with your configuration (especially OPENAI_API_KEY)"
fi

# Start database services first
echo "🗄️  Starting database services..."
docker-compose -f databases/docker-compose.yml up -d

# Wait for databases to be ready
echo "⏳ Waiting for databases to start..."
sleep 10

# Check database health
echo "🔍 Checking database health..."
docker exec likemind-postgres pg_isready -h localhost -p 5432 || {
    echo "❌ PostgreSQL is not ready"
    exit 1
}

docker exec likemind-redis redis-cli ping || {
    echo "❌ Redis is not ready"
    exit 1
}

echo "✅ Databases are ready!"

# Install frontend dependencies
if [ -d "frontend" ]; then
    echo "📦 Installing frontend dependencies..."
    cd frontend
    npm install
    cd ..
    echo "✅ Frontend dependencies installed"
fi

# Install backend dependencies
if [ -d "backend" ]; then
    echo "📦 Installing backend dependencies..."
    cd backend
    go mod tidy
    cd ..
    echo "✅ Backend dependencies installed"
fi

# Install AI layer dependencies
if [ -d "ai-layer" ] && [ -f "ai-layer/requirements.txt" ]; then
    echo "📦 Installing AI layer dependencies..."
    cd ai-layer
    python3 -m pip install -r requirements.txt
    cd ..
    echo "✅ AI layer dependencies installed"
fi

echo ""
echo "🎉 Setup complete!"
echo ""
echo "Next steps:"
echo "1. Edit .env file with your configuration"
echo "2. Add your OpenAI API key to .env"
echo "3. Run the development servers:"
echo ""
echo "   Frontend:  cd frontend && npm run dev"
echo "   Backend:   cd backend && go run cmd/server/main.go"
echo "   AI Layer:  cd ai-layer && python -m uvicorn main:app --reload"
echo ""
echo "Or start everything with Docker:"
echo "   docker-compose -f deployment/docker-compose.yml up"
echo ""
echo "Access the application at: http://localhost:3000"
