import React, { ReactNode, useState, useEffect, Suspense } from "react";

import { motion } from "framer-motion";

const ProblemComponent = React.lazy(() => import("../component/Problem/Problem.Component"));

function Problem() {
    return (
        <motion.div
            className="h-full w-full"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
        >
            <ProblemComponent />
        </motion.div>
    );
}

export default Problem;
