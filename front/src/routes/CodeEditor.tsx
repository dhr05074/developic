// eslint-disable-next-line import/extensions,import/no-unresolved
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

const getMenuValue = (value: string) => {
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

function CodeEditor() {
    return (
        <div id="CodeEditor" className="App w-screen h-screen">
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
                <article id="problem" style={{ width: style.problem.width }} className=" text-left overflow-auto">
                    <Problem markdown={useProblem} />
                </article>
                {/* 고정 */}
                <article className="flex flex-row w-full">
                    <div id="code" className="bg-Navy-900 flex flex-auto w-1/2 " />
                    <div
                        id="runner"
                        ref={runnerRef}
                        style={{ width: runnerWidth }}
                        className="h-full flex flex-none bg-Navy-900 "
                    >
                        <div
                            id="resizeBar"
                            onMouseDown={handleMouseDown}
                            className="h-full w-4 bg-Navy-800 cursor-col-resize"
                        />
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

export default CodeEditor;
