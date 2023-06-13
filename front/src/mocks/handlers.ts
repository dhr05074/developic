import { rest } from "msw";
import { RequestProblem202Response, Problem } from "api/api";
// src/mocks/handlers.js

// problem
const apiUrl = "http://15.165.21.53:3000";

const createProblem: RequestProblem202Response = {
    problem_id: "problem_id",
};

const getProblem = {
    problem: {
        problem_id: "",
        title: "",
        background: "",
        code: "",
        estimated_time: 0,
    },
};

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
    setTimeout(() => {
        getProblem.problem = {
            problem_id: "123",
            title: "Add Two Numbers",
            background: `<p class="has-line-data" data-line-start="3" data-line-end="4">You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.</p>
            <p class="has-line-data" data-line-start="5" data-line-end="6">You may assume the two numbers do not contain any leading zero, except the number 0 itself.</p>`,
            code: btoa(code),
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
        // return res(ctx.status(404), ctx.json(null));
        return res(ctx.status(200), ctx.json(createProblem));
    }),
    rest.get(`${apiUrl}/problems/:requestId`, async (req, res, ctx) => {
        const { requestId } = req.params;
        console.log("get ID", requestId);
        await onCreateProblem();

        return res(ctx.status(200), ctx.json(getProblem));
    }),
];
