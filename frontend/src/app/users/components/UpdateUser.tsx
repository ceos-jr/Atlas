"use client";

import React, { useState, useEffect } from "react";
import { updateUserBody, updateUser } from "@/services/UserServices/User";
import UpdateUserRole from "./UpdateUserRole";


interface UserProps {
    user: updateUserBody;
    closeModal: () => void;
}

const UpdateUser: React.FC<UserProps> = ({ user, closeModal }) => {
    const [userData, setUserData] = useState(user);
    const handleUpdateUser = async () => {
        const updatedData: updateUserBody = {
            id: userData.id,
            email: userData.email !== user.email ? userData.email : null,
            name: userData.name !== user.name ? userData.name : null,
            password: userData.password !== user.password ? userData.password : null,
        };

        const result = await updateUser(userData.id.toString(), updatedData);
    };

    const handleInputChange = (field: keyof updateUserBody, value: string | Number) => {
        setUserData((prevData) => {
            const newValue = value !== "" && value !== prevData[field] ? value : null;
            return {
                ...prevData,
                [field]: newValue !== undefined ? newValue : "",
            };
        });
    };

    useEffect(() => {
        setUserData(user);
    }, [user]);


    return (
        <div className="fixed z-10 inset-0 overflow-y-auto">
            <div className="flex items-end justify-center min-h-screen pt-4 px-4 pb-20 text-center sm:block sm:p-0">
                <div className="fixed inset-0 transition-opacity" aria-hidden="true">
                    <div className="absolute inset-0 bg-gray-500 opacity-75"></div>
                </div>

                <span className="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">
                    &#8203;
                </span>

                <div
                    className="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full"
                    role="dialog"
                    aria-modal="true"
                    aria-labelledby="modal-headline"
                >
                    <div className="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
                        <div className="mb-4">
                            <label htmlFor="name" className="block text-gray-700 text-sm font-bold mb-2">
                                Nome:
                            </label>
                            <input
                                type="text"
                                value={userData.name ?? undefined}
                                onChange={(e) => handleInputChange("name", e.target.value)}
                                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                            />
                        </div>
                        <div className="mb-4">
                            <label htmlFor="email" className="block text-gray-700 text-sm font-bold mb-2">
                                Email:
                            </label>
                            <input
                                type="text"
                                value={userData.email ?? undefined}
                                onChange={(e) => handleInputChange("email", e.target.value)}
                                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                            />
                        </div>
                        <div className="mb-4">
                            <label htmlFor="password" className="block text-gray-700 text-sm font-bold mb-2">
                                password:
                            </label>
                            <input
                                type="text"
                                value={userData.password ?? ""}
                                onChange={(e) => handleInputChange("password", e.target.value)}
                                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                            />
                        </div>




                        <div>
                            <label htmlFor="status" className="block text-gray-700 text-sm font-bold mb-2">
                                Status
                            </label>
                            <select
                                id="status"
                                value={userData.name?.toString() ?? ""}
                                onChange={(e) => handleInputChange("name", e.target.value)}
                                className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                            >
                                <option value="">Selecione o status</option>
                                <option value="1">Diretor de Projetos</option>
                                <option value="2">Presidente</option>
                                <option value="3">Gerente Atlas</option>
                            </select>
                        </div>
                        <div>
                            <UpdateUserRole/>
                        </div>


                        <button
                            onClick={() => {
                                handleUpdateUser();
                                closeModal();
                            }}
                            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                        >
                            Atualizar Usuário
                        </button>
                    </div>
                </div>
            </div>
        </div>
    );
};


export default UpdateUser;