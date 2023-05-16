// import { useRecoilState } from "recoil";
import { useState } from "react";

export default function useStepper() {
    const [step, setStep] = useState<number>(6);
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

    const stepperStateChanger = () => {
        if (step === 6) {
            StepperList.difficult.step = "loading";
            const changeDefficult = StepperList.difficult;
            setStepperList((prevState) => {
                return { ...prevState, difficult: changeDefficult };
            });
        } else if (step === 5) {
            StepperList.difficult.step = "complete";
            StepperList.language.step = "loading";
            const changeDefficult = StepperList.difficult;
            const changeLanguage = StepperList.language;
            setStepperList((prevState) => {
                return { ...prevState, difficult: changeDefficult, language: changeLanguage };
            });
        } else if (step === 4) {
            StepperList.language.step = "complete";
            StepperList.api.step = "loading";
            const changeLanguage = StepperList.language;
            const changeApi = StepperList.api;
            setStepperList((prevState) => {
                return { ...prevState, language: changeLanguage, api: changeApi };
            });
        } else if (step === 3) {
            StepperList.api.step = "complete";
            StepperList.comp.step = "loading";
            const changeApi = StepperList.api;
            const changeComp = StepperList.comp;
            setStepperList((prevState) => {
                return { ...prevState, api: changeApi, comp: changeComp };
            });
        } else if (step === 2) {
            StepperList.comp.step = "complete";
            const changeComp = StepperList.comp;
            setStepperList((prevState) => {
                return { ...prevState, comp: changeComp };
            });
        } else if (step === 0) {
            return false;
        }
        setStep(step - 1);
        return true;
    };

    return {
        step,
        StepperList,
        stepperStateChanger,
    };
}
