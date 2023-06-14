import React from "react";

import { motion } from "framer-motion";

const ProblemNav = React.lazy(() => import("../component/Problem/Problem.Nav"));
const ProblemBody = React.lazy(() => import("../component/Problem/Problem.Body"));

function Problem() {
    return (
        <motion.div
            className="h-full w-full"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
        >
            <ProblemNav />
            <ProblemBody />
        </motion.div>
    );
}

export default Problem;
