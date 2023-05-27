import React from "react";
import { RecoilRoot } from "recoil";
import "./styles/App.css";
import "./styles/index.css";
import "./styles/component.css";
import "./styles/designSystem.css";
// react-simple-code-editr
// import { useCodeMirror } from "@uiw/react-codemirror";
// import { javascript } from "@codemirror/lang-javascript";
// import CodeEditor from './component/CodeEditor/CodeEditor';
import { Route, Routes } from "react-router-dom";
import { AnimatePresence } from "framer-motion";
import NavBar from "./component/NavBar/NavBar";
// import Problem from "./component/Resizable/Problem";
import Problem from "./routes/Problem";
import ErrorPage from "./routes/Error";
import HomePage from "./routes/Home";
import Stepper from "./routes/Stepper";

function App(): JSX.Element {
    return (
        <RecoilRoot>
            <AnimatePresence>
                <div className=" h-screen w-screen">
                    <section id="header" className="absolute">
                        <NavBar />
                    </section>
                    <Routes>
                        <Route path="/" element={<HomePage />} errorElement={<ErrorPage />} />
                        <Route path="/stepper" element={<Stepper />} errorElement={<ErrorPage />} />
                        <Route path="/problem" element={<Problem />} errorElement={<ErrorPage />} />
                    </Routes>
                </div>
            </AnimatePresence>
        </RecoilRoot>
    );
}

export default App;
