import { apiProvider } from '@/provider/provider';

export type User = {
  id: number;
  name: string;
  email: string;
  status: number;
};
export type updateUserBody = {
  id: number;
  name?: string | null;
  email?: string | null;
  password?: string | null;
}
export type updateUserRoleBody = {
  Id: number;
  UserId: number;
  RoleId: number;
}
export type updateProjectBody={
  Id: number;
  
}
export async function deleteUser(userId: string) {
  try {
    const path = `user/disable/${userId}`;
    const response = await apiProvider.delete<{ message: string }>(path);

    console.log(response);
    console.log(`Usuário com ID ${userId} foi deletado com sucesso.`);

    return { ok: true, data: response };
  } catch (error) {
    console.error(`Erro ao deletar usuário com ID ${userId}:`, error);

    return { ok: false, err: error };
  }
}
export async function listUsers() {
  try {
    const path = `listusers`;
    const response = await apiProvider.get<{ message: string, users: User[] }>(path);
    console.log(`Sucesso na Listagem.`);
    return { ok: true, data: response };

  } catch (error) {
    console.error(`Erro ao listar usuários:`, error);
  }
}
export async function updateUser(userId: string, updateUserData: updateUserBody) {
  console.log(updateUserData);
  if (updateUserData.password) {
    console.log("não é nulo")
    console.log(updateUserData.password)
  }
  try {
    const path = `user/update/${userId}`;
    const response = await apiProvider.put<{ message: string }, updateUserBody>(path, updateUserData);

    console.log(response);
    console.log(`Usuário com ID ${userId} foi atualizado com sucesso.`);
    return { ok: true, data: response };
  } catch (error) {
    console.error(`Erro ao atualizar usuário com ID ${userId}: `, error);
    return { ok: false, err: error }
  }
}

export async function updateUserDataAndRelated(userId: string, updateUserData: updateUserBody) {
  // Lógica para atualizar informações do usuário
  await updateUser(userId, updateUserData);

  // Lógica para atualizar função do usuário
  //await updateUserRole(userId, updateUserRoleData);

  // Lógica para atualizar projeto do usuário
  //await updateUserProject(userId, updateUserProjectData);
}