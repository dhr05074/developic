import React, { useEffect } from "react";

import useStepper from "@/hook/Stepper.hook";
import useProblem from "@/hook/Problem.hook";
import { motion } from "framer-motion";

type StepTypes = "idle" | "loading" | "complete";
type StepperStyle = {
    [component: StepTypes]: StepTypes;
    idle: string;
    loading: string;
    complete: string;
};

function StepperComponent() {
    const { StepperList, stepperStateChanger } = useStepper();
    const { problemId, problem, getProblemData } = useProblem();

    // if (step === 1) {
    //     stepButton = "문제 풀러가기";
    // }
    const style: StepperStyle = {
        idle: "border-gray-300 bg-gray-100 text-gray-900",
        loading: "border-Navy-500 bg-Navy-600 text-white",
        complete: "border-green-300 bg-green-50 text-green-700",
        inviserble: "opacity-0 invisible",
    };
    const stepChanger = (selectStep?: StepType) => {
        const timer = (step: StepType, time: number) => {
            setTimeout(() => {
                if (step === "lang") {
                    if (problemId) {
                        timer("api", 1000);
                    }
                }
                if (step === "clear") {
                    timer("end", 1000);
                }
                stepperStateChanger(step);
            }, time);
        };
        if (selectStep) {
            timer(selectStep, 1000);
            return;
        }

        timer("start", 1000);
        timer("level", 2000);
        timer("lang", 3000);
    };
    useEffect(() => {
        if (problemId) {
            console.log("problemId", problemId);
            stepChanger();
            getProblemData();
        }
    }, [problemId]);
    useEffect(() => {
        if (problem) stepChanger("clear");
    }, [problem]);
    return (
        <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
        >
            <div
                className={`motion_basic absolute flex h-full w-full flex-row items-center justify-center bg-Navy-800 `}
            >
                <ol className="w-72 space-y-4">
                    {Object.entries(StepperList).map(([key, value]) => (
                        <li key={key}>
                            <div
                                className={`motion_basic  stepper w-full rounded-lg  border p-4
                                ${style[value.step as StepTypes]}
                               `}
                                role="alert"
                            >
                                <div className="flex items-center justify-between">
                                    <span className="sr-only" />
                                    <h3 className="font-medium">{value.value}</h3>

                                    {value.step === "loading" && (
                                        <div role="status">
                                            <svg
                                                aria-hidden="true"
                                                className="mr-2 h-8 w-8 animate-spin fill-coco-green_500 text-gray-200 dark:text-gray-600"
                                                viewBox="0 0 100 101"
                                                fill="none"
                                                xmlns="http://www.w3.org/2000/svg"
                                            >
                                                <path
                                                    d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                                                    fill="currentColor"
                                                />
                                                <path
                                                    d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                                                    fill="currentFill"
                                                />
                                            </svg>
                                            <span className="sr-only">Loading...</span>
                                        </div>
                                    )}
                                    {value.step === "complete" && (
                                        <svg
                                            aria-hidden="true"
                                            className="h-5 w-5"
                                            fill="currentColor"
                                            viewBox="0 0 20 20"
                                            xmlns="http://www.w3.org/2000/svg"
                                        >
                                            <path
                                                fillRule="evenodd"
                                                d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                                                clipRule="evenodd"
                                            />
                                        </svg>
                                    )}
                                </div>
                            </div>
                        </li>
                    ))}
                </ol>
            </div>
        </motion.div>
    );
}

export default StepperComponent;
