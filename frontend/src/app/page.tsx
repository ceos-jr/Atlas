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
    <div className="h-screen text-center font-quick bg-white text-black flex flex-col lg:flex-row">
      <div className="flex flex-col flex-grow items-center justify-evenly">
        <div>
          Atlas
        </div>
        <div className="flex gap-y-4 flex-col px-10">
          <h1 className="text-secondary-orange text-5xl font-bold lg:text-[70px]">Fortaleça os laços do seu time.</h1>
          <p className="text-[20px] text-gray-700 px-10 lg:text-[25px]">Otimize o desempenho da sua equipe com o Atlas.</p>
        </div>
        <div className="space-x-4 space-y-4">
          <button
            onClick={goToSignupPage}
            className="px-3 py-2 bg-secondary-orange text-white rounded-2xl transition-transform transform hover:scale-110 hover:bg-white hover:text-secondary-orange hover:outline focus:outline-none focus:ring focus:ring-secondary-orange"
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

      <div className="lg:w-1/3 w-full h-[100px] flex flex-row lg:flex lg:flex-col lg:h-full">
        <div className="bg-primary-blue flex-1 lg:w-7/12 lg:h-1/4 lg:flex-none"></div>
        <div className="bg-secondary-blue flex-1 lg:w-7/12 lg:h-1/4 lg:flex-none"></div>
        <div className="bg-primary-orange flex-1 w-7/12 lg:h-1/4 lg:flex-none"></div>
        <div className="bg-primary-yellow flex-1 w-7/12 lg:h-1/4 lg:flex-none"></div>
      </div>
    </div>
  );
};

export default Home;
