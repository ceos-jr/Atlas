"use client";

import React from "react";
import { useState } from "react";
import { useRouter } from "next/navigation";

const CreateTaskPopUp = () => {
    const router = useRouter();
    const goHomePage = () => {
        router.push("/");
    };
    const [isOpen, setIsOpen] = useState(false);
    const toggle = () => {
        setIsOpen(!isOpen);
      };

    return (
        <section className="h-screen w-full items-center  bg-white content-center">
            <button onClick={toggle}  className="bg-blue-100 ease-in rounded-md m-10 p-10 "> {isOpen ? (<div> Botão </div>) : (<div> Botão </div>)}</button>
            <div className={`${isOpen ? "hidden" : "bg-teal-300 rounded-md  m-5 p-5"}`}>
                <div className="">
                    <form className="grid  gap-2 font-semibold text-md ">
                        <h1>Task's Name</h1>
                        <input type="text" placeholder="Golden Cloud" id="name" name="name" className="ring ring-red-300 rounded px-2 mx-2 w-11/12"></input>
                        <h1>Choose tag</h1>
                        <input type="text" placeholder="Sonho" id="tag" name="tag" className="ring ring-red-300 rounded px-2 mx-2 w-11/12 "></input>
                        <h1>Assign to </h1>
                        <input type="text" placeholder="Alan" id="assignee" name="assignee" className="ring ring-red-300 rounded bg-white px-2 mx-2 w-11/12"></input>
                        <h1>Deadline</h1>
                        <input type="text" placeholder="2024" id="deadline" name="deadline" className="ring ring-red-300 rounded px-2 mx-2 w-11/12"></input>
                        <h1>Description</h1>
                        <input type="text" placeholder="Trabaia junto do Mestre Sariz" id="description" name="description" className="ring ring-red-300 rounded px-2 mx-2 w-11/12"></input>
                    </form>
                    </div>
            </div>
        </section>
    )
}
export default CreateTaskPopUp;