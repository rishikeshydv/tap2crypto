/**
 * v0 by Vercel.
 * @see https://v0.dev/t/DK84tBjCr6O
 * Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
 */
import React from "react"
import { Button } from "@/components/ui/button"
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogDescription, DialogFooter } from "@/components/ui/dialog"
import { useState } from "react"

export default function CallsPopup() {
  const [isOpen, setIsOpen] = useState(false)
  return (
    <>
      <Button className="flex-1" onClick={() => setIsOpen(true)}>
          Buy Call
          <ArrowRightIcon className="ml-2 h-4 w-4" />
        </Button>
      {isOpen && (
        <Dialog open={isOpen} onOpenChange={setIsOpen}>
          <DialogContent className="sm:max-w-[425px]">
            <DialogHeader>
              <DialogTitle>Confirmation</DialogTitle>
            </DialogHeader>
            <div className="grid gap-4 py-4">
              <DialogDescription>Your options have been purchased</DialogDescription>
            </div>
            <DialogFooter>
              <Button onClick={() => setIsOpen(false)}>OK</Button>
            </DialogFooter>
          </DialogContent>
        </Dialog>
      )}
    </>
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