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
import NavBar from "./component/NavBar/NavBar";
// import Problem from "./component/Resizable/Problem";
import Problem from "./pages/Problem";
import ErrorPage from "./pages/Error";
import { PageHome } from "./pages/Home/Usage";
import Stepper from "./pages/Stepper";
import Result from "./pages/Result";
import Profile from "./pages/Profile";
import LoadingComponent from "./component/Loading/Loading";

function App(): JSX.Element {
  return (
    <RecoilRoot>
      <div className="flex h-screen w-screen flex-col">
        <LoadingComponent />
        <section id="header" className=" h-16 w-full">
          <NavBar />
        </section>
        <section id="body" className="h-full w-full">
          <Routes>
            <Route
              path="/"
              element={<PageHome />}
              errorElement={<ErrorPage />}
            />
            <Route
              path="/stepper"
              element={<Stepper />}
              errorElement={<ErrorPage />}
            />
            <Route
              path="/problem"
              element={<Problem />}
              errorElement={<ErrorPage />}
            />
            <Route
              path="/result"
              element={<Result />}
              errorElement={<ErrorPage />}
            />
            <Route
              path="/profile"
              element={<Profile />}
              errorElement={<ErrorPage />}
            />
          </Routes>
        </section>
      </div>
    </RecoilRoot>
  );
}

export default App;
