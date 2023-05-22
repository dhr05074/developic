import { useState, useRef, useEffect } from "react";
import { GetProblem200Response } from "api/api";
import { useSearchParams } from "react-router-dom";
import NavBar from "@/component/NavBar/NavBar";
import MarkDown from "@/component/Resizable/MarkDown";
import useProblem from "@/hook/Problem.hook";
import { generateProblem } from "@/api/problem.api";
import wrapPromise from "@/api/Suspencer";
import useResizable from "@/hook/Resizable.hook";

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
// const resource = fetchProblem().create();

// const resource2 = fetchProblem().problem();
export default function ProblemComponent() {
    const [problem, setProblem] = useState<GetProblem200Response | null>(null);
    const [searchParams] = useSearchParams();
    const difficulty = Number(searchParams.get("difficulty"));
    const language = searchParams.get("language");

    const runner = useRef<HTMLDivElement>(null);
    const body = useRef<HTMLSelectElement>(null);
    const { getRef, handleMouseMove, handleMouseDown, handleMouseUp, runnerWidth } = useResizable();

    const resorce = wrapPromise(generateProblem().create(language, difficulty));

    const prom = generateProblem.create(language, difficulty);

    prom.then().catch((e: Error) => {
        throw e;
    });

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
        // setProblem(res.request_id);

        resorce.read();
    }, []);
    return (
        <div id="CodeEditor" className="App h-screen w-screen">
            <section id="header">
                <NavBar currentLang={currentLang} />
                {problem}
            </section>
            <section
                role="presentation"
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
                            role="presentation"
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
