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
export async function createUser(userId: string, userName: string, userEmail: string, userStatus: string) {
  try{
    const path = "register";
    const data = {
      ID: userId,
      name: userName,
      email: userEmail,
      status: userStatus,
    }

    const response = await apiProvider.post<{message: string}, typeof data>(path, data);

    console.log(response);
    console.log("Usuário cadastrado com sucesso!");

    return {ok: true, data: response};
  } catch (error) {
    console.error("Erro ao criar usuário: ", error);

    return {ok: false, err: error};
  }
}
