"use client";

import React, { useState } from "react";
import { useRouter } from "next/navigation";
import { deleteuser } from "@/provider/provider";
const UsersPage = () => {
  const router = useRouter();


  const userIdToDelete = "43";
  const handleDeleteUser = async () => {
    try {
      const result = await deleteuser(userIdToDelete);
      if (result.err) {
        console.error("Erro ao excluir usuário:",userIdToDelete, result.err);
      } else {
        console.log("Usuário excluído com sucesso: ",userIdToDelete, result.ok);
      }
    } catch (error) {
      console.error("Erro inesperado:", userIdToDelete, error);
    }
  };


  return (
    <div className="bg-white text-black">
      <div className="flex">
        <div className="flex flex-col w-full">
          <div className="p-4">
            <h1 className="text-2xl font-bold">Users Page Content</h1>
            {/* Adicione o conteúdo da página de usuários aqui */}
          </div>
          <div>
            <h1>Next.js + Golang Fiber</h1>
            <button onClick={handleDeleteUser}>Excluir Usuário</button>
            {/* Renderize outros componentes conforme necessário */}
          </div>
        </div>
      </div>
    </div>
  )
}

export default UsersPage;

