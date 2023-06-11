import Button from "../Button/Function.Button";

function ProblemNav() {
    return (
        <section id="problem_nav" className="flex h-10 w-full flex-row items-center justify-between bg-Navy-800 px-8">
            <article id="problem_nav_options"></article>
            <article id="problem_nav_buttons" className="flex flex-row gap-2">
                <Button name="Reset" type="line" size="small" />
                <Button name="Submit" type="full" size="small" />
            </article>
        </section>
    );
}

export default ProblemNav;
