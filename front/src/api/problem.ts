import axios from "axios";
import { apiErrorHandler } from "./errorhandler";
const instance = axios.create({
    baseURL: "http://localhost:8080",
    // timeout: 1000,
    // headers: {'X-Custom-Header': 'foobar'}
});
const headers = { "Content-Type": `application/json` };
type languageType = "Go" | "JavaScript" | "TypeScript";
export const postProblem = async (language: languageType, difficulty: number) => {
    try {
        const res = await instance.post("/problem", { language, difficulty }, { headers });
        console.log("🚀 ~ file: problem.ts:12 ~ res:", res);
    } catch (err) {
        apiErrorHandler(err);
    }
};
