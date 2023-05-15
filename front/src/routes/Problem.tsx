// eslint-disable-next-line import/extensions,import/no-unresolved
import { useState, useRef } from "react";
import { useSearchParams } from "react-router-dom";
import { motion } from "framer-motion";
import NavBar from "@/component/NavBar/NavBar";
import Problem from "@/component/Resizable/Problem";
import Stepper from "@/component/Loading/Stepper";
import { generateProblem } from "@/api/problem";

const languages: string[] = [
    "go",
    "javascript",
    "typescript",
    "python3",
    "c",
    "cpp",
    "c#",
    "clojure",
    "dart",
    "elixir",
    "java",
    "kotlin",
    "php",
    "r",
    "ruby",
    "scala",
    "swift",
];
const currentLang = "javascript" as LanguageType;
function CodeEditor() {
    const [searchParams] = useSearchParams();
    const difficulty = Number(searchParams.get("difficulty"));
    const language = searchParams.get("language");

    const [useProblem, setProblem] = useState<string>("");
    // resizer
    const [runnerWidth, setRunnerWidth] = useState<number>(300);
    const [isResizing, setIsResizing] = useState<boolean>(false);
    const [step, setStep] = useState<number>(6);
    const [StepperList, setStepperList] = useState<StepperListTypes>({
        difficult: {
            value: "난이도 체크",
            step: "idle",
        },
        language: {
            value: "언어 체크",
            step: "idle",
        },
        api: {
            value: "API 불러오는 중",
            step: "idle",
        },
        comp: {
            value: "문제 출제 완료",
            step: "idle",
        },
    });
    const runnerRef = useRef<HTMLDivElement>(null);
    const bodyRef = useRef<HTMLDivElement>(null);

    const style = {
        problem: {
            width: "600px",
        },
    };

    const handleMouseMove = (event: MouseEvent) => {
        if (isResizing && runnerRef.current) {
            const current = bodyRef.current as HTMLDivElement;
            const diff = current.clientWidth - event.clientX + 10;
            setRunnerWidth(diff);
        }
    };
    // add a mousedown event listener to the sidebar to start resizing
    function handleMouseDown() {
        setIsResizing(true);
        // document.addEventListener("mousemove", handleMouseMove);
    }
    // remove the mousemove event listener when the user releases the mouse button
    function handleMouseUp() {
        console.log("🚀 ~ file: App.tsx:77 ~ handleMouseUp ~ handleMouseUp:");
        setIsResizing(false);
        // document.removeEventListener("mousemove", handleMouseMove);
    }
    console.log("CodeEditor");

    const stepperStateChanger = () => {
        if (step === 6) {
            StepperList.difficult.step = "loading";
            const changeDefficult = StepperList.difficult;
            setStepperList((prevState) => {
                return { ...prevState, difficult: changeDefficult };
            });
        } else if (step === 5) {
            StepperList.difficult.step = "complete";
            StepperList.language.step = "loading";
            const changeDefficult = StepperList.difficult;
            const changeLanguage = StepperList.language;
            setStepperList((prevState) => {
                return { ...prevState, difficult: changeDefficult, language: changeLanguage };
            });
        } else if (step === 4) {
            StepperList.language.step = "complete";
            StepperList.api.step = "loading";
            const changeLanguage = StepperList.language;
            const changeApi = StepperList.api;
            setStepperList((prevState) => {
                return { ...prevState, language: changeLanguage, api: changeApi };
            });
        } else if (step === 3) {
            StepperList.api.step = "complete";
            StepperList.comp.step = "loading";
            const changeApi = StepperList.api;
            const changeComp = StepperList.comp;
            setStepperList((prevState) => {
                return { ...prevState, api: changeApi, comp: changeComp };
            });
        } else if (step === 2) {
            StepperList.comp.step = "complete";
            const changeComp = StepperList.comp;
            setStepperList((prevState) => {
                return { ...prevState, comp: changeComp };
            });
        } else if (step === 0) {
            return false;
        }
        setStep(step - 1);
        return true;
    };

    return (
        <motion.div
            className=""
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
        >
            <div id="CodeEditor" className="App h-screen w-screen">
                <Stepper step={step} list={StepperList} testFunction={stepperStateChanger} />

                {/* 스테퍼 제거해야함. */}

                <section id="header">
                    <NavBar currentLang={currentLang} />
                </section>
                <section
                    id="body"
                    ref={bodyRef}
                    className="flex w-full"
                    onMouseUp={handleMouseUp}
                    onMouseMove={handleMouseMove}
                >
                    <article id="problem" style={{ width: style.problem.width }} className=" overflow-auto text-left">
                        <Problem markdown={useProblem} />
                    </article>
                    {/* 고정 */}
                    <article className="flex w-full flex-row">
                        <div id="code" className="flex w-1/2 flex-auto bg-Navy-900 " />
                        <div
                            id="runner"
                            ref={runnerRef}
                            style={{ width: runnerWidth }}
                            className="flex h-full flex-none bg-Navy-900 "
                        >
                            <div
                                id="resizeBar"
                                onMouseDown={handleMouseDown}
                                className="h-full w-4 cursor-col-resize bg-Navy-800"
                            />
                            {runnerWidth}
                        </div>
                    </article>
                    {/* 움직임 */}
                </section>
            </div>
        </motion.div>
    );
}

export default CodeEditor;
