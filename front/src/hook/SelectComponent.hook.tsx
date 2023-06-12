import { difficultState, languageState, selectOptionState } from "@/recoil/problem.recoil";
import { useRecoilState } from "recoil";
const useSelectComponent = () => {
    const [languages, setLanguages] = useRecoilState(languageState);
    const [difficultList, setDifficultList] = useRecoilState(difficultState);
    const [selectOptoin, setSelectOption] = useRecoilState(selectOptionState);

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
    };
};

export default useSelectComponent;
