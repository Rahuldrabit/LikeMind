# LikeMind - Agentic AI Application

A comprehensive AI-powered knowledge management and automation platform built with modern technologies.

## 🏗️ Architecture

```
LikeMind/
├── Frontend (React.js/Next.js)     # Modern web interface
├── Backend (Golang)                # High-performance API server
├── AI Layer (Python)               # Advanced AI capabilities
├── Databases                       # Data persistence layer
│   ├── PostgreSQL                  # Relational data
│   ├── Redis                       # Caching & pub/sub
│   └── Qdrant                      # Vector database
└── Deployment                      # Container orchestration
```

## 🚀 Features

### 🤖 AI-Powered Core
- **Generative AI**: OpenAI GPT integration for intelligent conversations
- **Semantic Search**: Vector-based knowledge retrieval
- **Agentic Framework**: Specialized AI agents for different tasks
- **Embedding Models**: Advanced text understanding

### 💻 Frontend Features
- **Modern Dashboard**: Real-time analytics and monitoring
- **Chat Interface**: Intuitive AI conversation experience
- **Semantic Search UI**: Natural language search capabilities
- **Responsive Design**: Mobile-first approach with Tailwind CSS

### ⚡ Backend Capabilities
- **REST & WebSocket APIs**: Real-time communication
- **JWT Authentication**: Secure user management
- **Task Scheduling**: Background job processing
- **High Performance**: Golang-powered efficiency

### 📊 Data Layer
- **PostgreSQL**: Reliable relational data storage
- **Redis**: Fast caching and real-time features
- **Vector Database**: Semantic search and embeddings
- **Real-time Sync**: Consistent data across services

## 🛠️ Tech Stack

### Frontend
- **Framework**: Next.js 14 with TypeScript
- **Styling**: Tailwind CSS with Headless UI
- **State Management**: Zustand
- **Data Fetching**: React Query
- **Icons**: Heroicons
- **Animation**: Framer Motion

### Backend
- **Language**: Go 1.21
- **Framework**: Gin (HTTP) + Gorilla WebSocket
- **Database ORM**: GORM
- **Authentication**: JWT with golang-jwt
- **Cache**: Redis with go-redis
- **Validation**: go-playground/validator

### AI Layer
- **Language**: Python 3.11
- **Framework**: FastAPI + Uvicorn
- **AI Libraries**: LangChain, OpenAI, Transformers
- **Vector DB**: Qdrant client
- **Data Processing**: NumPy, Pandas

### Databases
- **Primary**: PostgreSQL 15
- **Cache**: Redis 7
- **Vector**: Qdrant (latest)
- **Management**: pgAdmin, Redis Commander

### DevOps
- **Containerization**: Docker & Docker Compose
- **Orchestration**: Kubernetes (optional)
- **Monitoring**: Prometheus + Grafana
- **Reverse Proxy**: Nginx
- **CI/CD**: GitHub Actions (configurable)

## 📋 Prerequisites

- **Docker** & **Docker Compose**
- **Node.js** 18+ (for local frontend development)
- **Go** 1.21+ (for local backend development)
- **Python** 3.11+ (for local AI development)
- **OpenAI API Key** (for AI features)

## 🚀 Quick Start

### 1. Clone & Environment Setup
```bash
git clone <repository-url>
cd LikeMind
cp .env.example .env
# Edit .env with your configuration
```

### 2. Start with Docker (Recommended)
```bash
# Start all services
docker-compose -f deployment/docker-compose.yml up -d

# View logs
docker-compose -f deployment/docker-compose.yml logs -f
```

### 3. Database Setup
```bash
# Start only databases first
docker-compose -f databases/docker-compose.yml up -d

# Initialize database (automatic with docker-compose)
# Manual: psql -h localhost -U postgres -d likemind -f databases/postgresql/init.sql
```

### 4. Environment Variables
```bash
# Copy example environment file
cp .env.example .env

# Required variables:
OPENAI_API_KEY=your_openai_api_key
JWT_SECRET=your_super_secret_jwt_key
DATABASE_URL=postgres://postgres:postgres@localhost:5432/likemind
REDIS_URL=redis://localhost:6379
VECTOR_DB_URL=http://localhost:6333
```

## 🏃‍♂️ Development

### Frontend Development
```bash
cd frontend
npm install
npm run dev
# Runs on http://localhost:3000
```

### Backend Development
```bash
cd backend
go mod tidy
go run cmd/server/main.go
# Runs on http://localhost:8080
```

### AI Layer Development
```bash
cd ai-layer
pip install -r requirements.txt
python -m uvicorn main:app --reload
# Runs on http://localhost:8000
```

## 📖 API Documentation

### Authentication Endpoints
- `POST /api/v1/auth/register` - User registration
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/refresh` - Token refresh
- `POST /api/v1/auth/logout` - User logout

### Chat Endpoints
- `GET /api/v1/chat/sessions` - Get user chat sessions
- `POST /api/v1/chat/sessions` - Create new chat session
- `GET /api/v1/chat/sessions/:id/messages` - Get session messages
- `POST /api/v1/chat/sessions/:id/messages` - Send message

### Search Endpoints
- `POST /api/v1/search/semantic` - Semantic search
- `GET /api/v1/search/history` - Search history
- `POST /api/v1/search/index` - Index documents

### Knowledge Base Endpoints
- `GET /api/v1/knowledge/documents` - List documents
- `POST /api/v1/knowledge/documents` - Upload document
- `DELETE /api/v1/knowledge/documents/:id` - Delete document

### WebSocket Endpoints
- `WS /api/v1/ws/chat` - Real-time chat
- `WS /api/v1/ws/notifications` - System notifications

## 🧪 Testing

### Frontend Tests
```bash
cd frontend
npm run test
npm run test:e2e
```

### Backend Tests
```bash
cd backend
go test ./...
go test -race ./...
```

### AI Layer Tests
```bash
cd ai-layer
python -m pytest
python -m pytest --cov=.
```

## 📊 Monitoring & Observability

### Access Points
- **Application**: http://localhost:3000
- **API**: http://localhost:8080
- **AI Service**: http://localhost:8000
- **Database Admin**: http://localhost:5050 (pgAdmin)
- **Redis Admin**: http://localhost:8081 (Redis Commander)
- **Monitoring**: http://localhost:9090 (Prometheus)
- **Dashboards**: http://localhost:3001 (Grafana)

### Health Checks
```bash
# Backend health
curl http://localhost:8080/health

# AI service health
curl http://localhost:8000/health

# Database connection
docker exec likemind-postgres pg_isready
```

## 🔒 Security Features

- **JWT Authentication** with secure token handling
- **CORS Configuration** for cross-origin security
- **Input Validation** across all API endpoints
- **SQL Injection Protection** with parameterized queries
- **Rate Limiting** to prevent abuse
- **Secure Headers** with proper middleware

## 🌐 Deployment Options

### Development
```bash
docker-compose -f deployment/docker-compose.yml up
```

### Production
```bash
# With SSL and production optimizations
docker-compose -f deployment/docker-compose.prod.yml up -d
```

### Kubernetes (Optional)
```bash
kubectl apply -f deployment/k8s/
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Troubleshooting

### Common Issues

1. **Database Connection Failed**
   ```bash
   # Check if PostgreSQL is running
   docker ps | grep postgres
   
   # Check logs
   docker logs likemind-postgres
   ```

2. **Redis Connection Failed**
   ```bash
   # Test Redis connection
   docker exec likemind-redis redis-cli ping
   ```

3. **OpenAI API Issues**
   ```bash
   # Verify API key in environment
   echo $OPENAI_API_KEY
   
   # Check AI service logs
   docker logs likemind-ai
   ```

4. **Frontend Build Issues**
   ```bash
   # Clear cache and reinstall
   cd frontend
   rm -rf .next node_modules
   npm install
   npm run build
   ```

### Performance Optimization

- **Database**: Configure connection pooling and indexes
- **Redis**: Optimize memory usage and eviction policies
- **Backend**: Enable Go's runtime optimizations
- **Frontend**: Implement code splitting and lazy loading

## 📚 Documentation

- [API Documentation](docs/api.md)
- [Database Schema](docs/database.md)
- [AI Integration Guide](docs/ai-integration.md)
- [Deployment Guide](docs/deployment.md)
- [Contributing Guidelines](docs/contributing.md)

## 🎯 Roadmap

- [ ] Advanced agent capabilities
- [ ] Multi-modal AI support (images, audio)
- [ ] Real-time collaboration features
- [ ] Advanced analytics dashboard
- [ ] Mobile application
- [ ] Plugin system
- [ ] Enterprise SSO integration
- [ ] Advanced workflow automation

---

**Built with ❤️ by the LikeMind team**
