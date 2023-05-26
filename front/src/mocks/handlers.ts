import { rest } from "msw";
import { CreateProblem202Response, GetProblem200Response } from "../../api/api";
// src/mocks/handlers.js

// problem
const apiUrl = "http://15.165.21.53:3000";

const createProblem: CreateProblem202Response = {
    request_id: "problem_id",
};

const getProblem: GetProblem200Response = {
    problem: null,
};
const onCreateProblem = () => {
    setTimeout(() => {
        getProblem.problem = {
            problem_id: "123",
            title: "문제 제목",
            background: "문제 배경 설명",
            code: btoa("code"),
            estimated_time: 0,
        };
    }, 5000);
};

export default [
    // Handles a POST /login request
    //   rest.post("/login", null),
    // Handles a GET /user request
    /**
     * @prams request
     * @prams response
     */
    rest.post(`${apiUrl}/problems`, async (req, res, ctx) => {
        console.log("msw post problems : ", req.body);
        onCreateProblem();
        // return res(ctx.status(404), ctx.json(null));
        return res(ctx.status(200), ctx.json(createProblem));
    }),
    rest.get(`${apiUrl}/problems/:requestId`, async (req, res, ctx) => {
        const { requestId } = req.params;
        return res(ctx.status(200), ctx.json(getProblem));
    }),
];
