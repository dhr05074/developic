import { useEffect, useRef, useState } from "react";
import { GetProblem200Response } from "api/api";
import NavBar from "@/component/NavBar/NavBar";
import useResizable from "@/hook/resizable.hook";

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
const style = {
    problem: {
        width: "600px",
    },
};
export default function ProblemComponent() {
    const [problem, setProblem] = useState<GetProblem200Response>();

    const runner = useRef<HTMLDivElement>(null);
    const body = useRef<HTMLSelectElement>(null);
    const { getRef, handleMouseMove, handleMouseDown, handleMouseUp, runnerWidth } = useResizable();

    // const getProblemData = () => {
    //     setTimeout(async () => {

    //         console.log("🚀 ~ file: Problem.hook.tsx:28 ~ getProblemData ~ resource:", resource);
    //         if (!resource) return getProblemData();
    //         setProblem();
    //     }, 3000);
    // };
    useEffect(() => {
        getRef(runner, body);
        // getProblemData();
    }, []);
    return (
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
                    {/* <MarkDown markdown={problem.problem.code} /> */}
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
    );
}
