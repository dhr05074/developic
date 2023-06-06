import axios from "axios";
import { Problem, DefaultApi } from "../../api/api";
import { apiErrorHandler } from "./errorhandler";

const instance = axios.create();
const baseURL = "http://15.165.21.53:3000";
const headers = { "Content-Type": `application/json` };

const CancelToken = axios.CancelToken;

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
// const postProblem = async (language: Languages, difficulty: number) => {
//     try {
//         const res = await instance.post("/problems", { language, difficulty }, { headers });
//         const data = res.data as postProblemReturn;
//         console.log("🚀 ~ file: problem.ts:12 ~ res:", data);

//         return data;
//     } catch (err) {
//         apiErrorHandler(err);
//     }
// };

const generateProblem = () => {
    const api = new DefaultApi(undefined, baseURL, instance);
    const create = async (language: string, difficulty: number) => {
        const getCreate = await api.createProblem({
            language,
            difficulty,
        });
        return getCreate.data;
    };

    const get = async (requestId: string, cancel?: boolean) => {
        const problem = await api.getProblem(
            requestId,
            cancel
                ? {
                      cancelToken: new CancelToken(function executor(c) {
                          // excutor 함수는 cancel 함수를 매개 변수로 받습니다.
                          console.log("get problem cancel");
                      }),
                  }
                : undefined,
        );
        // const res = await instance.get(`/problems/${requestId}`, { headers });
        const data = problem.data.problem as Problem;
        // console.log(data);
        return data;
    };
    return {
        create,
        get,
    };
};

export { generateProblem };
