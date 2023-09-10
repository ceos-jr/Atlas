"use client";

//verificar o router
//revisar codigo linha por linha
import React, { useState } from "react";
import { useRouter } from "next/navigation";

const SignInPage = () => {
  const router = useRouter();
  
  const [data, setData] = useState({
    name: "",
    age: "",
    tel: "",
    email: "",
  });

  // esse handle submit envia pro console .log, mas dps mudar isso
  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log(data);
  };
  
  return (
    <section className=" container flex flex-col gap-8 justify-center items-center mx-auto h-screen">
      <h1 className="text-4xl text-teal-800">Cadastre-se</h1>
      <p className="text-neutral-400">Preecha os campos abaixo</p>
      <form 
        className="flex flex-col space-y-8 w-1/2"
        onSubmit={(e) => handleSubmit(e)}
      >
        <input
          type="text"
          placeholder="Nome"
          className="form-input rounded-lg border-2"
          onChange={(e) => setData({ ...data, name: e.target.value })}
        />
        <input
          type="tel"
          placeholder="Telefone para contato"
          className="form-input rounded-lg border-2"
          onChange={(e) => setData({ ...data, tel: e.target.value })}
        />
        <div>
        <input 
          type="checkbox"
          className="form-input rounded-lg"
          checked
           />
          <label className=""> Eu li e concordo com os termos de uso</label>
          </div>
        <input
          type="submit"
          value="Enviar"
          className="py-4 bg-stone-200 rounded-lg hover:bg-green-500 focus:outline-zinc-900 active:py-6z"
        />
      </form>
    </section>
  );
};

export default SignInPage;

// erro: 
//You provided a `checked` prop to a form field without an `onChange` handler. This will render a read-only field. If the field should be mutable use `defaultChecked`. Otherwise, set either `onChange` or `readOnly`.