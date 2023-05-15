import React from "react";
import { RecoilRoot } from "recoil";
import "./styles/App.css";
import "./styles/index.css";
import "./styles/component.css";
// react-simple-code-editr
// import { useCodeMirror } from "@uiw/react-codemirror";
// import { javascript } from "@codemirror/lang-javascript";
// import CodeEditor from './component/CodeEditor/CodeEditor';
import { Route, Routes } from "react-router-dom";
// import NavBar from "./component/NavBar/NavBar";
// import Problem from "./component/Resizable/Problem";
import { AnimatePresence } from "framer-motion";
import Problem from "./routes/Problem";
import ErrorPage from "./routes/Error";
import HomePage from "./routes/Home";
import Select from "./routes/Select";

function App(): JSX.Element {
    return (
        <RecoilRoot>
            <AnimatePresence>
                <Routes>
                    <Route path="/" element={<HomePage />} errorElement={<ErrorPage />} />
                    <Route path="/select" element={<Select />} errorElement={<ErrorPage />} />
                    <Route path="/codeEditor" element={<Problem />} errorElement={<ErrorPage />} />
                </Routes>
            </AnimatePresence>
        </RecoilRoot>
    );
}

export default App;
