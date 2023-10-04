import React from "react";
import { RecoilRoot } from "recoil";
import "./styles/App.css";
import "./styles/index.css";
import "./styles/component.css";
import "./styles/designSystem.css";
import { Route, Routes } from "react-router-dom";
import NavBar from "./component/NavBar/NavBar";
import Problem from "./pages/Problem";
import ErrorPage from "./pages/Error";
import { PageHome } from "./pages/Home/Usage";
import Stepper from "./pages/Loading";
import Result from "./pages/Result";
import Profile from "./pages/Profile";
import LoadingComponent from "./component/Loading/Loading";
import { GlobalStyle } from "./styles/GlobalStyle";
import { UsageFunnel } from "./component/Funnel/Usage";

function App(): JSX.Element {
  return (
    <RecoilRoot>
      <GlobalStyle />
      <div className="flex h-screen w-screen flex-col">
        <LoadingComponent />
        <section id="header" className=" h-16 w-full">
          <NavBar />
        </section>
        <section id="body" className="h-full w-full">
          <UsageFunnel />
          {/* <Routes>
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
          </Routes> */}
        </section>
      </div>
    </RecoilRoot>
  );
}

export default App;
