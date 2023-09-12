"use client";

//verificar o router
//revisar codigo linha por linha
import React, { useState } from "react";
import { useRouter } from "next/navigation";

const SignInPage = () => {
  const router = useRouter();
  
  const [data, setData] = useState({
    name: "",
    email: "",
  });

  // esse handle submit envia pro console .log, mas dps mudar isso
  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log(data);
  };
  
  return (
    <section className=" flex justify-center items-center mx-auto h-screen bg-gradient-to-r from-cyan-500 to-blue-500">
      <div className=" container flex md:mx-52 items-center justify-center my-20 py-6 rounded-se-full bg-white/40 ">
      <div>
       <div>
      <h1 className="text-4xl text-teal-800 pb-4">Cadastre-se</h1>
       </div>
      <form 
        className="flex flex-col space-y-8"
        onSubmit={(e) => handleSubmit(e)}
      >
        <input
          type="text"
          placeholder="Nome"
          className="form-input rounded-lg border-2"
          onChange={(e) => setData({ ...data, name: e.target.value })}
        />
        <input
          type="email"
          placeholder="Email para contato"
          className="form-input rounded-lg border-2"
          onChange={(e) => setData({ ...data, email: e.target.value })}
        />
        <input
          type="submit"
          value="Enviar"
          className="py-4 bg-stone-200 rounded-lg hover:bg-green-500 shadow-lg shadow-cyan-500/50 focus:outline-zinc-900 active:py-6z"
        />
      </form>
      </div>
      </div>
    </section>
  );
};

export default SignInPage;