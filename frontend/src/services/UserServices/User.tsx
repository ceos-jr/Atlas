import { apiProvider } from '@/provider/provider';

export type User = {
  id: number;
  name: string;
  email: string;
  status: number;
};

export async function deleteUser(userId: string) {
    try {
      const path = `user/disable/${userId}`;
      const response = await apiProvider.delete<{ message: string }>(path);
  
      console.log(response);
      console.log(`Usuário com ID ${userId} foi deletado com sucesso.`);

      return { ok: true, data: response};
    } catch (error) {
      console.error(`Erro ao deletar usuário com ID ${userId}:`, error);

      return { ok: false, err: error };
    }
}
export async function listUsers() {
  try{
    const path = `listusers`;
    const response = await apiProvider.get<{message: string, users: User[]}>(path);
    console.log(`Sucesso na Listagem.`);
    return {ok: true, data: response};
    
  } catch (error){
    console.error(`Erro ao listar usuários:`, error);
  }
}