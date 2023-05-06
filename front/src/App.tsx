import React,{ useState, useRef, MouseEvent, useEffect } from "react";
// import viteLogo from "../../../../../../vite.svg";
import "./styles/App.css";

// import Footer from "./component/Footer/Footer";
import TextArea from "./component/Textarea/Textarea";
import Select from "./component/Select/Select";

// react-simple-code-editor
// import { useCodeMirror } from "@uiw/react-codemirror";
// import { javascript } from "@codemirror/lang-javascript";
// import CodeEditor from "./component/CodeEditor/CodeEditor";
import NavBar from "./component/NavBar/NavBar";
import Problem from "./component/Resizable/Problem";
import {ReactJSXElement} from "@emotion/react/types/jsx-namespace";
import AppRouter from "@/routes/AppRouter";

function App() {

    return (
       <AppRouter />
    );
}

export default App;
