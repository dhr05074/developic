import { selectOptionState } from "@/recoil/problem.recoil";
import { useState } from "react";
import { useRecoilState } from "recoil";
const useSelectComponent = () => {
    const [languages, setLanguages] = useState(["Javascript", "Go", "Cpp"]);
    const [difficultList, setDifficultList] = useState(["Hard", "Normal", "Easy"]);
    const [selectOptoin, setSelectOption] = useRecoilState(selectOptionState);

    const initSelectOption = () => {
        setLang("");
        setDifficulty("");
    };

    const setLang = (value: LanguageType) => {
        setSelectOption((prevState) => {
            return { ...prevState, currentLang: value };
        });
    };
    const setDifficulty = (value: difficultyType) => {
        setSelectOption((prevState) => {
            return { ...prevState, defaultDifficulty: value };
        });
    };

    return {
        languages,
        difficultList,
        selectOptoin,
        setLang,
        setDifficulty,
        initSelectOption,
    };
};

export default useSelectComponent;
