import { BrowserRouter, Navigator, Route, Routes } from "react-router-dom";
import ErrorPage from "@/routes/Error";

function AppRouter() {
    return (
        <BrowserRouter>
            {/* <Navigator/> */}
            <Route path="/" element={<CodeEditor />} errorElement={<ErrorPage />} />
        </BrowserRouter>
    );
}

export default AppRouter;
