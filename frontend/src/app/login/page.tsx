"use client";

import React from "react";
import { useRouter } from "next/navigation";
import Image from 'next/image';
import smallLogin from './images/small-login-image.svg';
import largeLogin from './images/large-login-image.svg'

const LoginPage = () => {
  const router = useRouter();

  const goHomePage = () => {
    router.push("/");
  };

  return (
    <div className="bg-white flex justify-center lg:justify-normal  h-screen text-center font-quick">
      <div className="absolute top-10 mt-10 lg:absolute  lg:top-40 lg:right-0 lg:w-[45%] xl:w-[45%] xl:top-[25%]">
        <h1 className="text-secondary-gray font-roboto-bold">
          login de usuário
        </h1>
        <div className="md:flex md:justify-center md:items-center">
          <form className="bg-primary-gray flex flex-col items-center justify-center mx-4 rounded-lg p-8 shadow-xl mt-8 md:absulute lg:w-full lg:ml-2 xl:ml-[5%] xl:mr-[10%] ">
            <label className="mb-4 w-full">
              <input 
                type="text" 
                placeholder="Email"
                className=" p-4 border border-gray-300 rounded-lg w-full text-base text-secondary-gray font-roboto"
                />
            </label>
            <label className="mb-4 w-full">
              <input 
                type="text" 
                placeholder="Senha"
                className=" p-4 border border-gray-300 rounded-lg w-full mt-5 text-base text-secondary-gray font-roboto"
                />
            </label>
            <button 
              type="submit" 
              className="font-def bg-terciary-blue text-white p-2 rounded-lg  w-1/2 mt-5 mb-6 text-lg font-roboto-bold "
              >
              entrar  
            </button>
          </form>
        </div>
          
        <div>
          <img src={smallLogin} alt="" className="lg:hidden w-full absolute mt-20"/>
        </div>

      </div>

      <div className="lg:flex lg:w-[55%] lg:top-10 lg:absolute xl:w-full xl:top-10">
        <img src={largeLogin} alt="" className="hidden lg:flex"/>
      </div>
               
    </div>
  );
};

export default LoginPage;
