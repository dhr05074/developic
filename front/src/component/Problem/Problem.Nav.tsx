import Button from "../Button/Function.Button";
import SelectComponent from "../Select/Select";
import useSelectComponent from "@/hook/SelectComponent.hook";

function ProblemNav() {
    const { languages, difficultList, setLang, setDifficulty, selectOptoin } = useSelectComponent();
    return (
        <section id="problem_nav" className="flex w-full flex-row items-center justify-between bg-Navy-800 px-6 py-2">
            <article id="problem_nav_options" className=" relative flex h-10 w-52 flex-row">
                <SelectComponent
                    value={{ menu: languages, callback: setLang }}
                    disabled={selectOptoin.currentLang}
                    location="left-0"
                    size="small"
                />
                <SelectComponent
                    value={{ menu: difficultList, callback: setDifficulty }}
                    disabled={selectOptoin.defaultDifficulty}
                    location="right-0"
                    size="small"
                />
            </article>
            <article id="problem_nav_buttons" className="flex flex-row gap-3">
                <Button name="Reset" type="line" size="small" />
                <Button name="Submit" type="full" size="small" />
            </article>
        </section>
    );
}

export default ProblemNav;
