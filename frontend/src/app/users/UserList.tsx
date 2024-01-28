"use client";

import { useState, ChangeEvent } from "react";

import Link from "next/link";
import { FaWhatsapp } from "react-icons/fa";
import { CgProfile } from "react-icons/cg";
import { IoMdClose } from "react-icons/io";

type UserCardProps = { user: string };
function UserCard({ user }: UserCardProps) {
  return (
    <div className="flex m-2 max-h-14 justify-evenly rounded  bg-slate-400 text-black font-sans  text-center text-md">
      <CgProfile className="m-3 text-black bg-white rounded-full  " />

      <div className="flex-col my-2">
        <h1 className="text-sm font-bold text-left">{user}</h1>
        <div className="flex text-sm space-x-2">
          <a className="text-orange-700 ">Assessor</a>
          <Link href="">Relatório</Link>
        </div>
      </div>

      <FaWhatsapp className=" my-3 p-1 text-white bg-green-500 rounded-full" />

      <a
        href=""
        className="flex min-w-0 max-w-10 p-2 mx-4 my-2.5 space-x-1  bg-white text-sm"
      >
        <IoMdClose className="my-1" />
        <button>Desabilitar</button>
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
    <div className="flex-col">
      <div className="flex justify-evenly">
        <input
          placeholder="Pesquisar usuários"
          value={state}
          onChange={handleChange}
          className="mx-2 my-4 p-2 w-64  text-black text-sm border-2 rounded-md border-black"
        ></input>
        <button className="my-4 p-2 text-black text-sm border-2 rounded-md bg-blue-500">
          Adicionar
        </button>
      </div>
      <div className=" max-h-96 scroll-smooth overflow-y-scroll">
        <Listing state={state} />
      </div>
    </div>
  );
}
