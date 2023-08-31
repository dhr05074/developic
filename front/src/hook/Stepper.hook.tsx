import { useEffect, useState } from "react";
import { redirect, useNavigate } from "react-router-dom";

const useStepper = () => {
    const navigate = useNavigate();

    const [currentStep, setCurrentStep] = useState<StepType>("idle");
    const [StepperList, setStepperList] = useState<StepperListTypes>({
        difficult: {
            value: "난이도 체크",
            step: "idle",
        },
        language: {
            value: "언어 체크",
            step: "idle",
        },
        api: {
            value: "API 불러오는 중",
            step: "idle",
        },
        comp: {
            value: "문제 출제 완료",
            step: "idle",
        },
    });

    const stepInit = () => {
        setStepperList({
            difficult: {
                value: "난이도 체크",
                step: "idle",
            },
            language: {
                value: "언어 체크",
                step: "idle",
            },
            api: {
                value: "API 불러오는 중",
                step: "idle",
            },
            comp: {
                value: "문제 출제 완료",
                step: "idle",
            },
        });
    };
    const stepperStateChanger = (getStep: StepType) => {
        setCurrentStep(getStep);
        if (getStep === "start") {
            StepperList.difficult.step = "loading";
            const changeDifficult = StepperList.difficult;
            setStepperList((prevState) => {
                return { ...prevState, difficult: changeDifficult };
            });
        } else if (getStep === "level") {
            StepperList.difficult.step = "complete";
            StepperList.language.step = "loading";
            const changeDifficult = StepperList.difficult;
            const changeLanguage = StepperList.language;
            setStepperList((prevState) => {
                return { ...prevState, difficult: changeDifficult, language: changeLanguage };
            });
        } else if (getStep === "lang") {
            StepperList.language.step = "complete";
            StepperList.api.step = "loading";
            const changeLanguage = StepperList.language;
            const changeApi = StepperList.api;
            setStepperList((prevState) => {
                return { ...prevState, language: changeLanguage, api: changeApi };
            });
        } else if (getStep === "api") {
            StepperList.api.step = "complete";
            StepperList.comp.step = "loading";
            const changeApi = StepperList.api;
            const changeComp = StepperList.comp;
            setStepperList((prevState) => {
                return { ...prevState, api: changeApi, comp: changeComp };
            });
        } else if (getStep === "clear") {
            StepperList.comp.step = "complete";
            const changeComp = StepperList.comp;
            setStepperList((prevState) => {
                return { ...prevState, comp: changeComp };
            });
        } else if (getStep === "end") {
            navigate("/problem");
        }
        return true;
    };
    useEffect(() => {
        let count = 0;
        let timer = 0;
        timer = window.setInterval(() => {
            count++;
            console.log("count", count);
            // if (selectStep === "clear") {
            //     stepperStateChanger("clear");
            //     return;
            // }
            if (count === 1) stepperStateChanger("start");
            if (count === 2) stepperStateChanger("level");
            if (count === 3) stepperStateChanger("lang");
            if (count === 4) stepperStateChanger("api");
            if (count === 4) clearInterval(timer);
        }, 1000);

        return () => {
            stepInit();
            clearInterval(timer);
            console.log("stepper hook unmount");
        };
    }, []);
    const endStep = () => {
        // Stepper.tsx에서 실행
        stepperStateChanger("clear");
        setTimeout(() => {
            stepperStateChanger("end");
        }, 1000);
    };
    return {
        StepperList,
        stepperStateChanger,
        stepInit,
        endStep,
    };
};

export default useStepper;
