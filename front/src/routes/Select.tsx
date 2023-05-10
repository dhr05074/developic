import { Link } from "react-router-dom";
import SelectComponent from "@/component/Select/Select";
import TextArea from "@/component/Textarea/Textarea";

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
    return (
        <div>
            <section className="flex h-screen w-screen flex-col items-center justify-center gap-5">
                <div className="flex flex-row items-center justify-center gap-4">
                    <p>언어 : </p>
                    <SelectComponent value={{ menu: languages, callback: setLang }} />
                    <p>난이도 : </p>
                    <SelectComponent value={{ menu: difficulty, callback: setDifficulty }} />
                </div>
                <Link
                    to={{ pathname: "/codeEditor", search: `?difficulty=${c_difficulty}&language=${currentLang}` }}
                    className="motion_basic inline-flex items-center justify-center rounded-lg bg-coco-green_500 px-5 py-3 text-center text-base font-medium text-black  hover:bg-Navy-700 hover:text-white"
                >
                    문제 풀기
                </Link>
            </section>
        </div>
    );
}
