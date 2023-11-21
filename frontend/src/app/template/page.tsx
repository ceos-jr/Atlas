"use client";

//verificar o router
//revisar codigo linha por linha
import React, { useState } from "react";
import { useRouter } from "next/navigation";
import Card1 from "./card1"

export type Card1Type ={
    taskName: string
    sectorName: string
}

const cardsType1: Card1Type[] = [
    {
        taskName: "Fazer postagem da semana 5 no instagram",
        sectorName: "Marketing"
    },

    {
        taskName: "Preparar Gameficação",
        sectorName: "Pessoas"
    },

]

const CardsPage = () => {
  const router = useRouter();

  return (
    <section className="container flex flex-col gap-y-16 justify-center items-center mx-auto mt-24">
      <div>
        <h1 className="text-4xl font-bold text-gray-800">
          Nossos <span className="text-primary-500">Cards</span>
        </h1>
      </div>
      <div className="grid grid-cols-1 gap-8 p-12 w-full md:grid-cols-3">
        {cardsType1.map((cardsType1) => (
          <Card1 key={cardsType1.taskName} {...cardsType1} />
        ))}
      </div>
    </section>
  );
};

export default CardsPage;
//ServiceCard=Card1