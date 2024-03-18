import { apiProvider } from '@/provider/provider';

export type RoleBody = {
    id: number;
    name: string;
    description: string;
}

export async function listRoles() {
    try {
        const path = `/role/listroles`;
        const response = await apiProvider.get<{ message: string, roles: RoleBody[] }>(path);
        console.log(`Sucesso na Listagem.`);
        return { ok: true, data: response };

    } catch (error) {
        console.error(`Erro ao listar usuários:`, error);
    }
}