"use client"
import React from "react";

export default function Page() {
    return (
    
<div className="grid grid-cols-3 auto-rows-[300px] gap-4">
    <div className="p-2 flex items-center col-span-2">
      <div className="m-4 w-36 h-36 bg-white shadow-lg rounded-full ">
      </div>
      <div className="ml-4">
      <h2 className="text-2xl font-bold">Your Text Here</h2>
      <p>Your additional content goes here.</p>
    </div>
    </div>
    <div className="bg-neutral-100 border-2 rounded-xl p-2 flex flex-col items-center justify-center ">
    <div className="m-4 w-44 h-44 bg-white border border-blue-600 shadow-lg rounded-full flex items-center justify-center">
      <h1 className="text-black">100</h1>
    </div>
    </div>
    <div className="bg-neutral-100 border-2 rounded-xl p-2 flex flex-col items-center justify-center">
    <div className="ml-4">
      <h2 className="text-2xl font-bold ">Your Text Here</h2>
      <p>Your additional content goes here.</p>
    </div>
    </div>
    <div className="bg-neutral-100 border-2 rounded-xl p-2 flex flex-col items-center justify-center">
    <p>Your additional content goes here.</p>
    <p>Your additional content goes here.</p>
    <p>Your additional content goes here.</p>
    </div>
    <div className="bg-neutral-100 border-2 rounded-xl row-span-3 p-4 py-8 flex flex-col items-center gap-6 ">
    <div className=" w-full h-36 p-2 bg-white rounded-md shadow-md"/>
    <div className=" w-full h-36 p-2 bg-white rounded-md shadow-md"/>
    <div className=" w-full h-36 p-2 bg-white rounded-md shadow-md"/>
    <div className=" w-full h-36 p-2 bg-white rounded-md shadow-md"/>
    <div className=" w-full h-36 p-2 bg-white rounded-md shadow-md"/>

    </div>
    <div className="bg-neutral-100 border-2 rounded-xl row-span-2 col-span-2 p-6 flex flex-col gap-4 ">
      <div className=" w-full h-12 p-4 bg-neutral-100 rounded-md shadow-md flex">
        <h1 className="text-2xl font-bold ">Your Text Here</h1>
      </div>
      <div className=" w-full h-12 p-2 bg-neutral-100 rounded-md hover: shadow-md"/>
      <div className=" w-full h-12 p-2 bg-neutral-100 rounded-md shadow-md"/>
      <div className=" w-full h-12 p-2 bg-neutral-100 rounded-md shadow-md"/>
      <div className=" w-full h-12 p-2 bg-neutral-100 rounded-md shadow-md"/>

    </div>
</div>

    )
};

