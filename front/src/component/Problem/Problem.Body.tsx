import { useRef, useEffect } from "react";
import MarkDown from "@/component/Resizable/MarkDown";
import useResizable from "@/hook/resizable.hook";
import useProblem from "@/hook/Problem.hook";
import CodeEditor from "../CodeEditor/CodeEditor";

export default function ProblemComponent() {
    const { problem, initProblem } = useProblem();
    const runner = useRef<HTMLDivElement>(null);
    const body = useRef<HTMLSelectElement>(null);
    const { getRef, handleMouseMove, handleMouseDown, handleMouseUp, runnerWidth } = useResizable();

    useEffect(() => {
        getRef(runner, body);
        console.log("code", problem);
        return () => {
            initProblem();
        };
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
                <article className=" flex w-full  flex-col bg-Navy-800 p-6 text-left text-white ">
                    <h3 className="text-xl">{problem?.title ? problem?.title : "no title"}</h3>
                    <MarkDown markdown={problem?.description ? problem.description : "no description"} />
                </article>
                {/* <div
                    role="presentation"
                    id="resizeBar"
                    onMouseDown={handleMouseDown}
                    className="motion_basic h-full w-8 cursor-col-resize bg-Navy-800 hover:bg-Navy-300"
                /> */}
                <article className="flex w-full flex-row">
                    <div id="code" className=" flex h-full w-1/2 flex-auto bg-Navy-900 "></div>
                    <div
                        id="runner"
                        style={{ width: runnerWidth }}
                        ref={runner}
                        className="flex h-full flex-none bg-Navy-900 "
                    >
                        <CodeEditor code={problem?.code} />
                    </div>
                </article>
                {/* 움직임 */}
            </section>
        </div>
    );
}
