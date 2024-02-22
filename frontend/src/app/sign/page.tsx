"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import Image from "next/image";
import smGreet from "./images/Group.png";
import greet from "./images/Welcome.png";
import {z, ZodType} from "zod";
import { zodResolver } from '@hookform/resolvers/zod';
import {useForm} from "react-hook-form";

type formData = {
  userName: string,
  email: string,
  whatsapp: string,
}

const SignInPage = () => {
  const router = useRouter();

   const [hover, setHover] = useState(false);

  const handleMouseEnter = () => {
    setHover(true);
  };

  const handleMouseLeave = () => {
    setHover(false);
  };

  const whatsappRegex = new RegExp(
    /^(\d{1,3}|[(]+[0-9]+[)])([\s]?[0-9]){9,10}$/
  )
  
  const formSchema: ZodType<formData> = z.object({ 
    userName: z.string().
      min(3, {message: "Nome muito curto!"}).
      max(128, {message: "Nome muito grande!"}),
  
    email: z.string().
      min(3, {message: "email muito curto!"}).
      max(128, {message: "email muito grande!"}).
      email({message: "Não é um email Válido!"}),
  
    whatsapp: z.string().regex(whatsappRegex, {message: "Não é um número válido!"}),
  })

  const {
    register,
    handleSubmit,
    formState: { errors },
    getValues,
  } = useForm<formData>({
    resolver: zodResolver(formSchema),
  });

  return (
    <div className="flex flex-col md:flex-row justify-between items-center h-screen text-center font-quick mx-4 md:mx-0 lg:mx-12">
      <div className="flex flex-col items-center mt-10 w-5/6 md:w-1/2 mb-12 l:h-full">
          <p 
            className="text-gray-700 text-30 leading-35 mt-20 mb-10 color-[#727272] :hidden font-bold md:mb-3 xl:mb-10 xl:mt-2"
            style={{textShadow: "0px 4px 4px rgba(0, 0, 0, 0.25)"}}
            >cadastro de usuário</p>
          
          <div className="card flex flex-col justify-around bg-[#f4f9ff] rounded-lg shadow-[0px_4px_4px_0px_#00000040] px-6 py-4 w-full xl:w-2/3 md:mx-5">
            <form onSubmit={handleSubmit(() => createUser("0", getValues("userName"), getValues("email"), "1"))}>
              <input
                type="text"
                placeholder="Nome completo"
                className="form-input rounded-lg my-3 py-4 px-2  text-lg border border-gray-400 color-[#727272]"
                {...register('userName')}
              />
              {errors.userName?.message && <p className="text-sm font-bold color-[#ab0303]">{errors.userName?.message}</p>}
              <input
                type="text"
                placeholder="Email"
                className="form-input rounded-lg my-3 py-4 px-2 text-lg border border-gray-400"
                {...register('email')}
              />
              {errors.email?.message && <p className="text-sm font-bold color-[#ab0303]">{errors.email?.message}</p>}
              <input
                type="text"
                placeholder="Número do whatsapp"
                className="form-input rounded-lg my-3 py-4 px-2 text-lg border border-gray-400"
                {...register('whatsapp')}
              />
              {errors.whatsapp?.message && <p className="text-sm font-bold color-[#ab0303]">{errors.whatsapp?.message}</p>}
              <button
                type="submit"
                className='px-4 py-2 w-1/2 xl:w-1/2 xl:py-5 xl:scale-5 mx-auto my-7 text-white rounded-[10px] text-center hover:bg-blue-600 focus:outline-none focus:ring focus:ring-blue-300 text-lg xl:text-3xl font-bold'
                style={{textShadow: "0px 4px 4px rgba(0, 0, 0, 0.25)", backgroundColor : hover ? '#ff3d00' : '#ff8a00'}}
                onMouseEnter={handleMouseEnter}
                onMouseLeave={handleMouseLeave}
              >
                cadastrar
              </button>
            </form>
          </div>
        </div>
        {/*imagem mobile*/}
        <Image 
          src={smGreet}
          alt="boas vindas ao ATLAS"
          width={360}
          height={430}
          layout="responsive"
          className="md:hidden"
          />
          {/*imagem média e grande*/}
        <Image 
          src={greet}
          alt="boas vindas ao ATLAS"
          width={660}
          height={630}
          layout="responsive"
          className="hidden md:block xl:hidden"
          />
          {/*imagem da tela do computador*/}
        <Image 
          src={greet}
          alt="boas vindas ao ATLAS"
          width={800}
          height={800}
          className="hidden xl:block"
          />
      </div>
  );
};

export default SignInPage;
