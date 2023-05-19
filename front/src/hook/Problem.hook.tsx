import { useEffect, useState } from "react";
import { useRecoilState } from "recoil";
import { useSearchParams } from "react-router-dom";
import { GetProblem200Response } from "api/api";
import { setProblemState } from "../recoil/problem";
import { generateProblem } from "@/api/problem.api";

// recoil로 변경
let problemId = "";
let problemData = null;

const useProblem = () => {
    // const [getProblemState, setId, setData] = useRecoilState(setProblemState);
    const [searchParams] = useSearchParams();
    const [hookProblem, setProblem] = useState<GetProblem200Response>();
    const difficulty = Number(searchParams.get("difficulty"));
    const language = searchParams.get("language");

    const createProblem = async () => {
        console.log("createProblem");
        const problem = await generateProblem().create(language, difficulty);
        problemId = problem.request_id;
    };
    const getProblemData = async () => {
        const problemInterval = setInterval(async () => {
            console.log("problemInterval", problemId);
            if (problemId) {
                try {
                    const newProblem = await generateProblem().get(problemId);
                    console.log("🚀 ~ file: Problem.hook.tsx:27 ~ problemInterval ~ newProblem:", newProblem);
                    if (newProblem) {
                        clearInterval(problemInterval);
                        problemData = newProblem;
                    }
                } catch (err) {
                    throw err;
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
        setProblem(problemData);
    }, []);

    return {
        // getProblemState,
        hookProblem,
        createProblem,
        getProblemData,
    };
};

export default useProblem;
