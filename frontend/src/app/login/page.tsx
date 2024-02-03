"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";

const LoginPage = () => {
  const router = useRouter();

  const goHomePage = () => {
    router.push("/");
  };

  const [searchString, setSearchString] = useState({
    name: "",
    email: "",
    whatsapp: "",
  });

  return (
    <div className="flex flex-col md:flex-row justify-center md:items-center h-screen text-center font-quick mx-4 md:mx-8 lg:mx-12">
        <p className="text-gray-700 text-3xl my-12">cadastro de usuário</p>
        <div className="card flex flex-col bg-gray-200 rounded-lg shadow-xl px-6 py-2">
          <input
            type="text"
            placeholder="Nome Completo"
            className="form-input rounded-lg border-2 my-2 py-3 px-2 text-base border border-gray-400"
            onChange={(e) => setSearchString({ ...searchString, name: e.target.value })}
          />
          <input
            type="text"
            placeholder="Email"
            className="form-input rounded-lg border-2 my-2 py-3 px-2 text-base border border-gray-400"
            onChange={(e) => setSearchString({ ...searchString, email: e.target.value })}
          />
          <input
            type="text"
            placeholder="Whatsapp"
            className="form-input rounded-lg border-2 my-2 py-3 px-2 text-base border border-gray-400"
            onChange={(e) => setSearchString({ ...searchString, whatsapp: e.target.value })}
          />
          <button
            onClick={goHomePage}
            className='px-4 py-2 w-1/3 mx-auto my-7 bg-primary-blue text-white rounded hover:bg-blue-600 focus:outline-none focus:ring focus:ring-blue-300 text-sm text-shadow'
          >
            Home
          </button>
        </div>
      </div>
  );
};

export default LoginPage;
