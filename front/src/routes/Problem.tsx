import React, { ReactNode, useState, useEffect, Suspense } from "react";

import { motion } from "framer-motion";
import { useSearchParams } from "react-router-dom";
import Stepper from "@/component/Stepper/Stepper";
import wrapPromise from "@/api/Suspencer";

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
            <Suspense fallback={<Stepper />}>
                <ErrorBoundary fallback={<Stepper />}>
                    <ProblemComponent />
                </ErrorBoundary>
            </Suspense>
        </motion.div>
    );
}

class ErrorBoundary extends React.Component {
    constructor(props) {
        super(props);
        this.state = { hasError: false };
    }

    static getDerivedStateFromError() {
        return { hasError: true };
    }

    componentDidCatch(error, errorInfo) {
        // logErrorToMyService(error, errorInfo);
    }

    render() {
        if (this.state.hasError) {
            return <h1>Something went wrong.</h1>;
        }

        return this.props.children;
    }
}
export default Problem;
