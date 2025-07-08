"""
LikeMind AI Layer - Agentic Framework Implementation
This module provides the core AI functionality including:
- OpenAI GPT integration
- Embedding generation and semantic search
- Agentic framework with LangChain
- Custom agent implementations
"""

from typing import List, Dict, Any, Optional
import openai
import numpy as np
from langchain.agents import initialize_agent, Tool
from langchain.agents import AgentType
from langchain.memory import ConversationBufferMemory
from langchain.llms import OpenAI
from langchain.tools import BaseTool
from langchain.embeddings import OpenAIEmbeddings
from langchain.vectorstores import Qdrant
from langchain.schema import Document
import logging
import asyncio
from datetime import datetime
import json

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

class LikeMindAI:
    """Main AI service class that orchestrates all AI operations"""
    
    def __init__(self, openai_api_key: str, qdrant_url: str = "http://localhost:6333"):
        self.openai_api_key = openai_api_key
        self.qdrant_url = qdrant_url
        
        # Initialize OpenAI
        openai.api_key = openai_api_key
        
        # Initialize embeddings
        self.embeddings = OpenAIEmbeddings(openai_api_key=openai_api_key)
        
        # Initialize vector store
        self.vector_store = Qdrant(
            url=qdrant_url,
            collection_name="likemind_knowledge",
            embeddings=self.embeddings
        )
        
        # Initialize LLM
        self.llm = OpenAI(temperature=0.7, openai_api_key=openai_api_key)
        
        # Initialize memory
        self.memory = ConversationBufferMemory(
            memory_key="chat_history",
            return_messages=True
        )
        
        # Initialize custom tools
        self.tools = self._create_custom_tools()
        
        # Initialize agent
        self.agent = self._create_agent()
        
    def _create_custom_tools(self) -> List[BaseTool]:
        """Create custom tools for the AI agent"""
        
        class KnowledgeSearchTool(BaseTool):
            name = "knowledge_search"
            description = "Search through the knowledge base for relevant information"
            
            def _run(self, query: str) -> str:
                try:
                    # Perform semantic search
                    docs = self.vector_store.similarity_search(query, k=5)
                    
                    if not docs:
                        return "No relevant information found in the knowledge base."
                    
                    # Format results
                    results = []
                    for doc in docs:
                        results.append({
                            "content": doc.page_content,
                            "metadata": doc.metadata
                        })
                    
                    return json.dumps(results, indent=2)
                except Exception as e:
                    logger.error(f"Error in knowledge search: {e}")
                    return f"Error searching knowledge base: {str(e)}"
            
            async def _arun(self, query: str) -> str:
                return self._run(query)
        
        class DataAnalysisTool(BaseTool):
            name = "data_analysis"
            description = "Analyze data and provide insights"
            
            def _run(self, data: str) -> str:
                try:
                    # Simple data analysis implementation
                    # In production, this would integrate with pandas, numpy, etc.
                    return f"Analysis of data: {data[:100]}... (truncated)"
                except Exception as e:
                    logger.error(f"Error in data analysis: {e}")
                    return f"Error analyzing data: {str(e)}"
            
            async def _arun(self, data: str) -> str:
                return self._run(data)
        
        class TaskSchedulerTool(BaseTool):
            name = "task_scheduler"
            description = "Schedule and manage tasks"
            
            def _run(self, task: str) -> str:
                try:
                    # Task scheduling logic
                    return f"Task scheduled: {task}"
                except Exception as e:
                    logger.error(f"Error scheduling task: {e}")
                    return f"Error scheduling task: {str(e)}"
            
            async def _arun(self, task: str) -> str:
                return self._run(task)
        
        return [
            KnowledgeSearchTool(),
            DataAnalysisTool(),
            TaskSchedulerTool()
        ]
    
    def _create_agent(self):
        """Create the main AI agent"""
        return initialize_agent(
            tools=self.tools,
            llm=self.llm,
            agent=AgentType.CONVERSATIONAL_REACT_DESCRIPTION,
            memory=self.memory,
            verbose=True
        )
    
    async def generate_response(self, user_input: str, context: Optional[Dict] = None) -> Dict[str, Any]:
        """Generate AI response to user input"""
        try:
            # Add context if provided
            if context:
                user_input = f"Context: {json.dumps(context)}\n\nUser: {user_input}"
            
            # Get response from agent
            response = await asyncio.to_thread(self.agent.run, user_input)
            
            return {
                "response": response,
                "timestamp": datetime.now().isoformat(),
                "status": "success"
            }
        except Exception as e:
            logger.error(f"Error generating response: {e}")
            return {
                "response": "I encountered an error processing your request. Please try again.",
                "timestamp": datetime.now().isoformat(),
                "status": "error",
                "error": str(e)
            }
    
    async def add_to_knowledge_base(self, documents: List[Dict[str, Any]]) -> Dict[str, Any]:
        """Add documents to the knowledge base"""
        try:
            # Convert to LangChain documents
            docs = []
            for doc in documents:
                docs.append(Document(
                    page_content=doc["content"],
                    metadata=doc.get("metadata", {})
                ))
            
            # Add to vector store
            await asyncio.to_thread(self.vector_store.add_documents, docs)
            
            return {
                "status": "success",
                "added_documents": len(docs),
                "timestamp": datetime.now().isoformat()
            }
        except Exception as e:
            logger.error(f"Error adding to knowledge base: {e}")
            return {
                "status": "error",
                "error": str(e),
                "timestamp": datetime.now().isoformat()
            }
    
    async def semantic_search(self, query: str, k: int = 5) -> Dict[str, Any]:
        """Perform semantic search on knowledge base"""
        try:
            # Search for similar documents
            docs = await asyncio.to_thread(
                self.vector_store.similarity_search_with_score, 
                query, 
                k=k
            )
            
            # Format results
            results = []
            for doc, score in docs:
                results.append({
                    "content": doc.page_content,
                    "metadata": doc.metadata,
                    "similarity_score": float(score)
                })
            
            return {
                "query": query,
                "results": results,
                "timestamp": datetime.now().isoformat(),
                "status": "success"
            }
        except Exception as e:
            logger.error(f"Error in semantic search: {e}")
            return {
                "query": query,
                "results": [],
                "timestamp": datetime.now().isoformat(),
                "status": "error",
                "error": str(e)
            }
    
    async def generate_embeddings(self, texts: List[str]) -> Dict[str, Any]:
        """Generate embeddings for given texts"""
        try:
            # Generate embeddings
            embeddings = await asyncio.to_thread(
                self.embeddings.embed_documents, 
                texts
            )
            
            return {
                "embeddings": embeddings,
                "count": len(embeddings),
                "timestamp": datetime.now().isoformat(),
                "status": "success"
            }
        except Exception as e:
            logger.error(f"Error generating embeddings: {e}")
            return {
                "embeddings": [],
                "count": 0,
                "timestamp": datetime.now().isoformat(),
                "status": "error",
                "error": str(e)
            }
    
    async def analyze_sentiment(self, text: str) -> Dict[str, Any]:
        """Analyze sentiment of text"""
        try:
            # Use OpenAI for sentiment analysis
            response = await asyncio.to_thread(
                openai.Completion.create,
                engine="text-davinci-003",
                prompt=f"Analyze the sentiment of this text and provide a score from -1 (very negative) to 1 (very positive):\n\n{text}\n\nSentiment analysis:",
                max_tokens=50,
                temperature=0.3
            )
            
            return {
                "text": text,
                "sentiment_analysis": response.choices[0].text.strip(),
                "timestamp": datetime.now().isoformat(),
                "status": "success"
            }
        except Exception as e:
            logger.error(f"Error analyzing sentiment: {e}")
            return {
                "text": text,
                "sentiment_analysis": "neutral",
                "timestamp": datetime.now().isoformat(),
                "status": "error",
                "error": str(e)
            }

# Agent configurations
AGENT_CONFIGS = {
    "research_agent": {
        "name": "Research Agent",
        "description": "Specialized in research and information gathering",
        "tools": ["knowledge_search", "web_search"],
        "temperature": 0.3
    },
    "creative_agent": {
        "name": "Creative Agent", 
        "description": "Specialized in creative tasks and content generation",
        "tools": ["text_generation", "idea_generation"],
        "temperature": 0.8
    },
    "analytical_agent": {
        "name": "Analytical Agent",
        "description": "Specialized in data analysis and insights",
        "tools": ["data_analysis", "chart_generation"],
        "temperature": 0.2
    },
    "task_agent": {
        "name": "Task Agent",
        "description": "Specialized in task management and scheduling",
        "tools": ["task_scheduler", "calendar_integration"],
        "temperature": 0.1
    }
}

class AgentManager:
    """Manages multiple specialized agents"""
    
    def __init__(self, openai_api_key: str, qdrant_url: str = "http://localhost:6333"):
        self.openai_api_key = openai_api_key
        self.qdrant_url = qdrant_url
        self.agents = {}
        self._initialize_agents()
    
    def _initialize_agents(self):
        """Initialize all configured agents"""
        for agent_id, config in AGENT_CONFIGS.items():
            self.agents[agent_id] = LikeMindAI(
                openai_api_key=self.openai_api_key,
                qdrant_url=self.qdrant_url
            )
    
    async def route_to_agent(self, query: str, agent_type: Optional[str] = None) -> Dict[str, Any]:
        """Route query to appropriate agent"""
        if agent_type and agent_type in self.agents:
            return await self.agents[agent_type].generate_response(query)
        
        # Simple routing logic based on keywords
        query_lower = query.lower()
        
        if any(word in query_lower for word in ["research", "find", "search", "information"]):
            return await self.agents["research_agent"].generate_response(query)
        elif any(word in query_lower for word in ["create", "generate", "write", "creative"]):
            return await self.agents["creative_agent"].generate_response(query)
        elif any(word in query_lower for word in ["analyze", "data", "chart", "graph"]):
            return await self.agents["analytical_agent"].generate_response(query)
        elif any(word in query_lower for word in ["schedule", "task", "remind", "calendar"]):
            return await self.agents["task_agent"].generate_response(query)
        else:
            # Default to research agent
            return await self.agents["research_agent"].generate_response(query)

# Example usage
if __name__ == "__main__":
    import asyncio
    
    async def main():
        # Initialize AI system
        ai = LikeMindAI(
            openai_api_key="your-openai-key-here",
            qdrant_url="http://localhost:6333"
        )
        
        # Test semantic search
        result = await ai.semantic_search("What is machine learning?")
        print("Search result:", result)
        
        # Test response generation
        response = await ai.generate_response("Explain artificial intelligence")
        print("AI Response:", response)
    
    # Run the example
    # asyncio.run(main())
