// eslint-disable-next-line import/extensions,import/no-unresolved
import React from "react";
import { motion } from "framer-motion";

import ProblemComponent from "../component/Problem/Problem";
import Stepper from "@/component/Stepper/Stepper";

const resource = fetchProfileData();

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
                <ProblemComponent />
            </React.Suspense>
        </motion.div>
    );
}

function ProfileDetails() {
    const user = resource.user.read();
    return <h1>{user.name}</h1>;
}
export default Problem;
