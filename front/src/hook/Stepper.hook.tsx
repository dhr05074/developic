import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

const useStepper = () => {
    const navigate = useNavigate();

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
    const stepperStateChanger = (getStep: StepType) => {
        if (getStep === "start") {
            StepperList.difficult.step = "loading";
            const changeDefficult = StepperList.difficult;
            setStepperList((prevState) => {
                return { ...prevState, difficult: changeDefficult };
            });
        } else if (getStep === "level") {
            StepperList.difficult.step = "complete";
            StepperList.language.step = "loading";
            const changeDefficult = StepperList.difficult;
            const changeLanguage = StepperList.language;
            setStepperList((prevState) => {
                return { ...prevState, difficult: changeDefficult, language: changeLanguage };
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

    return {
        StepperList,
        stepperStateChanger,
    };
};

export default useStepper;
