"use client"

import React, { ReactNode } from 'react';
import { useRouter } from 'next/navigation';



const LoginPage = () => {
  const router = useRouter();

  const goHomePage = () => {
    router.push('/');
  };

  return (
    <div className="flex justify-center items-center h-screen text-center font-quick">
      <div className="card bg-white rounded-lg shadow-md p-6">
          <h1 className="text-3xl font-semibold mb-4 text-primary-blue">Teste Login</h1>
          <p className="text-lg text-gray-600"> Uma plataforma para gerenciamento de demandas e equipes. </p>
          <button onClick = {goHomePage}className='className="px-4 py-2 bg-primary-blue text-white rounded hover:bg-blue-600 focus:outline-none focus:ring focus:ring-blue-300 text-xs"'> Home </button>
      </div>
    </div>
  );
};


export default LoginPage;