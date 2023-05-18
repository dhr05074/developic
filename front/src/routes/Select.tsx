import { Link } from "react-router-dom";
import React from "react";
import { motion } from "framer-motion";
import SelectComponent from "@/component/Select/Select";
import ButtonBasic from "@/component/Button/Basic.Button";

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
    let c_difficulty = "normal";
    const setLang = (value: LanguageType) => {
        currentLang = value;
    };
    const setDifficulty = (value: string) => {
        c_difficulty = value;
    };
    const buttonOption = {
        pathName: "/problem",
        search: `?difficulty=${c_difficulty}&language=${currentLang}`,
    };
    return (
        <motion.div
            className=""
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            exit={{ opacity: 0 }}
            transition={{ duration: 0.5 }}
        >
            <div>
                <section className="flex h-screen w-screen flex-col items-center justify-center gap-5">
                    <div className="flex flex-row items-center justify-center gap-4">
                        <p>언어 : </p>
                        <SelectComponent value={{ menu: languages, callback: setLang }} />
                        <p>난이도 : </p>
                        <SelectComponent value={{ menu: difficulty, callback: setDifficulty }} />
                    </div>
                    <ButtonBasic link={buttonOption} name="문제 풀기" />
                </section>
            </div>
        </motion.div>
    );
}
