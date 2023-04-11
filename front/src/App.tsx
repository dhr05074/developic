import { useState, useRef, MouseEvent } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "../../../../../../vite.svg";
import "./App.css";
import Footer from "./component/Footer/Footer";
import TextArea from "./component/Textarea/Textarea";
import Select from "./component/Select/Select";
import { postProblem } from "./api/problem";

// react-simple-code-editor
import { useCodeMirror } from "@uiw/react-codemirror";
import { javascript } from "@codemirror/lang-javascript";
import CodeEditor from "./component/CodeEditor/CodeEditor";
import NavBar from "./component/NavBar/NavBar";
import Resizable from "react-resizable-layout";

function App() {
    const languages = [
        "Go",
        "TypeScript",
        "JavaScript",
        "Python3",
        "C",
        "C++",
        "C#",
        "Clojure",
        "Dart",
        "Elixir",
        "Java",
        "Kotlin",
        "PHP",
        "R",
        "Ruby",
        "Scala",
        "Swift",
    ];
    let currentLang = "";
    const onClickAPI = (e) => {
        e.preventDefault();

        if (currentLang) {
            postProblem(currentLang, 90);
            return;
        }
        postProblem("Go", 90);
    };
    const getMenuValue = (value: string) => {
        console.log("🚀 ~ file: App.tsx:41 ~ getMenuValue ~ value:", value);
        currentLang = value;
    };
    const style = {
        problem: {
            width: "20rem",
        },
    };
    const bodyRef = useRef<HTMLDivElement>(null);
    //resizer
    const [runnerWidth, setRunnerWidth] = useState<number>(300);
    const runnerRef = useRef<HTMLDivElement>(null);
    const [isResizing, setIsResizing] = useState<boolean>(false);

    const handleMouseMove = (event: MouseEvent) => {
        if (isResizing && runnerRef.current) {
            const diff = bodyRef.current.clientWidth - event.clientX + 10;
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
        <div className="App w-screen h-screen">
            <section id="header">
                <NavBar />
            </section>
            <section
                id="body"
                ref={bodyRef}
                className="flex w-full h-full"
                onMouseUp={handleMouseUp}
                onMouseMove={handleMouseMove}
            >
                <article id="problem" style={{ width: style.problem.width }} className=" bg-Navy-800" />
                {/* 고정 */}
                <article className="flex flex-row w-full">
                    <div id="code" className="bg-Navy-900 flex flex-auto w-1/2"></div>
                    <div
                        id="runner"
                        ref={runnerRef}
                        style={{ width: runnerWidth }}
                        className="h-full flex flex-none bg-Navy-900 "
                    >
                        <div
                            id="resizeBar"
                            onMouseDown={handleMouseDown}
                            className={"h-full w-4 bg-Navy-800 cursor-col-resize"}
                        ></div>
                        {runnerWidth}
                    </div>
                </article>
                {/* 움직임 */}
            </section>
            {/* <section className="flex flex-row justify-between">
                <Select value={{ menu: languages, callback: getMenuValue }} />
                <button
                    onClick={onClickAPI}
                    type="submit"
                    className="inline-flex items-center px-5 py-2.5 text-sm font-medium text-center text-white bg-blue-700 rounded-lg focus:ring-4 focus:ring-blue-200 dark:focus:ring-blue-900 hover:bg-blue-800"
                >
                    postProblem
                </button>
            </section>
            <h4 className="mb-4 text-4xl font-extrabold leading-none tracking-tight text-gray-900 md:text-5xl lg:text-6xl dark:text-white"></h4>
            <section className="bg-white dark:bg-gray-900">
                <TextArea />
            </section>
            <section>
                <CodeEditor />
            </section> */}
            {/* <Footer /> */}
        </div>
    );
}

export default App;
