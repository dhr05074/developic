import React, { ReactNode, useState, useEffect, Suspense } from "react";

import { motion } from "framer-motion";
import { useSearchParams } from "react-router-dom";
import Stepper from "@/component/Stepper/Stepper";
import wrapPromise from "@/api/Suspencer";
import { generateProblem } from "@/api/problem.api";

const ProblemComponent = React.lazy(() => import("../component/Problem/Problem.Component"));

function Problem() {
    return (
        <motion.div
            className=""
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
        >
            <Stepper />
        </motion.div>
    );
}

export default Problem;
