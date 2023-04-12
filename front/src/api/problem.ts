import axios from "axios";
import { apiErrorHandler } from "./errorhandler";

const instance = axios.create({
    // baseURL: "http://localhost:8080",
    baseURL: "http://15.165.21.53:3000",
    // timeout: 1000,
    // headers: {'X-Custom-Header': 'foobar'}
});
const headers = { "Content-Type": `application/json` };

type postProblemReturn = {
    functions: string[];
    requirements: string[];
    statement: string;
    test_cases: string[];
};
/**
 * @param {Languages} language
 * @param {number}difficulty
 * @returns {postProblemReturn}
 */
export const postProblem = async (language: Language, difficulty: number) => {
    try {
        const res = await instance.post("/problems", { language, difficulty }, { headers });
        const data = res.data as postProblemReturn;
        console.log("🚀 ~ file: problem.ts:12 ~ res:", data);

        return data;
    } catch (err) {
        apiErrorHandler(err);
    }
};

type getProblemReturnType = {
    id: string;
    language: Languages;
    content: string;
};

export const getProblem = async (problem_id: string) => {
    try {
        const res = await instance.get(`/problems/${problem_id}`, { headers });
        const data = res.data as postProblemReturn;
        console.log(data);
        // content : base64
        // id
        // language
        return data;
    } catch (err) {
        apiErrorHandler(err);
    }
};
