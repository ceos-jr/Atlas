import axios, { AxiosInstance, AxiosRequestConfig } from "axios";

class ApiProvider {
  private instance: AxiosInstance;

  constructor(baseUrl: string) {
    this.instance = axios.create({
      baseURL: baseUrl,
    });
  }

  async get<T>(path: string, options?: AxiosRequestConfig): Promise<T> {
    try {
      const response = await this.instance.get<T>(path, options);
      return response.data;
    } catch (error) {
      throw error;
    }
  }

  async post<T, D>(
    path: string,
    data?: D,
    options?: AxiosRequestConfig,
  ): Promise<T> {
    try {
      const response = await this.instance.post<T>(path, data, options);
      return response.data;
    } catch (error) {
      throw error;
    }
  }

  async put<T, D>(
    path: string,
    data?: D,
    options?: AxiosRequestConfig,
  ): Promise<T> {
    try {
      const response = await this.instance.put<T>(path, data, options);
      return response.data;
    } catch (error) {
      throw error;
    }
  }

  async delete<T>(path: string, options?: AxiosRequestConfig): Promise<T> {
    try {
      const response = await this.instance.delete<T>(path, options);
      return response.data;
    } catch (error) {
      throw error;
    }
  }
}

export const apiProvider = new ApiProvider('http://127.0.0.1:8000/');
