import {
    createBrowserRouter,
} from "react-router-dom";
import "./styles/index.css";
import ErrorPage from "@/routes/Error";
import App from "@/App";

const router = createBrowserRouter([
    {
        path: "/",
        element: <App />,
        errorElement:<ErrorPage/>,
        loader: rootLoader,
        children: [

        ],
    },
]);

export default router
