import React, { useEffect, useRef } from "react";

import { useNavigate } from "react-router-dom";
import { motion } from "framer-motion";
import useSelectComponent from "@/hook/SelectComponent.hook";
import useProblem from "@/hook/Problem.hook";
import SelectComponent from "@/component/Select/Select";
import Button from "@/component/Button/Function.Button";
import useResizable from "@/hook/resizable.hook";
import CodeEditor from "@/component/CodeEditor/CodeEditor";
import MarkDown from "@/component/Resizable/MarkDown";

// const ProblemNav = React.lazy(() => import("../component/Problem/Problem.Nav"));
// const ProblemBody = React.lazy(() => import("../component/Problem/Problem.Body"));

function Problem() {
    const { languages, difficultList, setLang, setDifficulty, selectOptoin } = useSelectComponent();
    const { initEditor, onClickSubmit, problem, initProblem, setLoading } = useProblem();
    const runner = useRef<HTMLDivElement>(null);
    const body = useRef<HTMLSelectElement>(null);
    const { getRef, handleMouseMove, handleMouseUp, runnerWidth } = useResizable();
    const navigate = useNavigate();

    useEffect(() => {
        if (!problem) {
            console.log("code", problem);

            navigate("/");
        }
        getRef(runner, body);
        return () => {
            initProblem();
            setLoading(false);
        };
    }, []);
    return (
        <motion.div
            className="h-full w-full"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
        >
            <section
                id="problem_nav"
                className="flex w-full flex-row items-center justify-between bg-Navy-800 px-6 py-2"
            >
                <article id="problem_nav_options" className=" relative flex h-10 w-52 flex-row">
                    <SelectComponent
                        value={{ menu: languages, callback: setLang }}
                        disabled={selectOptoin.currentLang}
                        location="left-0"
                        size="small"
                    />
                    <SelectComponent
                        value={{ menu: difficultList, callback: setDifficulty }}
                        disabled={selectOptoin.defaultDifficulty}
                        location="right-0"
                        size="small"
                    />
                </article>
                <article id="problem_nav_buttons" className="flex flex-row gap-3">
                    <Button name="Reset" type="line" size="small" func={initEditor} />
                    <Button name="Submit" type="full" size="small" func={onClickSubmit} />
                </article>
            </section>
            <section id="CodeEditor" className="App h-full w-full">
                <section
                    role="presentation"
                    id="ce_body"
                    ref={body}
                    className="flex h-full w-full flex-row "
                    onMouseUp={handleMouseUp}
                    onMouseMove={handleMouseMove}
                >
                    {/* 고정 */}
                    <article className=" flex w-full flex-col gap-10 bg-Navy-800 p-6 text-left text-white ">
                        <h3 className="pretendard_bold_24">{problem?.title ? problem?.title : "no title"}</h3>
                        <div id="problem_description" className="flex flex-col gap-6">
                            <h4 className="pretendard_bold_20">문제 설명</h4>
                            <MarkDown markdown={problem?.description ? problem.description : "no description"} />
                        </div>
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
            </section>
        </motion.div>
    );
}

export default Problem;
