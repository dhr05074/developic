import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter } from "react-router-dom";
import App from "./App";

if (process.env.NODE_ENV === "development") {
    const { worker } = await import("@/mocks/browser");
    worker.start();
}

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
    <BrowserRouter>
        <App />
    </BrowserRouter>,
);
