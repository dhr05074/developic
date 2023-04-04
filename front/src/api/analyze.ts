import axios from "axios";
import { apiErrorHandler } from "./errorhandler";

const instance = axios.create({
    baseURL: "http://localhost:8080",
    // timeout: 1000,
    // headers: {'X-Custom-Header': 'foobar'}
});

const headers = { "Content-Type": `application/json` };

export const analyze = async (data: string) => {
    const utf8Str = unescape(encodeURIComponent(data));
    const base64 = btoa(utf8Str);
    try {
        const response = await instance.post("/analyze", { base64 }, { headers });
    } catch (error) {
        apiErrorHandler(error);
    }
};
