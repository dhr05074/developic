// eslint-disable-next-line import/extensions,import/no-unresolved
import { useState, useRef } from "react";
import { motion } from "framer-motion";
import NavBar from "@/component/NavBar/NavBar";
import MarkDown from "@/component/Resizable/MarkDown";
import Stepper from "@/component/Stepper/Stepper";

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
    const [useProblem, setProblem] = useState<string>("");
    // resizer
    const [runnerWidth, setRunnerWidth] = useState<number>(300);
    const [isResizing, setIsResizing] = useState<boolean>(false);

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

    return (
        <motion.div
            className=""
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
        >
            <div id="CodeEditor" className="App h-screen w-screen">
                <Stepper />

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
                        <MarkDown markdown={useProblem} />
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

export default Problem;
