import { useEffect } from "react";
import { useRecoilState } from "recoil";
import { useSearchParams } from "react-router-dom";
import { api, CancelToken } from "@/api/defaultApi";
import { problemIdState, problemState } from "../recoil/problem.recoil";
import { ProgrammingLanguage } from "api/api";

// recoil로 변경
// const problemId = "";

const useProblem = () => {
    const [problemId, setProblemId] = useRecoilState(problemIdState);
    const [problem, setProblem] = useRecoilState(problemState);
    const [searchParams] = useSearchParams();
    const difficulty = Number(searchParams.get("difficulty"));
    const language = searchParams.get("language") as ProgrammingLanguage;

    const initProblem = () => {
        if (problemId) {
            api.getProblem(problemId, {
                cancelToken: new CancelToken(function executor(c) {
                    console.log("get problem cancel");
                }),
            });
        }
        setProblemId(null);
        setProblem(null);
    };
    const createProblem = async () => {
        console.log("createProblem", language);
        if (language) {
            const getCreate = await api.requestProblem({
                language,
            });
            return getCreate.data.problem_id;
        }
    };

    const getProblemData = async (problemId: string) => {
        console.log("getProblemData", problemId);
        const interval = setInterval(async () => {
            const p_data = await api.getProblem(problemId);
            if (p_data) {
                setProblem(p_data.data);
                clearInterval(interval);
            }
        }, 3000);
    };
    // didMount 대용

    useEffect(() => {
        createProblem().then(async (r: string) => {
            getProblemData(r);
        });
    }, [problem]);

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
