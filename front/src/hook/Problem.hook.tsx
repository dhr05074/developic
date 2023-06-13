import { useEffect } from "react";
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
    const [searchParams] = useSearchParams();
    const getDifficulty = Number(searchParams.get("difficulty"));
    const getLanguage = searchParams.get("language") as ProgrammingLanguage;

    const initEditor = () => {
        setEditorCode(problem?.code);
        console.log("initEditor", editorCode);
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
            const p_data = await api.getProblem(problemId);
            if (p_data) {
                setProblem(p_data.data);
                setEditorCode(p_data.data.code);
                clearInterval(interval);
            }
        }, 3000);
    };
    // didMount 대용
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
        editorInCode,
        initEditor,
        // getProblemState,
    };
};

export default useProblem;
