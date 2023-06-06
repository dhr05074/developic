import { useEffect } from "react";
import { useRecoilState } from "recoil";
import { useSearchParams } from "react-router-dom";
import { generateProblem } from "@/api/problem.api";
import { problemIdState, problemState } from "../recoil/problem.recoil";

// recoil로 변경
// const problemId = "";

const useProblem = () => {
    const [problemId, setId] = useRecoilState(problemIdState);
    const [problem, setProblem] = useRecoilState(problemState);
    const [searchParams] = useSearchParams();
    const difficulty = Number(searchParams.get("difficulty"));
    const language = searchParams.get("language");

    const initProblem = () => {
        generateProblem().get("", true);
        setId(null);
        setProblem(null);
    };
    const createProblem = async () => {
        if (language) {
            const p = await generateProblem().create(language, difficulty);
            setId(p.request_id);
        }
    };

    const getProblemData = async () => {
        // let re = null;

        const interval = setInterval(async () => {
            if (problemId) {
                const p_data = await generateProblem().get(problemId);
                if (p_data) {
                    setProblem(p_data);
                    clearInterval(interval);
                }
            }
        }, 3000);
    };
    // didMount 대용

    useEffect(() => {
        if (!problemId && !problem) createProblem();
        if (problemId && !problem) getProblemData();
    }, [problem, problemId]);
    useEffect(() => {
        return () => {
            console.log("problem hook unmount");
            initProblem();
        };
    }, []);

    return {
        createProblem,
        getProblemData,
        initProblem,
        problemId,
        problem,
        // getProblemState,
    };
};

export default useProblem;
