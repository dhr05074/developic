import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";
import Footer from "./component/Footer/Footer";
import TextArea from "./component/Textarea/Textarea";
import Select from "./component/Select/Select";
import { postProblem } from "./api/problem";

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

    return (
        <div className="App">
            <section id="header" className="flex flex-row justify-between">
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
            {/* <Footer /> */}
        </div>
    );
}

export default App;
