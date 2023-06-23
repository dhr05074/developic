import axios from "axios";
import { DefaultApi } from "../../api/api";

const instance = axios.create();
// const baseURL = "http://localhost:3000";
const baseURL = "https://api.developic.kr";
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

const CancelToken = axios.CancelToken;

const api = new DefaultApi(undefined, baseURL, instance);

export { api, CancelToken };
