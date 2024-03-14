"use client";
import React from "react";
import nextimage from "next/image";

export default function Homepage() {
  return (
    <div className="grid grid-cols-2 md:grid-cols-2 lg:grid-cols-3 auto-rows-[300px] gap-4">
      <div className="p-2 flex items-center col-span-2">
        <div className="m-4 w-36 h-36 bg-white shadow-lg rounded-full "></div>
        <div className="ml-4">
          <h2 className="text-4xl text-secondary-blue m-4">Olá,</h2>
          <p className="text-5xl font-light text-secondary-blue m-4 hover:underline decoration-1 cursor-pointer">Antônio Rego.</p>
        </div>
      </div>

      <div className="bg-neutral-100 border-2 rounded-xl p-2 flex flex-col items-center justify-center relative">
        <div className="md:m-4 w-44 h-44 bg-white border border-blue-600 shadow-lg rounded-full overflow-hidden">
          <div className="w-full h-full flex items-center justify-center">
            <h1 className="text-black">100 </h1>
          </div>
        </div>
      </div>

      <div className=" bg-neutral-100 border-2 rounded-xl p-2 flex flex-col items-center justify-center ">
        <div className="ml-4">
          <h1 className="text-10xl font-bold">Segunda-feira</h1>
          <p>Your additional content goes here.</p>
        </div>
      </div>

      <div className="col-span-2 lg:col-span-1 bg-neutral-100 border-2 rounded-xl p-2 flex flex-col items-center justify-center">
        <div className="ml-4">
          <h2 className="text-2xl font-bold ">Your Text Here</h2>
          <h2 className="text-2xl font-bold ">Your Text Here</h2>
          <h2 className="text-2xl font-bold ">Your Text Here</h2>
        </div>
      </div>


      <div className="col-span-2 lg:col-span-1 bg-neutral-100 border-2 rounded-xl row-span-3 p-4 flex flex-col gap-6 relative">
        <div className="flex-grow flex flex-col justify-end gap-4">
        <div className="relative w-full h-36 p-2 bg-white rounded-md shadow-md">
        <span className="absolute inset-y-0 left-0 bg-blue-400 w-4 "></span>
          <div className="flex -space-x-2 rtl:space-x-reverse overflow-hidden absolute bottom-0 left-0 m-2">
            <img className="w-10 h-10 border-2 border-white rounded-full dark:border-gray-800 " src="/img1.jpg" alt="img1" />
            <img className="w-10 h-10 border-2 border-white rounded-full dark:border-gray-800" src="/img1.jpg" alt="img1" />
            <img className="w-10 h-10 border-2 border-white rounded-full dark:border-gray-800" src="/img1.jpg" alt="img1" />
          </div>
          </div>
        <div className=" w-full h-36 p-2 bg-white rounded-md shadow-md" />
        <div className=" w-full h-36 p-2 bg-white rounded-md shadow-md" />
        <div className=" w-full h-36 p-2 bg-white rounded-md shadow-md" />
      </div>
      </div>

      <div className="bg-neutral-100 border-2 rounded-xl row-span-2 col-span-2 p-16 flex flex-col gap-4 ">
        <div className="flex items-center gap-10 bg-neutral-100 rounded hover:shadow-lg cursor-pointer ">
          <div className="flex aspect-square w-10 m-2 rounded bg-blue-600 "></div>
          <div className="text-md text-gray-400 "> Documentação Atlas
          </div>
          <div className="text-md text-gray-400 "> Deadline: minha bunda
          </div>
        </div>
      </div>

    </div>
  );
}
