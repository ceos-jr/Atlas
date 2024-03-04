"use client";


import React, { useState, useEffect } from "react";
import { updateUserBody, updateUser } from "@/services/UserServices/User";

interface UserProps {
    user: updateUserBody;
}

const UpdateUser: React.FC<UserProps> = ({ user }) => {
    const [userData, setUserData] = useState(user);

    const handleUpdateUser = async () => {
        const updatedData: updateUserBody = {
            id: userData.id,
            email: userData.email !== user.email ? userData.email : null,
            name: userData.name !== user.name ? userData.name : null,
            status: userData.status !== user.status ? userData.status : null,
        };

        const result = await updateUser(userData.id.toString(), updatedData);
    };

    const handleInputChange = (field: keyof updateUserBody, value: string) => {
        setUserData((prevData) => {
            const newValue = value !== "" && value !== prevData[field] ? value : null;
    
            return {
                ...prevData,
                [field]: newValue !== undefined ? newValue: "",
            };
        });
    };

    useEffect(() => {
        setUserData(user);
    }, [user]);

    return (
        <div className="h-screen bg-white text-black overflow-scroll">
            <div>
                <label>
                    UserID:
                    <input
                        type="text"
                        value={userData.id}
                        onChange={(e) => handleInputChange("id", e.target.value)}
                    />
                </label>
                <br />
                <label>
                    Name:
                    <input
                        type="text"
                        value={userData.name??undefined }
                        onChange={(e) => handleInputChange("name", e.target.value)}
                    />
                </label>
                <br />
                <label>
                    Email:
                    <input
                        type="text"
                        value={userData.email??undefined}
                        onChange={(e) => handleInputChange("email", e.target.value)}
                    />
                </label>
                <br />
                <label>
                    Status:
                    <input
                        type="text"
                        value={userData.status??undefined}
                        onChange={(e) => handleInputChange("status", e.target.value)}
                    />
                </label>
                <br />
                <a href=""><button onClick={handleUpdateUser}>Atualizar Usuário</button></a>
            </div>
        </div>
    );
};

export default UpdateUser;
