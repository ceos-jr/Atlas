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
    <div className="flex justify-center items-center h-screen text-center font-quick">
      <div className="card bg-white rounded-lg shadow-md p-6">
        <h1 className="text-3xl font-semibold mb-4 text-primary-blue">
          Boas vindas ao Atlas!
        </h1>
        <p className="text-lg text-gray-600">
          Uma plataforma para gerenciamento de demandas e equipes.
        </p>
        <button
          onClick={goToLoginPage}
          className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 focus:outline-none focus:ring focus:ring-blue-300 mb-2"
        >
          Login
        </button>
        <button
          onClick={goToSignupPage}
          className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600 focus:outline-none focus:ring focus:ring-blue-300"
        >
          Cadastre-se
        </button>
      </div>
    </div>
  );
};

export default Home;
