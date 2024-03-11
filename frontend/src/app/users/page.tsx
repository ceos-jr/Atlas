import React from "react";

import UserList from "./UserList";
import UpdateUser from "./components/UpdateUser";
import { updateUserBody, updateUser } from "@/services/UserServices/User";

export default function User() {
    const updatedData: updateUserBody = {
        id: 65,
        email: "joeliopinheiro@teste.com",
        name: "joelio pinheiro",
        status: 3,
    };
    return (
        <div className="h-screen bg-white text-black overflow-scroll">
            Hello Atlas
            <UserList />
            
        </div>
    );
}
