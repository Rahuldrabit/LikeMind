import Dashboard from '@/components/dashboard/Dashboard'

export default function Home() {
  return (
    <div className="space-y-6">
      <div>
        <h1 className="text-3xl font-bold text-gray-900">
          Welcome to LikeMind
        </h1>
        <p className="mt-2 text-gray-600">
          Your intelligent AI-powered knowledge management platform
        </p>
      </div>
      <Dashboard />
    </div>
  )
}
