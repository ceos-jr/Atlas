"use client";


import Link from "next/link";
import { FaWhatsapp } from "react-icons/fa";
import { CgProfile } from "react-icons/cg";
import { IoMdClose } from "react-icons/io";
import { User, deleteUser } from "@/services/UserServices/User";

type UserCardProps = { user: User };
export default function UserCard({ user }: UserCardProps) {
    const DisableUser = async (userId:string) => {
            try{
                await deleteUser(userId);
                console.log("Usuário desabilitado com sucesso")
            } catch(error){
                console.error(`Erro ao desabilitar usuário`, error)
            }
    };
  return (
    
    <div className="flex m-2 max-h-14 lg:max-h-24 2xl:max-h-32 justify-evenly rounded  bg-slate-400 text-black font-sans  text-center text-md">
      
      <CgProfile className="m-3 md:m-2 lg:my-4 md:w-10 lg:w-12 xl:w-16 2xl:w-20 md:h-10 lg:h-12 xl:h-16 2xl:h-20 text-black bg-white rounded-full  " />

      <div className="flex-col my-4 md:my-2.5 lg:my-4">
        <h1 className="text-sm lg:text-lg xl:text-xl 2xl:text-3xl font-bold  text-left">{user.name}</h1>
        <div className="flex text-sm lg:text-base xl:text-xl 2xl:text-3xl space-x-2">
          <a className="text-orange-700 ">Assessor</a>
          <Link href="" className="hover:text-blue-600">Relatório</Link>
        </div>
      </div>

      <FaWhatsapp  className=" my-3 md:my-2 lg:my-4 xl:my-5 p-1 md:w-10 lg:w-12 xl:w-14 2xl:w-20 md:h-10 lg:h-12 xl:h-14 2xl:h-20 text-white bg-green-500 rounded-full" />

      <a
        href=""
        className="flex min-w-0 max-w-10 lg:w-28 xl:w-32 2xl:w-52 lg:h-14 2xl:h-20 p-2 mx-4 my-2.5 lg:my-3.5 xl:my-5  space-x-1 bg-white hover:bg-red-600">
        <IoMdClose className="my-1 md:my-0 xl:-my-1 2xl:my-1 md:w-6 lg:w-8 xl:w-12 2xl:w-16 md:h-6 lg:h-8 xl:h-12 2xl:h-16 " />
        <button className="text-base xl:text-lg 2xl:text-3xl" onClick={() => DisableUser(`${user.id}`)}>Desabilitar</button>
      </a>
    </div>
  );
}