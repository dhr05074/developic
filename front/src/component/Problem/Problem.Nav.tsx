import ButtonBasic from "../Button/Link.Button";

function ProblemNav() {
    return (
        <section id="problem_nav" className="flex h-10 w-full flex-row items-center justify-between bg-Navy-800">
            <article id="problem_nav_options"></article>
            <article id="problem_nav_buttons">
                <ButtonBasic name="Submit" />
            </article>
        </section>
    );
}

export default ProblemNav;
