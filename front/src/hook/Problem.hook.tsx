import { useEffect, useState } from "react";
import { useRecoilState } from "recoil";
import { useSearchParams } from "react-router-dom";
import { GetProblem200Response, Problem } from "api/api";
import { setProblemState } from "../recoil/problem";
import { generateProblem } from "@/api/problem.api";
import fetchSuspeceData from "../api/Suspencer";

// recoil로 변경
let problemId = "";
const problemData: Problem | null = null;

const useProblem = () => {
    // const [getProblemState, setId, setData] = useRecoilState(setProblemState);
    const [searchParams] = useSearchParams();
    const difficulty = Number(searchParams.get("difficulty"));
    const language = searchParams.get("language");
    const [currentProblem, setProblem] = useState<Problem>();

    const createProblem = async () => {
        console.log("createProblem");
        const problem = await generateProblem().create(language, difficulty);
        problemId = problem.request_id;
    };
    const suspenceFunction = () => {
        return new Promise((resolve) => {
            resolve(generateProblem().get(problemId));
        });
    };
    const getProblemData = async () => {
        // let re = null;

        // const interval = setInterval(async () => {
        const resource = await fetchSuspeceData(suspenceFunction());
        // resource.read();
        // console.log("resource.read()", await resource.read());

        // re = await resource.read();
        // if (re) clearInterval(interval);
        // }, 1000);

        return resource;
    };

    // didMount 대용
    useEffect(() => {
        console.log("problem.hook : useEffect");
        createProblem();
    }, []);

    return {
        // getProblemState,
        suspenceFunction,
    };
};

export default useProblem;
