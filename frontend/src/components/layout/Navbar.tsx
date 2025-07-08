'use client'

import { BellIcon, UserCircleIcon } from '@heroicons/react/24/outline'

export function Navbar() {
  return (
    <header className="bg-white shadow-sm border-b border-gray-200 fixed top-0 right-0 left-64 z-40">
      <div className="px-6 py-4">
        <div className="flex items-center justify-between">
          <div className="flex items-center space-x-4">
            <h2 className="text-lg font-semibold text-gray-900">
              AI Knowledge Platform
            </h2>
          </div>
          
          <div className="flex items-center space-x-4">
            <button className="p-2 text-gray-500 hover:text-gray-700 hover:bg-gray-100 rounded-lg transition-colors duration-200">
              <BellIcon className="w-5 h-5" />
            </button>
            
            <div className="flex items-center space-x-3">
              <UserCircleIcon className="w-8 h-8 text-gray-500" />
              <span className="text-sm font-medium text-gray-700">Admin User</span>
            </div>
          </div>
        </div>
      </div>
    </header>
  )
}
