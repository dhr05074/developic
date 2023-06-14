import { useEffect, useState } from "react";
import { useRecoilState } from "recoil";
import { useSearchParams } from "react-router-dom";
import { api, CancelToken } from "@/api/defaultApi";
import { problemIdState, problemState, editorInCode } from "../recoil/problem.recoil";
import { ProgrammingLanguage } from "api/api";

// recoil로 변경
// const problemId = "";

const useProblem = () => {
    // const [languages, setLanguages] = useRecoilState(languageState);
    // const [difficultList, setDifficultList] = useRecoilState(difficultState);

    const [problemId, setProblemId] = useRecoilState(problemIdState);
    const [problem, setProblem] = useRecoilState(problemState);
    const [editorCode, setEditorCode] = useRecoilState(editorInCode);
    const [isCodeReset, setIsCodeReset] = useState(false);
    const [searchParams] = useSearchParams();
    const getDifficulty = Number(searchParams.get("difficulty"));
    const getLanguage = searchParams.get("language") as ProgrammingLanguage;

    const initEditor = () => {
        console.log("problem.hook ");
        const code = problem?.code;
        if (code) setEditorCode(atob(code));
        setIsCodeReset(!isCodeReset);
    };
    const initProblem = () => {
        console.log("initProblem");
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
        console.log("createProblem", getLanguage);
        if (getLanguage) {
            const getCreate = await api.requestProblem({
                language: getLanguage,
            });
            return getCreate.data.problem_id;
        }
    };

    const getProblemData = async (problemId: string) => {
        console.log("getProblemData", problemId);
        if (!problemId) {
            console.log("no problemId. Problem.hook.tsx 45");
            return false;
        }
        const interval = setInterval(async () => {
            console.log("problem interval");
            const p_data = await api.getProblem(problemId);
            if (p_data) {
                console.log("getProblemData End!!", p_data);
                setProblem(p_data.data);
                setEditorCode(atob(p_data.data.code));
                clearInterval(interval);
            }
        }, 3000);
    };

    return {
        createProblem,
        getProblemData,
        initProblem,
        problemId,
        problem,
        editorCode,
        isCodeReset,
        initEditor,
        // getProblemState,
    };
};

export default useProblem;
