"use client";

import React from "react";
import { useRouter } from "next/navigation";

const LoginPage = () => {
  const router = useRouter();

  const goHomePage = () => {
    router.push("/");
  };

  return (
    <div className="bg-white flex justify-center items-center h-screen text-center font-quick">
      <div className="absolute inset-x-0 top-10">
        <h1 className="text-black">
          login de usuário
        </h1>
      </div>
      <form className="bg-neutral-100 flex flex-col items-center justify-center mx-4  sm:w-2/3 md:w-1/2 lg:w-1/3 rounded-lg p-8 ">
        <label className="mb-4 w-full">
          <input 
            type="text" 
            placeholder="Email"
            className=" p-4 border border-gray-300 rounded-lg w-full text-base text-zinc-500"
          />
        </label>
        <label className="mb-4 w-full">
          <input 
            type="text" 
            placeholder="Senha"
            className=" p-4 border border-gray-300 rounded-lg w-full mt-5 text-base text-black"
          />
        </label>
        <button 
          type="submit" 
          className="font-def bg-sky-600 text-white p-1 rounded-lg hover:bg-cyan-700 w-1/2 mt-5 mb-6 text-lg "
        >
          entrar  
        </button>

          
      </form>
      
    </div>
  );
};

export default LoginPage;
