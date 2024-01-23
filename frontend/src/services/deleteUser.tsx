import { deleteuser } from '@/provider/provider'

export async function deleteUser(userId:string) {
    const data = await deleteuser(userId);
    if (data) return data
    return new Error("Erro ao deletar usuario")
}

