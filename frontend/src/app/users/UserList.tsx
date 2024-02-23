"use client";

import { useState, useEffect, ChangeEvent } from "react";
import { User, listUsers} from "@/services/UserServices/User";
import UserCard from "./components/UserCard";


type ListingProps = { state: string };
function Listing({ state }: ListingProps) {
  const [users, setUsers] = useState<User[]>([]);
  
    useEffect(() => {
        async function fetchUsers() {
          try {
            const result = await listUsers();
            if (result && result.ok) {
                const fetchedUsers: User[] = result.data.users || [];
                setUsers(fetchedUsers);
            } else {
              console.error('Erro desconhecido ao obter dados dos usuários.', result);
            }
          } catch (error) {
            console.error('Erro desconhecido ao obter dados dos usuários:', error);
          }
        }
    
        fetchUsers();
      }, []);
  const show = [];

  for (let item of users) {
    if (item.name.startsWith(state) && item.status !== 1) {
      show.push(<UserCard key={item.id} user={item}/>);
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
