import { useEffect } from "react";
import { useRecoilState } from "recoil";
import { useSearchParams } from "react-router-dom";
import { setProblemState } from "../recoil/problem";
import { generateProblem } from "@/api/problem";

const useProblem = () => {
    const [getProblemState, setId, setData] = useRecoilState(setProblemState);
    const [searchParams] = useSearchParams();
    const difficulty = Number(searchParams.get("difficulty"));
    const language = searchParams.get("language");

    const createProblem = async () => {
        console.log("createProblem");
        const problem = await generateProblem().create(language, difficulty);
        // setId(problem.request_id);
    };
    const getProblemData = async (requestId: string) => {
        const problemInterval = setInterval(async () => {
            console.log("problemInterval");
            const newProblem = await generateProblem().get(requestId);
            if (newProblem) {
                clearInterval(problemInterval);
                // setData(newProblem);
            }
        }, 3000);
    };
    // didMount 대용
    useEffect(() => {
        console.log("problem.hook : useEffect");
        createProblem();
    }, []);

    return {
        getProblemState,
        setId,
        setData,
        createProblem,
        getProblemData,
    };
};

export default useProblem;
