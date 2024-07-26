"use client"
import { Input } from "@/components/ui/input"
import { Popover, PopoverTrigger, PopoverContent } from "@/components/ui/popover"
import { Button } from "@/components/ui/button"
import { Command, CommandGroup, CommandItem } from "@/components/ui/command"
import { useState } from "react"
import PutPopup from "./PutPopup"
import CallsPopup from "./CallsPopup"
export default function Options() {
  const [showOptions, setShowOptions] = useState(false)
  return (
    <div className="bg-background rounded-lg border p-6 w-full max-w-3xl">
      <div className="grid grid-cols-2 gap-6">
        <div>
          <label htmlFor="ticker" className="block text-sm font-medium text-muted-foreground">
            Stock Ticker
          </label>
          <div className="relative mt-1">
            <Input id="ticker" placeholder="Enter ticker symbol" className="pr-10" />
            <div className="absolute inset-y-0 right-0 flex items-center pr-3 pointer-events-none">
              <SearchIcon className="h-5 w-5 text-muted-foreground" />
            </div>
          </div>
        </div>
        <div>
          <label htmlFor="contracts" className="block text-sm font-medium text-muted-foreground">
            Number of Contracts
          </label>
          <Input id="contracts" type="number" min="1" defaultValue="1" className="w-full" onChange={()=>setShowOptions(true)}/>
        </div>
      </div>
      <div className="mt-6 rounded-lg bg-muted p-4">
        <div className="flex items-center justify-between">
          <p className="text-sm font-medium">Total Cost:</p>
          <p className="text-lg font-bold">${showOptions? "654.00" : "0"}</p>
        </div>
        <div className="flex items-center justify-between mt-2">
          <p className="text-sm font-medium">Potential Profit/Loss:</p>
          <p className="text-lg font-bold text-green-500">+${showOptions? "200.00" : "0"}</p>
        </div>
      </div>
      <div className="mt-6 flex gap-2">
        <PutPopup />
        <CallsPopup />
      </div>
    </div>
  )
}

function ArrowRightIcon(props:any) {
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
      <path d="M5 12h14" />
      <path d="m12 5 7 7-7 7" />
    </svg>
  )
}


function ChevronsUpDownIcon(props:any) {
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
      <path d="m7 15 5 5 5-5" />
      <path d="m7 9 5-5 5 5" />
    </svg>
  )
}


function SearchIcon(props:any) {
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
      <circle cx="11" cy="11" r="8" />
      <path d="m21 21-4.3-4.3" />
    </svg>
  )
}


function XIcon(props:any) {
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
      <path d="M18 6 6 18" />
      <path d="m6 6 12 12" />
    </svg>
  )
}
