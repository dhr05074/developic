import { useState } from "react";
import { motion } from "framer-motion";
import SelectComponent from "@/component/Select/Select";
import ButtonLink from "@/component/Button/Link.Button";
import mainImage from "@/assets/main_image.svg";

const languages: LanguageType[] = [
    "javascript",
    "go",
    "c++",

    // "typescript",
    // "python3",
    // "cpp",
    // "c#",
    // "clojure",
    // "dart",
    // "elixir",
    // "java",
    // "kotlin",
    // "php",
    // "r",
    // "ruby",
    // "scala",
    // "swift",
];
const difficulty: difficultyType[] = ["hard", "normal", "easy"];
export default function Select() {
    const [selectOptoin, setSelectOption] = useState({
        currentLang: "" as LanguageType,
        defaultDifficulty: "",
    });
    const setLang = (value: LanguageType) => {
        setSelectOption((prevState) => {
            return { ...prevState, currentLang: value };
        });
    };
    const setDifficulty = (value: string) => {
        setSelectOption((prevState) => {
            return { ...prevState, defaultDifficulty: value };
        });
    };
    const buttonOption = {
        pathName: "/stepper",
        search: `?difficulty=${selectOptoin.defaultDifficulty}&language=${selectOptoin.currentLang}`,
    };
    return (
        <motion.div
            className="h-full w-full"
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
        >
            <section className="flex h-full w-full flex-row  bg-Navy-900">
                <article className="flex h-full w-1/2 flex-row items-center justify-center">
                    <img className="h-[40%]" src={mainImage} alt="refactor your code" />
                </article>
                <article className=" h-full w-1/2 gap-4">
                    <div className="flex h-full  w-[440px] flex-col items-start justify-center  gap-4 text-left">
                        <h3 className="pretendard_extrabold_32 w-full">Code Refactoring</h3>
                        <p className="pretendard_medium_20 w-11/12">
                            Here at Flowbite we focus on markets where technology, innovation, ,and capital can unlock
                            long-term value and drive economic growth.
                        </p>
                        <section id="home_selectButtons" className="relative mt-10 flex h-14 w-full  flex-row">
                            <SelectComponent
                                value={{ menu: languages, callback: setLang }}
                                disabled="Language"
                                location="left-0"
                            />
                            <SelectComponent
                                value={{ menu: difficulty, callback: setDifficulty }}
                                disabled="Difficulty"
                                location="right-0"
                            />
                        </section>
                        <section id="home_startButton" className="flex w-full flex-row justify-center">
                            <div className="mt-3 w-[80%]">
                                {selectOptoin.currentLang && selectOptoin.defaultDifficulty ? (
                                    <ButtonLink link={buttonOption} name="Start" />
                                ) : (
                                    <ButtonLink name="Start" />
                                )}
                            </div>
                        </section>
                    </div>
                </article>
            </section>
        </motion.div>
    );
}
