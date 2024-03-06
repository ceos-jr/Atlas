"use client";

import React from "react";
import { useRouter } from "next/navigation";

const Home: React.FC = () => {
  const router = useRouter();

  const goToLoginPage = () => {
    router.push("/login");
  };

  const goToSignupPage = () => {
    router.push("/sign"); 
  };

  return (
    <div className="h-screen text-center font-quick bg-white text-black flex flex-col lg:flex-row tela">
      <div className="flex flex-col flex-grow items-center justify-evenly lg:w-[70%] lg:items-start lg:justify-start">
        <div className="lg:h-[10%] lg:pl-10 lg:pt-5">
          Atlas
        </div>
        <div className="flex gap-y-4 flex-col px-10 lg:h-2/4 lg:px-0 lg:items-start lg:justify-end lg:pt-10 lg:pb-14 lg:mb-4 lg:pl-10 lg:mt-28">
          <h1 className="text-secondary-orange text-5xl font-bold lg:text-[100px] lg:text-left lg:leading-[117.19px]">Fortaleça os laços do seu time.</h1>
          <p className="text-[20px] text-gray-600 px-10 lg:text-[35px] lg:px-0 lg:text-left">Otimize o desempenho da sua equipe com o Atlas.</p>
        </div>
        <div className="space-x-4 space-y-4 lg:pl-10">
          <button
            onClick={goToSignupPage}
            className="px-3 py-2 lg:px-4 lg:mr-5 lg:py-3 lg:rounded-[10px] bg-secondary-orange text-white rounded-2xl transition-transform transform hover:scale-110 hover:bg-white hover:text-secondary-orange hover:outline focus:outline-none focus:ring focus:ring-secondary-orange"
          >
            Cadastre-se
          </button>
          <button
            onClick={goToLoginPage}
            className="text-secondary-blue rounded-2xl underline transition-transform transform hover:scale-110">
            Faça Login
          </button>
        </div>
      </div>

      <div className="lg:w-1/3 w-full h-[100px] flex flex-row lg:hidden lg:flex-col lg:h-full">
        <div className="bg-primary-blue flex-1 lg:w-7/12 lg:h-1/4 lg:flex-none"></div>
        <div className="bg-secondary-blue flex-1 lg:w-7/12 lg:h-1/4 lg:flex-none"></div>
        <div className="bg-primary-orange flex-1 w-7/12 lg:h-1/4 lg:flex-none"></div>
        <div className="bg-primary-yellow flex-1 w-7/12 lg:h-1/4 lg:flex-none"></div>
      </div>

      <div className="lg:flex hidden w-[30%] h-full flex-col relative">
        <div className="h-[30%] flex flex-row justify-end">
          <div className="bg-primary-orange  h-[100%] w-[50%]"></div>
        </div>
        <div className="h-[30%] w-[100%] flex flex-row justify-start absolute top-[23%] left-[15%]">
          <div className="bg-primary-yellow h-[100%] w-[50%]"></div>
        </div>
        <div className="h-[30%] w-[100%] flex flex-row justify-end absolute top-[46%]">
          <div className="bg-primary-blue h-[100%] w-[50%]"></div>
        </div>
        <div className="h-[30%] w-[100%] flex flex-row justify-start absolute top-[70%] left-[15%]">
          <div className="bg-secondary-blue h-[100%] w-[50%]"></div>
        </div>
      </div>

    </div>
  );
};

export default Home;
