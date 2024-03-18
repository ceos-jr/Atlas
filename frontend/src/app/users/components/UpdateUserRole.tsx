"use client";
import { useState, useEffect, ChangeEvent } from "react";
import { RoleBody, listRoles } from "@/services/RoleServices/Role";

export default function UpdateUserRole() {


    function Listing() {
        const [roles, setRoles] = useState<RoleBody[]>([]);
        useEffect(() => {
            async function ListRoles() {
                try {
                    const result = await listRoles();
                    if (result && result.ok) {
                        const fetchedRoles: RoleBody[] = result.data.roles || [];
                        setRoles(fetchedRoles);
                    } else {
                        console.error('Erro ao listar cargos:', result);
                    }
                } catch (error) {
                    console.error('Erro inesperado:', error);
                }
            }

            console.log(ListRoles());
        }, []);
        const show = [];
        console.log(roles);
        for (let item of roles) {
            console.log(item.name);
              show.push(
                <option key={item.id} value={item.id}>
                    {item.name}
                </option>
              );
          }
          console.log(show);
          return show;
    }

    return (
        <div>
            <select 
            id="role"
            >
                <option value="">Selecione um Cargo</option>
                <Listing/>
            </select>
        </div>
    );
}
