import { javascript } from '@codemirror/lang-javascript';
import { rest } from "msw";
import { RequestProblem202Response } from "api/api";
import { SubmitSolutionRequest } from "api/api";
import { Record } from "api/api";
// src/mocks/handlers.js

// problem
// const apiUrl = "http://localhost:3000";
const apiUrl = "https://api.developic.kr";

const createProblem: RequestProblem202Response = {
    problem_id: "problem_id_msw",
};

let problem = {
    id: "",
    title: "",
    description: "",
    code: "",
};
let record = {
    // record로 받는값.
    id: "record_id",
    problem_id: "",
    problem_title: "",
    efficiency: 0,
    readability: 0,
    robustness: 0,
    code: "",
};
const records: Record[] = [];

const js = `
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
const go =`
// Prime Sieve in Go.
// Taken from the Go specification.
// Copyright © The Go Authors.

package main

import "fmt"

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i  // Send 'i' to channel 'ch'
	}
}`
const cpp = `#include <iostream>
#include "mystuff/util.h"

namespace {
enum Enum {
  VAL1, VAL2, VAL3
};

char32_t unicode_string = U"\U0010FFFF";
string raw_string = R"delim(anything
you
want)delim";

int Helper(const MyType& param) {
  return 0;
}
} // namespace

class ForwardDec;

template <class T, class V>
class Class : public BaseClass {
  const MyType<T, V> member_;

 public:
  const MyType<T, V>& Method() const {
    return member_;
  }

  void Method2(MyType<T, V>* value);
}

template <class T, class V>
void Class::Method2(MyType<T, V>* value) {
  std::out << 1 >> method();
  value->Method3(member_);
  member_ = value;
}
`
const onCreateProblem = (language:string) => {
    console.log("문제 제출 시작",language);
    let code = "";

    setTimeout(() => {
        problem = {
            id: createProblem.problem_id,
            title: "Add Two Numbers",
            description: `## 2023

            ### 4/1
            
            -   개발환경 세팅 완료.
                -   mui5, iconify, lint ...
                -   v0.0.0:1
            
            ### 4/2`,
            code:""
        };
        if(language === "Cpp"){ problem.code =  btoa(cpp)}else
        if(language === "Javascript"){problem.code = btoa(js)} else
        if(language === "Go") problem.code = btoa(go)
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
    }, 10000);
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
            const body:any = req.body
        onCreateProblem(body?.language); // 문제 생성 시작
        return res(ctx.status(200), ctx.json(createProblem));
    }),
    rest.get(`${apiUrl}/problems/:requestId`, (req, res, ctx) => {
        const { requestId } = req.params;
        console.log("msw : get ID", problem);
        if (problem.id) {
            return res(ctx.status(200), ctx.json(problem));
        } else {
            return res(ctx.status(409), ctx.json("아직 문제 생성 안됌."));
        }
    }),
    // submit
    rest.post(`${apiUrl}/submit`, (req, res, ctx) => {
        const { problem_id, code } = req.body as SubmitSolutionRequest;
        onCreateRecord(); // 채점 시작
        return res(
            ctx.status(200),
            ctx.json({
                record_id: record.id,
            }),
        );
    }),
    // record
    rest.get(`${apiUrl}/records/:record_id`, (req, res, ctx) => {
        const { record_id } = req.params;
        if (record.code) {
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
