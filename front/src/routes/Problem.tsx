import * as React from "react";

import { useState, useEffect, Suspense } from "react";
import { motion } from "framer-motion";
import { useSearchParams } from "react-router-dom";
import Stepper from "@/component/Stepper/Stepper";
import useProblem from "@/hook/Problem.hook";

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
            <ErrorBoundary fallback={<Stepper />}>
                <Suspense fallback={<Stepper />}>
                    <ProblemComponent />
                </Suspense>
            </ErrorBoundary>
        </motion.div>
    );
}
class ErrorBoundary extends React.Component {
    constructor(props) {
        super(props);
        this.state = { hasError: false };
    }

    static getDerivedStateFromError(error) {
        // 다음 렌더링에서 폴백 UI가 보이도록 상태를 업데이트 합니다.
        return { hasError: true };
    }

    componentDidCatch(error, errorInfo) {
        // 에러 리포팅 서비스에 에러를 기록할 수도 있습니다.
        //   logErrorToMyService(error, errorInfo);
    }

    render() {
        if (this.state.hasError) {
            // 폴백 UI를 커스텀하여 렌더링할 수 있습니다.
            // <Stepper />;
            return null;
        }
        return this.props.children;
    }
}
//   <ErrorBoundary>
//     <MyWidget />
//   </ErrorBoundary>;
export default Problem;
