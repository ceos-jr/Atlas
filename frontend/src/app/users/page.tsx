"use client";
import { useState, useEffect, ChangeEvent } from "react";


import UserList from "./UserList";

export default function User() {
    
    return (
        <div className="h-screen bg-white text-black overflow-scroll">
            Hello Atlas
            <UserList />
        </div>
    );
}
