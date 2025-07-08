'use client'

import { useState, useEffect } from 'react'
import { 
  ChartBarIcon, 
  ChatBubbleLeftRightIcon, 
  DocumentTextIcon,
  CpuChipIcon,
  UsersIcon,
  BoltIcon
} from '@heroicons/react/24/outline'

interface DashboardStats {
  totalQueries: number
  activeAgents: number
  knowledgeBase: number
  responseTime: number
}

interface RecentActivity {
  id: string
  type: 'query' | 'agent' | 'knowledge'
  message: string
  timestamp: Date
}

export default function Dashboard() {
  const [stats, setStats] = useState<DashboardStats>({
    totalQueries: 0,
    activeAgents: 0,
    knowledgeBase: 0,
    responseTime: 0
  })
  
  const [recentActivity, setRecentActivity] = useState<RecentActivity[]>([])
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    // Simulate API calls
    setTimeout(() => {
      setStats({
        totalQueries: 1247,
        activeAgents: 5,
        knowledgeBase: 12890,
        responseTime: 0.8
      })
      
      setRecentActivity([
        {
          id: '1',
          type: 'query',
          message: 'User queried about machine learning algorithms',
          timestamp: new Date(Date.now() - 300000)
        },
        {
          id: '2',
          type: 'agent',
          message: 'New agent deployed for data analysis',
          timestamp: new Date(Date.now() - 600000)
        },
        {
          id: '3',
          type: 'knowledge',
          message: 'Knowledge base updated with 50 new documents',
          timestamp: new Date(Date.now() - 900000)
        }
      ])
      
      setLoading(false)
    }, 1000)
  }, [])

  const statCards = [
    {
      title: 'Total Queries',
      value: stats.totalQueries.toLocaleString(),
      icon: ChartBarIcon,
      color: 'bg-blue-500'
    },
    {
      title: 'Active Agents',
      value: stats.activeAgents.toString(),
      icon: CpuChipIcon,
      color: 'bg-green-500'
    },
    {
      title: 'Knowledge Base',
      value: stats.knowledgeBase.toLocaleString(),
      icon: DocumentTextIcon,
      color: 'bg-purple-500'
    },
    {
      title: 'Avg Response Time',
      value: `${stats.responseTime}s`,
      icon: BoltIcon,
      color: 'bg-yellow-500'
    }
  ]

  if (loading) {
    return (
      <div className="animate-pulse">
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          {[1, 2, 3, 4].map((i) => (
            <div key={i} className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
              <div className="h-4 bg-gray-200 rounded mb-4"></div>
              <div className="h-8 bg-gray-200 rounded"></div>
            </div>
          ))}
        </div>
        <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
          <div className="h-4 bg-gray-200 rounded mb-4"></div>
          <div className="space-y-3">
            {[1, 2, 3].map((i) => (
              <div key={i} className="h-12 bg-gray-200 rounded"></div>
            ))}
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className="space-y-6">
      {/* Stats Cards */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
        {statCards.map((card, index) => (
          <div key={index} className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-600">{card.title}</p>
                <p className="text-2xl font-bold text-gray-900 mt-1">{card.value}</p>
              </div>
              <div className={`p-3 rounded-full ${card.color}`}>
                <card.icon className="w-6 h-6 text-white" />
              </div>
            </div>
          </div>
        ))}
      </div>

      {/* Recent Activity */}
      <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
        <h3 className="text-lg font-semibold text-gray-900 mb-4">Recent Activity</h3>
        <div className="space-y-3">
          {recentActivity.map((activity) => (
            <div key={activity.id} className="flex items-center space-x-3 p-3 bg-gray-50 rounded-lg">
              <div className="flex-shrink-0">
                {activity.type === 'query' && <ChatBubbleLeftRightIcon className="w-5 h-5 text-blue-500" />}
                {activity.type === 'agent' && <CpuChipIcon className="w-5 h-5 text-green-500" />}
                {activity.type === 'knowledge' && <DocumentTextIcon className="w-5 h-5 text-purple-500" />}
              </div>
              <div className="flex-1">
                <p className="text-sm font-medium text-gray-900">{activity.message}</p>
                <p className="text-xs text-gray-500">{activity.timestamp.toLocaleTimeString()}</p>
              </div>
            </div>
          ))}
        </div>
      </div>

      {/* Quick Actions */}
      <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-200">
        <h3 className="text-lg font-semibold text-gray-900 mb-4">Quick Actions</h3>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          <button className="flex items-center justify-center space-x-2 p-4 bg-blue-50 text-blue-700 rounded-lg hover:bg-blue-100 transition-colors duration-200">
            <ChatBubbleLeftRightIcon className="w-5 h-5" />
            <span className="font-medium">Start Chat</span>
          </button>
          <button className="flex items-center justify-center space-x-2 p-4 bg-green-50 text-green-700 rounded-lg hover:bg-green-100 transition-colors duration-200">
            <CpuChipIcon className="w-5 h-5" />
            <span className="font-medium">Deploy Agent</span>
          </button>
          <button className="flex items-center justify-center space-x-2 p-4 bg-purple-50 text-purple-700 rounded-lg hover:bg-purple-100 transition-colors duration-200">
            <DocumentTextIcon className="w-5 h-5" />
            <span className="font-medium">Add Knowledge</span>
          </button>
        </div>
      </div>
    </div>
  )
}
