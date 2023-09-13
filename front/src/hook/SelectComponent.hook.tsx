import { SelectDifficulty, SelectLang } from "@/recoil/problem.recoil";
import { useRecoilState } from "recoil";
const useSelectComponent = () => {
    const [optionLength, setOptionLength] = useRecoilState(SelectLang);
    const [optionDifficulty, setOptionDifficulty] = useRecoilState(SelectDifficulty);

    const initSelectOption = () => {
        setLang("");
        setDifficulty("");
    };

    const setLang = (value: LanguageType) => {
        setOptionLength(value)

    };
    const setDifficulty = (value: difficultyType) => {
        setOptionDifficulty(value)

    };

    return {
        optionLength,
        optionDifficulty,
        setLang,
        setDifficulty,
        initSelectOption,
    };
};

export default useSelectComponent;
