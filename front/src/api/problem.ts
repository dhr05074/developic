import axios from "axios";
import { apiErrorHandler } from "./errorhandler";
const instance = axios.create({
    baseURL: "http://localhost:8080",
    // timeout: 1000,
    // headers: {'X-Custom-Header': 'foobar'}
});
const headers = { "Content-Type": `application/json` };
type languageType = "Go" | "JavaScript" | "TypeScript";

type postProblemReturn = {
    functions: string[];
    requirements: string[];
    statement: string;
    test_cases: string[];
};
/**
 * @param {languageType} language
 * @param {number}difficulty
 * @returns {postProblemReturn}
 */
export const postProblem = async (language: languageType, difficulty: number) => {
    try {
        const res = await instance.post("/problem", { language, difficulty }, { headers });
        const data = res.data as postProblemReturn;
        console.log("🚀 ~ file: problem.ts:12 ~ res:", data);

        return data;
    } catch (err) {
        apiErrorHandler(err);
    }
};
