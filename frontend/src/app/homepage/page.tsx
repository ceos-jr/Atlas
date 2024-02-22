"use client"
import React from "react";

export default function Page() {
    return (
    
<div className="grid grid-cols-3 auto-rows-[300px] gap-4">
    <div className="p-2 flex flex-col items-center justify-center col-span-2 row-[200px]">3</div>
    <div className="bg-neutral-100 border-2 rounded-xl p-2 flex flex-col items-center justify-center col-span-1.5">4</div>
    <div className="bg-neutral-100 border-2 rounded-xl p-2 flex flex-col items-center justify-center">5</div>
    <div className="bg-neutral-100 border-2 rounded-xl p-2 flex flex-col items-center justify-center">6</div>
    <div className="bg-neutral-100 border-2 rounded-xl row-span-3 p-2 flex flex-col items-center justify-center">7</div>
    <div className="bg-neutral-100 border-2 rounded-xl row-span-2 col-span-2 p-2 flex flex-col items-center justify-center">8</div>
</div>
    )
  }

