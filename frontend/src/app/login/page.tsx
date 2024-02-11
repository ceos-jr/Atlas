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
        <p className="text-gray-700 text-3xl my-12 color-[#727272]">cadastro de usuário</p>
        <div className="card flex flex-col bg-[#f4f9ff] rounded-lg shadow-[0px_4px_4px_0px_#00000040] px-6 py-4 mx-6">
          <input
            type="text"
            placeholder="Nome Completo"
            className="form-input rounded-lg my-3 py-3 px-2 text-base border border-gray-400"
            onChange={(e) => setSearchString({ ...searchString, name: e.target.value })}
          />
          <input
            type="text"
            placeholder="Email"
            className="form-input rounded-lg my-3 py-3 px-2 text-base border border-gray-400"
            onChange={(e) => setSearchString({ ...searchString, email: e.target.value })}
          />
          <input
            type="text"
            placeholder="Whatsapp"
            className="form-input rounded-lg my-3 py-3 px-2 text-base border border-gray-400"
            onChange={(e) => setSearchString({ ...searchString, whatsapp: e.target.value })}
          />
          <button
            onClick={goHomePage}
            className='px-4 py-2 w-2/5 mx-auto my-7 bg-[#ff8a00] text-white rounded-[10px] hover:bg-blue-600 focus:outline-none focus:ring focus:ring-blue-300 text-sm text-shadow-[0px_4px_4px_0px_#00000040]'
          >
            Cadastrar
          </button>
        </div>
      </div>
  );
};

export default LoginPage;
