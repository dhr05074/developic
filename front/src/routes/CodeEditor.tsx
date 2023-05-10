// eslint-disable-next-line import/extensions,import/no-unresolved
import { useState, useRef } from "react";
import { useSearchParams } from "react-router-dom";
import NavBar from "@/component/NavBar/NavBar";
import Problem from "@/component/Resizable/Problem";

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
let currentLang = "javascript" as LanguageType;
function CodeEditor() {
    const [searchParams] = useSearchParams();
    const queryList = [...searchParams]; // [['key1', 'test1'], ['key2', 'test2']]
    console.log("🚀 ~ file: CodeEditor.tsx:31 ~ CodeEditor ~ queryList:", queryList);

    const getMenuValue = (value: LanguageType) => {
        console.log("🚀 ~ file: App.tsx:41 ~ getMenuValue ~ value:", value);
        currentLang = value;
    };
    const [useProblem, setProblem] = useState<string>("");

    const style = {
        problem: {
            width: "600px",
        },
    };
    const bodyRef = useRef<HTMLDivElement>(null);
    // resizer
    const [runnerWidth, setRunnerWidth] = useState<number>(300);
    const runnerRef = useRef<HTMLDivElement>(null);
    const [isResizing, setIsResizing] = useState<boolean>(false);

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

    const getProblem = () => {
        // problem: string
        // setProblem(problem);
    };

    return (
        <div id="CodeEditor" className="App h-screen w-screen">
            <section id="header">
                <NavBar currentLang={currentLang} getProblem={getProblem} />
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
    );
}

export default CodeEditor;
