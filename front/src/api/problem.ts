import axios from "axios";
import { DefaultApi, CreateProblem202Response } from "../../api/api";
import { apiErrorHandler } from "./errorhandler";

const instance = axios.create();
const baseURL = "http://15.165.21.53:3000";
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
const postProblem = async (language: Language, difficulty: number) => {
    try {
        const res = await instance.post("/problems", { language, difficulty }, { headers });
        const data = res.data as postProblemReturn;
        console.log("🚀 ~ file: problem.ts:12 ~ res:", data);

        return data;
    } catch (err) {
        apiErrorHandler(err);
    }
};

// type getProblemReturnType = {
//     id: string;
//     language: Languages;
//     content: string;
// };

// export const getProblem = async (problem_id: string) => {
//     try {
//         const res = await instance.get(`/problems/${problem_id}`, { headers });
//         const data = res.data as postProblemReturn;
//         console.log(data);
//         // content : base64
//         // id
//         // language
//         return data;
//     } catch (err) {
//         apiErrorHandler(err);
//     }
// };
const generateProblem = () => {
    const api = new DefaultApi(undefined, baseURL, instance);
    const create = async (language: string, difficulty: number): CreateProblem202Response => {
        const getCreate = await api.createProblem({
            language,
            difficulty,
        });
        return getCreate;
    };

    const get = async (requestId: string) => {
        const problem = await api.getProblem(requestId);
        return problem;
    };
    return {
        create,
        get,
    };
};
export { generateProblem, postProblem };
