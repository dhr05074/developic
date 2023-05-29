import { useRef, useEffect } from "react";
import NavBar from "@/component/NavBar/NavBar";
import MarkDown from "@/component/Resizable/MarkDown";
import useResizable from "@/hook/Resizable.hook";
import { Problem } from "api/api";
import useProblem from "../../hook/Problem.hook";
import CodeEditor from "../CodeEditor/CodeEditor";

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
    const { problem } = useProblem();
    const runner = useRef<HTMLDivElement>(null);
    const body = useRef<HTMLSelectElement>(null);
    const { getRef, handleMouseMove, handleMouseDown, handleMouseUp, runnerWidth } = useResizable();

    useEffect(() => {
        getRef(runner, body);
    }, []);
    return (
        <div id="CodeEditor" className="App h-full w-full">
            <section
                role="presentation"
                id="ce_body"
                ref={body}
                className="flex h-full w-full flex-row "
                onMouseUp={handleMouseUp}
                onMouseMove={handleMouseMove}
            >
                {/* 고정 */}
                <article className=" flex w-full  flex-col bg-Navy-900 p-6 text-left ">
                    <h3 className="text-xl">{problem?.title ? problem?.title : "Add Two Numbers"}</h3>
                    <MarkDown
                        markdown={
                            problem?.background
                                ? problem?.background
                                : `You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.`
                        }
                    />
                </article>
                <article className="flex w-full flex-row">
                    <div id="code" className=" flex h-full w-1/2 flex-auto bg-Navy-900 "></div>

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
                            className="motion_basic h-full w-4 cursor-col-resize bg-Navy-800 hover:bg-Navy-300"
                        />
                        <CodeEditor code={problem?.code} />
                    </div>
                </article>
                {/* 움직임 */}
            </section>
        </div>
    );
}
