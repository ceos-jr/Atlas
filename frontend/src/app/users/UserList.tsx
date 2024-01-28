"use client";

import { useState, ChangeEvent } from "react";

import Link from "next/link";
import { FaWhatsapp } from "react-icons/fa";
import { CgProfile } from "react-icons/cg";
import { IoMdClose } from "react-icons/io";

type UserCardProps = { user: string };
function UserCard({ user }: UserCardProps) {
  return (
    <div className="flex m-2 max-h-14 lg:max-h-24 2xl:max-h-32 justify-evenly rounded  bg-slate-400 text-black font-sans  text-center text-md">
      
      <CgProfile className="m-3 md:m-2 lg:my-4 md:w-10 lg:w-12 xl:w-16 2xl:w-20 md:h-10 lg:h-12 xl:h-16 2xl:h-20 text-black bg-white rounded-full  " />

      <div className="flex-col my-4 md:my-2.5 lg:my-4">
        <h1 className="text-sm lg:text-lg xl:text-xl 2xl:text-3xl font-bold  text-left">{user}</h1>
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
        <button className="text-base xl:text-lg 2xl:text-3xl">Desabilitar</button>
      </a>
    </div>
  );
}

type ListingProps = { state: string };
function Listing({ state }: ListingProps) {
  const list = [
    "Alan Cabral",
    "LG Bessa",
    "Naomi Takimaki",
    "Vivi Nenao",
    "Joelio",
    "Wagner",
    "Alan Cabral",
    "LG Bessa",
    "Naomi Takimaki",
    "Vivi Nenao",
    "Joelio",
    "Wagner",
  ];
  const show = [];

  for (let item of list) {
    if (item.startsWith(state)) {
      show.push(<UserCard user={item} />);
    }
  }
  return show;
}

export default function UserList() {
  const [state, setState] = useState<string>("");

  function handleChange(e: ChangeEvent<HTMLInputElement>) {
    const value = e.target.value;

    setState(value);
  }

  return (
    <div className="flex-col p-2 2xl:p-6">
      <div className="flex justify-evenly md:justify-start">
        <input
          placeholder="Pesquisar usuários"
          value={state}
          onChange={handleChange}
          className="mx-2 my-4 p-2 w-60 lg:w-80 2xl:w-[20%] 2xl:h-16 text-black text-sm lg:text-lg 2xl:text-3xl border-2 rounded-md border-black"
        ></input>
        <button className="my-4 p-2 lg:w-28 xl:w-32 2xl:w-48 2xl:h-16 text-black text-sm lg:text-lg 2xl:text-3xl border-2 rounded-md bg-blue-500 hover:bg-blue-700">
          Adicionar
        </button>
      </div>
      <div className="max-h-96 lg:max-h-[35rem] 2xl:max-h-[72rem] scroll-smooth overflow-y-scroll bg-white">
        <Listing state={state} />
      </div>
    </div>
  );
}
