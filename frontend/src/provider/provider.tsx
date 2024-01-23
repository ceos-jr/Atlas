import axios from 'axios'
const baseUrl = 'http://127.0.0.1:8000/';


export async function deleteuser(userId:string) {
    console.log(baseUrl);
    try {
        const response = await axios.delete(`${baseUrl}/user/disable/${userId}`)
        return response.data;
    } catch (error) {
        console.log(error);
        // throw error;
    }
}