import { useEffect, useState } from "react";
import { useRecoilState } from "recoil";
import { useSearchParams } from "react-router-dom";
import { api, CancelToken } from "@/api/defaultApi";
import { problemState, editorInCode } from "../recoil/problem.recoil";
import { ProgrammingLanguage, SubmitSolutionRequest } from "api/api";
import { profileState } from "@/recoil/profile.recoil";
import useProfile from "./Profile.hook";
import { loadingState } from "@/recoil/component.recoil";

// recoil로 변경
// const problemId = "";

const useProblem = () => {
    // const [languages, setLanguages] = useRecoilState(languageState);
    // const [difficultList, setDifficultList] = useRecoilState(difficultState);
    const [profile] = useRecoilState(profileState);
    const [problem, setProblem] = useRecoilState(problemState);
    const [isLoading, setLoading] = useRecoilState(loadingState);
    // 에디터 내부 코드
    const [editorCode, setEditorCode] = useRecoilState(editorInCode);
    const { getSingleRecord } = useProfile();
    const [isCodeReset, setIsCodeReset] = useState(false);
    const [repeatProblem, setRepeatProblem] = useState(true);
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
        if (problem?.id) {
            api.getProblem(problem.id, {
                headers: profile.headers,
                cancelToken: new CancelToken(function executor(c) {
                    console.log("get problem cancel");
                }),
            });
        }
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
        api.getProblem(problemId, { headers: profile.headers })
            .then((res) => {
                console.log("getProblemData End!!", res);
                setProblem(res.data);
                setEditorCode(atob(res.data.code));
            })
            .catch((error) => {
                if (error.response.status === 409) {
                    setTimeout(() => {
                        if (repeatProblem) getProblemData(problemId);
                    }, 3000);
                }
            });
    };

    const onClickSubmit = () => {
        console.log("onClickSubmit", problem?.id);
        // problemId = null이다 조치해야함.
        const submit = {
            problem_id: problem?.id,
            code: editorCode,
        } as SubmitSolutionRequest;

        api.submitSolution(submit, { headers: profile.headers })
            .then((res) => {
                getSingleRecord(res.data.record_id);
                setLoading(true);
            })
            .catch((error) => {
                console.error("onClickSubmit error", error);
            });
    };
    useEffect(() => {
        return () => {
            console.log("problem.hook out");
        };
    }, []);

    return {
        createProblem,
        getProblemData,
        initProblem,
        problem,
        editorCode,
        isCodeReset,
        initEditor,
        setEditorCode,
        onClickSubmit,
        setLoading,
        setRepeatProblem,
        // getProblemState,
    };
};

export default useProblem;
