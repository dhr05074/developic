import { useEffect, useState } from "react";
import { useRecoilState } from "recoil";
import { useSearchParams } from "react-router-dom";
import { setProblemState } from "../recoil/problem";
import { generateProblem } from "@/api/problem";

// recoil로 변경
let problemId = "";
let problemData = "";

const useProblem = () => {
    // const [getProblemState, setId, setData] = useRecoilState(setProblemState);
    const [searchParams] = useSearchParams();
    const difficulty = Number(searchParams.get("difficulty"));
    const language = searchParams.get("language");

    const createProblem = async () => {
        console.log("createProblem");
        const problem = await generateProblem().create(language, difficulty);
        problemId = problem.data.request_id;
    };
    const getProblemData = async () => {
        const problemInterval = setInterval(async () => {
            console.log("problemInterval", problemId);
            if (problemId) {
                const newProblem = await generateProblem().get(problemId);
                console.log("🚀 ~ file: Problem.hook.tsx:27 ~ problemInterval ~ newProblem:", newProblem);
                if (newProblem) {
                    clearInterval(problemInterval);
                    problemData = newProblem;
                }
            } else {
                console.log("getProblemData : problemId 없음.");
            }
        }, 3000);
    };
    // didMount 대용
    useEffect(() => {
        console.log("problem.hook : useEffect");
        createProblem();
    }, []);

    return {
        // getProblemState,
        problemData,
        createProblem,
        getProblemData,
    };
};

export default useProblem;
