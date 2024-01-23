import { apiProvider } from '@/provider/provider';

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
