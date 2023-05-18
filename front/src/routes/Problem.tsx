// eslint-disable-next-line import/extensions,import/no-unresolved
import { useState, useRef, Suspense, useEffect } from "react";
import { motion } from "framer-motion";
import NavBar from "@/component/NavBar/NavBar";
import MarkDown from "@/component/Resizable/MarkDown";
import Stepper from "@/component/Stepper/Stepper";
import useResizable from "../hook/Resizable.hook";
import useProblem from "@/hook/Problem.hook";

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

function Problem() {
    const { problemData, getProblemData } = useProblem();

    const [problem, setProblem] = useState<string>("");

    const runner = useRef<HTMLDivElement>(null);
    const body = useRef<HTMLSelectElement>(null);
    const { getRef, handleMouseMove, handleMouseDown, handleMouseUp, runnerWidth } = useResizable();

    useEffect(() => {
        getRef(runner, body);
        getProblemData();
        console.log("problemData : ", problemData);
    }, []);
    const style = {
        problem: {
            width: "600px",
        },
    };
    // 컴포넌트화 해야함.
    return (
        <motion.div
            className=""
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
        >
            {/* 화면 분할해서 컴포넌트화 진행해야함. */}
            <div id="CodeEditor" className="App h-screen w-screen">
                <section id="header">
                    <NavBar currentLang={currentLang} />
                </section>
                <section
                    id="body"
                    ref={body}
                    className="flex w-full"
                    onMouseUp={handleMouseUp}
                    onMouseMove={handleMouseMove}
                >
                    <article id="problem" style={{ width: style.problem.width }} className=" overflow-auto text-left">
                        <MarkDown markdown={problem} />
                    </article>
                    {/* 고정 */}
                    <article className="flex w-full flex-row">
                        <div id="code" className="flex w-1/2 flex-auto bg-Navy-900 " />
                        <div
                            id="runner"
                            style={{ width: runnerWidth }}
                            ref={runner}
                            className="flex h-full flex-none bg-Navy-900 "
                        >
                            <div
                                id="resizeBar"
                                onMouseDown={handleMouseDown}
                                className="h-full w-4 cursor-col-resize bg-Navy-800"
                            />
                        </div>
                    </article>
                    {/* 움직임 */}
                </section>
            </div>
        </motion.div>
    );
}

export default Problem;