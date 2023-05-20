// eslint-disable-next-line import/extensions,import/no-unresolved
import React, { useState } from "react";
import { motion } from "framer-motion";
import { useEffect } from "react";
import { useSearchParams } from "react-router-dom";
import Stepper from "@/component/Stepper/Stepper";
import useProblem from "../hook/Problem.hook";
import fetchProblem from "@/api/Suspencer";

const ProblemComponent = React.lazy(() => import("../component/Problem/Problem"));

function Problem() {
    return (
        <motion.div
            className=""
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
        >
            <React.Suspense fallback={<Stepper />}>
                <ProfileDetails />
            </React.Suspense>
        </motion.div>
    );
}
const resource = fetchProblem().create();
// const resource2 = fetchProblem().problem();

function ProfileDetails() {
    const [id, setId] = useState("asdsad");
    const [searchParams] = useSearchParams();
    const difficulty = Number(searchParams.get("difficulty"));
    const language = searchParams.get("language");
    // const test = () => {
    //     console.log("???", id);

    //     setTimeout(() => {
    //         if (id) {
    //             setId(resource.read());
    //         } else {
    //             test();
    //         }

    //         // const getProb = fetchProblem().problem(id.request_id);
    //         // const get = getProb.read();
    //         // if (get) {
    //         //     setAdvice(get);
    //         // } else {
    //         //     test();
    //         // }
    //     }, 1000);
    // };
    // useEffect(() => {
    //     test();
    // });
    resource.read();

    // setAdvice(problem);

    // useEffect(() => {});
    // setAdvice(getProblemData());
    return <h1>{id.request_id}</h1>;
    // <ProblemComponent />
}
export default Problem;
