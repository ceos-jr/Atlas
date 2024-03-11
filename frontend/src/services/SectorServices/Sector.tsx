import { apiProvider } from '@/provider/provider';

export async function createSector(name: string, description: string, admId: string) {
  try{
    const path = "sector/register";
    const data = {
      name: name,
      description: description,
      admId: admId,
    }

    const response = await apiProvider.post<{message: string}, typeof data>(path, data);

    console.log(response);
    console.log("Setor criado com sucesso!");

  } catch (error) {
    console.error("Erro ao criar setor", error);
  }
}