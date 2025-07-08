'use client'

import { useState } from 'react'
import { MagnifyingGlassIcon, DocumentTextIcon, ClockIcon } from '@heroicons/react/24/outline'

interface SearchResult {
  id: string
  title: string
  content: string
  similarity: number
  source: string
  timestamp: Date
}

export default function SemanticSearch() {
  const [query, setQuery] = useState('')
  const [results, setResults] = useState<SearchResult[]>([])
  const [isLoading, setIsLoading] = useState(false)
  const [searchPerformed, setSearchPerformed] = useState(false)

  const handleSearch = async () => {
    if (!query.trim()) return

    setIsLoading(true)
    setSearchPerformed(true)

    // Simulate API call
    setTimeout(() => {
      const mockResults: SearchResult[] = [
        {
          id: '1',
          title: 'Machine Learning Fundamentals',
          content: 'Machine learning is a subset of artificial intelligence that enables computers to learn and improve from experience without being explicitly programmed...',
          similarity: 0.95,
          source: 'knowledge_base',
          timestamp: new Date('2024-01-15')
        },
        {
          id: '2',
          title: 'Neural Network Architecture',
          content: 'Neural networks are computing systems inspired by biological neural networks. They consist of layers of interconnected nodes that process information...',
          similarity: 0.87,
          source: 'documents',
          timestamp: new Date('2024-01-10')
        },
        {
          id: '3',
          title: 'Deep Learning Applications',
          content: 'Deep learning has revolutionized many fields including computer vision, natural language processing, and speech recognition...',
          similarity: 0.82,
          source: 'research_papers',
          timestamp: new Date('2024-01-08')
        }
      ]
      setResults(mockResults)
      setIsLoading(false)
    }, 1500)
  }

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter') {
      e.preventDefault()
      handleSearch()
    }
  }

  const getSimilarityColor = (similarity: number) => {
    if (similarity >= 0.9) return 'text-green-600'
    if (similarity >= 0.8) return 'text-yellow-600'
    return 'text-red-600'
  }

  const getSimilarityBadge = (similarity: number) => {
    if (similarity >= 0.9) return 'bg-green-100 text-green-800'
    if (similarity >= 0.8) return 'bg-yellow-100 text-yellow-800'
    return 'bg-red-100 text-red-800'
  }

  return (
    <div className="space-y-6">
      {/* Search Header */}
      <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
        <h1 className="text-2xl font-bold text-gray-900 mb-2">Semantic Search</h1>
        <p className="text-gray-600 mb-6">
          Search through your knowledge base using natural language queries with AI-powered semantic understanding
        </p>

        {/* Search Input */}
        <div className="flex items-center space-x-2">
          <div className="flex-1 relative">
            <MagnifyingGlassIcon className="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 text-gray-400" />
            <input
              type="text"
              value={query}
              onChange={(e) => setQuery(e.target.value)}
              onKeyPress={handleKeyPress}
              placeholder="Ask a question or describe what you're looking for..."
              className="w-full pl-10 pr-4 py-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-primary-500 focus:border-transparent"
              disabled={isLoading}
            />
          </div>
          <button
            onClick={handleSearch}
            disabled={!query.trim() || isLoading}
            className="px-6 py-3 bg-primary-600 text-white rounded-lg hover:bg-primary-700 disabled:opacity-50 disabled:cursor-not-allowed transition-colors duration-200"
          >
            {isLoading ? 'Searching...' : 'Search'}
          </button>
        </div>
      </div>

      {/* Search Results */}
      {searchPerformed && (
        <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-lg font-semibold text-gray-900">
              Search Results {results.length > 0 && `(${results.length})`}
            </h2>
            {results.length > 0 && (
              <span className="text-sm text-gray-500">
                Sorted by relevance
              </span>
            )}
          </div>

          {isLoading ? (
            <div className="space-y-4">
              {[1, 2, 3].map((i) => (
                <div key={i} className="animate-pulse">
                  <div className="h-4 bg-gray-200 rounded mb-2"></div>
                  <div className="h-3 bg-gray-200 rounded mb-2 w-3/4"></div>
                  <div className="h-3 bg-gray-200 rounded w-1/2"></div>
                </div>
              ))}
            </div>
          ) : results.length > 0 ? (
            <div className="space-y-4">
              {results.map((result) => (
                <div key={result.id} className="border border-gray-200 rounded-lg p-4 hover:border-primary-300 transition-colors duration-200">
                  <div className="flex items-start justify-between mb-2">
                    <h3 className="text-lg font-medium text-gray-900">{result.title}</h3>
                    <span className={`px-2 py-1 rounded-full text-xs font-medium ${getSimilarityBadge(result.similarity)}`}>
                      {(result.similarity * 100).toFixed(0)}% match
                    </span>
                  </div>
                  
                  <p className="text-gray-600 mb-3 line-clamp-3">
                    {result.content}
                  </p>
                  
                  <div className="flex items-center justify-between text-sm text-gray-500">
                    <div className="flex items-center space-x-4">
                      <div className="flex items-center space-x-1">
                        <DocumentTextIcon className="w-4 h-4" />
                        <span className="capitalize">{result.source.replace('_', ' ')}</span>
                      </div>
                      <div className="flex items-center space-x-1">
                        <ClockIcon className="w-4 h-4" />
                        <span>{result.timestamp.toLocaleDateString()}</span>
                      </div>
                    </div>
                    <button className="text-primary-600 hover:text-primary-700 font-medium">
                      View Full Document
                    </button>
                  </div>
                </div>
              ))}
            </div>
          ) : (
            <div className="text-center py-8">
              <MagnifyingGlassIcon className="w-12 h-12 text-gray-400 mx-auto mb-4" />
              <h3 className="text-lg font-medium text-gray-900 mb-2">No results found</h3>
              <p className="text-gray-500">
                Try adjusting your search query or using different keywords
              </p>
            </div>
          )}
        </div>
      )}

      {/* Search Tips */}
      {!searchPerformed && (
        <div className="bg-blue-50 p-6 rounded-lg border border-blue-200">
          <h3 className="text-lg font-medium text-blue-900 mb-3">Search Tips</h3>
          <ul className="space-y-2 text-blue-800">
            <li>• Use natural language to describe what you're looking for</li>
            <li>• Ask questions like "How does machine learning work?"</li>
            <li>• Search for concepts, not just keywords</li>
            <li>• The AI understands context and semantics</li>
          </ul>
        </div>
      )}
    </div>
  )
}
