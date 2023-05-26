import { Link } from "react-router-dom";
import React from "react";
import { motion } from "framer-motion";
import SelectComponent from "@/component/Select/Select";
import ButtonBasic from "@/component/Button/Basic.Button";
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
const difficulty: string[] = ["hard", "normal", "easy"];
export default function Select() {
    let currentLang = "javascript" as LanguageType;
    let defaultDifficulty = "normal";
    const setLang = (value: LanguageType) => {
        currentLang = value;
    };
    const setDifficulty = (value: string) => {
        defaultDifficulty = value;
    };
    const buttonOption = {
        pathName: "/stepper",
        search: `?difficulty=${defaultDifficulty}&language=${currentLang}`,
    };
    return (
        <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
        >
            <div>
                <section className="flex h-screen w-screen flex-row  bg-Navy-900">
                    <article className="flex h-full w-1/2 flex-row items-center justify-center">
                        <img className="h-[40%]" src={mainImage} alt="refactor your code" />
                    </article>
                    <article className=" h-full w-1/2 gap-4">
                        <div className="flex h-full  w-[400px] flex-col items-center  justify-center text-left">
                            <h3 className="pretendard_extrabold_32 w-full">Code Refactoring</h3>
                            <p className="pretendard_medium_20 w-full">
                                Here at Flowbite we focus on markets where technology, innovation, ,and capital can
                                unlock long-term value and drive economic growth.
                            </p>
                            <div className="flex flex-row gap-2">
                                <SelectComponent value={{ menu: languages, callback: setLang }} />
                                <SelectComponent value={{ menu: difficulty, callback: setDifficulty }} />
                            </div>
                            <ButtonBasic link={buttonOption} name="Start" />
                        </div>
                    </article>
                </section>
            </div>
        </motion.div>
    );
}
