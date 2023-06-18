import { useEffect, useState } from "react";
import { useRecoilState } from "recoil";
import { useSearchParams, useNavigate } from "react-router-dom";
import { api, CancelToken } from "@/api/defaultApi";
import { problemIdState, problemState, editorInCode } from "../recoil/problem.recoil";
import { ProgrammingLanguage, SubmitSolutionRequest } from "api/api";
import { profileState } from "@/recoil/profile.recoil";

// recoil로 변경
// const problemId = "";

const useProblem = () => {
    // const [languages, setLanguages] = useRecoilState(languageState);
    // const [difficultList, setDifficultList] = useRecoilState(difficultState);
    const navigate = useNavigate();
    const [profile, setProfile] = useRecoilState(profileState);
    const [problemId, setProblemId] = useRecoilState(problemIdState);
    const [problem, setProblem] = useRecoilState(problemState);
    // 에디터 내부 코드
    const [editorCode, setEditorCode] = useRecoilState(editorInCode);
    const [isCodeReset, setIsCodeReset] = useState(false);
    const [searchParams] = useSearchParams();
    const getDifficulty = Number(searchParams.get("difficulty"));
    const getLanguage = searchParams.get("language") as ProgrammingLanguage;

    const initEditor = () => {
        console.log("initEditor ");
        const code = problem?.code;
        if (code) setEditorCode(atob(code));
        setIsCodeReset(!isCodeReset);
    };
    const initProblem = () => {
        console.log("initProblem");
        if (problemId) {
            api.getProblem(problemId, {
                headers: profile.headers,
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
            const getCreate = await api.requestProblem(
                {
                    language: getLanguage,
                },
                { headers: profile.headers },
            );
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
            const p_data = await api.getProblem(problemId, { headers: profile.headers });
            if (p_data) {
                console.log("getProblemData End!!", p_data);
                setProblem(p_data.data);
                setEditorCode(atob(p_data.data.code));
                clearInterval(interval);
            }
        }, 3000);
    };

    const onClickSubmit = () => {
        console.log(problem?.id);
        // problemId = null이다 조치해야함.
        const submit = {
            problem_id: problemId,
            code: editorCode,
        } as SubmitSolutionRequest;

        api.submitSolution(submit, { headers: profile.headers })
            .then(() => {
                navigate("/result");
            })
            .catch((error) => {
                console.error("onClickSubmit error", error);
            });
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
        onClickSubmit,
        // getProblemState,
    };
};

export default useProblem;
