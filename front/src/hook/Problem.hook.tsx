import { useEffect, useState } from "react";
import { useRecoilState } from "recoil";
import { useSearchParams } from "react-router-dom";
import { GetProblem200Response, Problem } from "api/api";
import { generateProblem } from "@/api/problem.api";
import { setProblemId } from "../recoil/problem";

// recoil로 변경
// const problemId = "";
const problemData: Problem | null = null;

const useProblem = () => {
    const [getId, setId] = useRecoilState(setProblemId);
    const [searchParams] = useSearchParams();
    const difficulty = Number(searchParams.get("difficulty"));
    const language = searchParams.get("language");
    const [currentProblem, setProblem] = useState<Problem>();

    const createProblem = async () => {
        const problem = await generateProblem().create(language, difficulty);
        console.log("🚀 ~ file: Problem.hook.tsx:21 ~ createProblem ~ problem:", problem);
        setId(problem.request_id);
        console.log("getId", getId);
    };

    const getProblemData = async () => {
        // let re = null;

        // const interval = setInterval(async () => {
        // const resource = await fetchSuspeceData(suspenceFunction());
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
    }, []);

    return {
        createProblem,
        // getProblemState,
    };
};

export default useProblem;
