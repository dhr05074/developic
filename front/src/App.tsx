import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
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
import ResizableComponent from "./component/Resizerble/Resizerble";

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
    // const [textareaValue, setText] = useState("입력하세요");
    // const [] = useState("");
    // const onTextChange = (e) => {
    //     setText(e.target.value);
    // };
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
            width: "26rem",
        },
    };

    return (
        <div className="App w-screen h-screen">
            <section id="header">
                <NavBar />
            </section>
            <section id="body" className="flex flex-row w-full h-full">
                <article id="problem" style={{ width: style.problem.width }} className=" bg-Navy-800"></article>
                {/* 고정 */}
                <article id="code" className="w-[60%] bg-Navy-900"></article>
                {/* 움직임 */}
                <article id="runner" className="w-[40%] h-full">
                    <ResizableComponent />
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
