"use client";

import { useState, ChangeEvent } from "react";

import Link from "next/link";
import { FaWhatsapp } from "react-icons/fa";
import { CgProfile } from "react-icons/cg";
import { IoMdClose } from "react-icons/io";


type UserCardProps = {user : string}
function UserCard({user} : UserCardProps) {
    return (
        <div className="flex my-2 mx-4 max-h-14 space-x-4 rounded  bg-slate-400 text-black font-sans  text-center text-md">
            
            <CgProfile className="text-black bg-white rounded-full  "/>

            <div className="flex-col my-1">
                <h1 className="text-sm font-bold text-left">{user}</h1>
                <div className="flex text-sm space-x-2">
                    <a>Cargo</a>
                    <Link href="">Ver Relatório</Link>
                </div>
            </div>
            
            <FaWhatsapp/>

            <div className="flex min-w-0 max-w-8 p-2 my-2 bg-white text-sm">
                <IoMdClose className="m-1"/>
                <button>Desabilitar</button>
            </div>

        </div>
    )
}

type ListingProps = { state: string }
function Listing({ state }: ListingProps) {
    const list = ["Alan Cabral", "LG Bessa", "Naomi Takimaki", "Vivi Nenao"];
    const show = [];

    for (let item of list) {
        if (item.startsWith(state)) {
            show.push(<UserCard user={item} />)
        }
    }
    return show;
}

export default function UserList() {

    const [state, setState] = useState<string>("")

    function handleChange(e: ChangeEvent<HTMLInputElement>) {
        const value = e.target.value

        setState(value)
    }

    return (
        <div className="flex-col max-h-56 scroll-smooth overflow-y-scroll">
            <input value={state} onChange={handleChange} className=""></input>
            <Listing state={state} />
        </div>
    )
}