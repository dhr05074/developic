import { rest } from "msw";
import { RequestProblem202Response } from "api/api";
import { SubmitSolutionRequest } from "api/api";
import { Record } from "api/api";
// src/mocks/handlers.js

// problem
const apiUrl = "http://localhost:3000";

const createProblem: RequestProblem202Response = {
    problem_id: "problem_id_msw",
};

let problem = {
    problem_id: "",
    title: "",
    description: "",
    code: "",
};
let record = {
    // record로 받는값.
    id: "",
    problem_id: "",
    problem_title: "",
    efficiency: 0,
    readability: 0,
    robustness: 0,
    code: "",
};
const records: Record[] = [];

const code = `
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {

};`;
const onCreateProblem = () => {
    console.log("문제 제출 시작");

    setTimeout(() => {
        problem = {
            problem_id: createProblem.problem_id,
            title: "Add Two Numbers",
            description: `<p class="has-line-data" data-line-start="3" data-line-end="4">You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.</p>
            <p class="has-line-data" data-line-start="5" data-line-end="6">You may assume the two numbers do not contain any leading zero, except the number 0 itself.</p>`,
            code: btoa(code),
        };
    }, 5000);
};

const onCreateRecord = async () => {
    console.log("채점 시작");
    setTimeout(() => {
        record = {
            id: "record_id",
            problem_id: createProblem.problem_id,
            problem_title: "중첩된 If문 없애기",
            efficiency: 85,
            readability: 85,
            robustness: 85,
            code: "7J6s7ZmY7J207ZiVIOuwlOuztA==",
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
    rest.post(`${apiUrl}/problems`, (req, res, ctx) => {
        console.log("msw post problems : ", req.body);

        onCreateProblem(); // 문제 생성 시작
        return res(ctx.status(200), ctx.json(createProblem));
    }),
    rest.get(`${apiUrl}/problems/:requestId`, (req, res, ctx) => {
        const { requestId } = req.params;
        console.log("msw : get ID", requestId);
        if (problem.problem_id) {
            return res(ctx.status(200), ctx.json(problem));
        } else {
            return res(ctx.status(409), ctx.json("아직 문제 생성 안됌."));
        }
    }),
    // submit
    rest.post(`${apiUrl}/submit`, (req, res, ctx) => {
        const { problem_id, code } = req.body as SubmitSolutionRequest;
        onCreateRecord(); // 채점 시작
        console.log("problem_id : ", problem_id, "code : ", code);
        return res(
            ctx.status(200),
            ctx.json({
                record_id: record.id,
            }),
        );
    }),
    // record
    rest.get(`${apiUrl}/records/:recordId`, (req, res, ctx) => {
        const { recordId } = req.params;
        if (record.id) {
            return res(ctx.status(200), ctx.json(record));
        } else {
            return res(ctx.status(409), ctx.json("아직 채점 안됌."));
        }
    }),
    rest.get(`${apiUrl}/records`, (req, res, ctx) => {
        for (let i = 0; i < 10; i++) {
            const setRecord: Record = {
                id: `record_id_${i}`,
                problem_id: `problem_${i}`,
                problem_title: `problem_title_${i}`,
                efficiency: i,
                readability: i,
                robustness: i,
                code: `code_${i}`,
            };
            records.push(setRecord);
        }
        //대량 조회
        return res(ctx.status(200), ctx.json({ records }));
    }),
    // Me
    rest.get(`${apiUrl}/me`, (req, res, ctx) => {
        // 나를 호출
        const authHeader = req.headers.get("Authorization");
        console.log("🚀 ~ file: handlers.ts:142 ~ rest.get ~ authHeader:", authHeader);
        // console.log(atob(authHeader));
        return res(
            ctx.status(200),
            ctx.json({
                nickname: authHeader,
                elo_score: 1000,
            }),
        );
    }),
];
