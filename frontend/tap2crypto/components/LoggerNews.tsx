import React from "react";
const LoggerNews = ({ LoggerStatements }: { LoggerStatements: string[] }) => {
  return (
    <div className="flex flex-col min-h-screen bg-gray-900 text-zinc-100">
      <header className="flex items-center justify-between px-4 py-2 border-b border-zinc-700">
        <h1 className="text-xl font-bold">CLI Logger</h1>
      </header>
      <main className="flex-1 overflow-auto p-4">
        <div className="space-y-2">
        {
          LoggerStatements.map((statement, index) => (
            <div key={index} className="flex items-center gap-2 text-green-400">
              <CircleIcon className="w-4 h-4" />
              <span className="font-mono">{statement}</span>
            </div>
          ))
        }
        </div>
      </main>
    </div>
  )
}

export default LoggerNews

function CircleIcon(props:any) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      width="24"
      height="24"
      viewBox="0 0 24 24"
      fill="none"
      stroke="currentColor"
      strokeWidth="2"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <circle cx="12" cy="12" r="10" />
    </svg>
  )
}

