'use client'
import React from 'react';
import { HiMenuAlt3 } from 'react-icons/hi';
import {Link} from 'react-router-dom';
import { BiHomeAlt2 } from "react-icons/bi";
import { AiOutlineUser } from "react-icons/ai"; 
import { RiTeamLine } from "react-icons/ri";
import { CgGames } from "react-icons/cg";
import { GrConfigure } from "react-icons/gr";
import { IoIosLogOut } from 'react-icons/io';

const SideBar = () => {
  return (

    <section className="flex gap-10">
      <div className="bg-white min-h-screen w-72 px-4">
        <div className="py-10 m-4 flex justify-end">
          <HiMenuAlt3 size={26} className="cursor-pointer bg-black" />
        </div>
        <div>
          <a
            href="#logo"
            className="flex aspect-square min-h-[32px] w-24 flex-col items-center justify-center gap-2 m-10 rounded-md p-10 bg-primary-gray shadow-md"
        />
        </div>
        <div>
        <a
            href="#homepage"
            className="flex aspect-square min-h-[32px] w-24 flex-col items-center justify-center gap-2 m-10 rounded-md p-10 bg-primary-gray text-gray-700 hover:bg-blue-300 shadow-md"
        >
          <div> 
            <BiHomeAlt2 size={26} />
          </div>
            <small className="text-center text-xs font-medium"> Home </small>
        </a>

        </div>
        <div>
        <a
            href="#meuperfil"
            className="flex aspect-square min-h-[32px] w-24 flex-col items-center justify-center gap-2 m-10 rounded-md p-10 bg-primary-gray text-gray-700 hover:bg-blue-300 shadow-md"
        >
          <div>
            <AiOutlineUser size={26} />
          </div>
            <small className="text-center text-xs font-medium"> Meu Perfil </small>
        </a>

        </div>
        <div>
        <a
            href="#equipes"
            className="flex aspect-square min-h-[32px] w-24 flex-col items-center justify-center gap-2 m-10 rounded-md p-10 bg-primary-gray text-gray-700 hover:bg-blue-300 shadow-md"
        >
          <div>
              <RiTeamLine size={26} />
            </div>
            <small className="text-xs font-medium">Equipes</small>
            
        </a>
        </div>
        <div>
        <a
            href="#meuspontos"
            className="flex aspect-square min-h-[32px] w-24 flex-col items-center justify-center gap-2 m-10 rounded-md p-10 bg-primary-gray text-gray-700 hover:bg-blue-300 shadow-md"
        >
          <div>
            <CgGames size={26} />
          </div>
            <small className="text-center text-xs font-medium"> Meus pontos </small>
        </a>
        </div>
        <div>
        <a
            href="#config"
            className="flex aspect-square min-h-[32px] w-24 flex-col items-center justify-center gap-2 m-10 rounded-md p-10 bg-primary-gray text-gray-700 hover:bg-blue-300 shadow-md"
        >
          <div>
            <GrConfigure size={26} />
          </div>
            <small className="text-center text-xs font-medium"> Configurações </small>
        </a>
        </div>
        <div>
        <a
            href="#logout"
            className="flex aspect-square min-h-[32px] w-24 flex-col items-center justify-center gap-2 m-10 rounded-md p-10 bg-primary-gray text-gray-700 hover:bg-blue-300 shadow-md"
        >
            <div>
              <IoIosLogOut size={26} />
            </div>
            <small className="text-center text-xs font-medium"> Sair </small>
        </a>
        </div>
      </div>
    </section>
  );
};

export default SideBar;