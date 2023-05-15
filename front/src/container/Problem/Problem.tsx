import { useRecoilState } from "recoil";

const createProblem = async () => {
    const problem = await generateProblem().create(language, difficulty);
    problem.request_id;
};

const getProblem = async (requestId: string) => {
    const problem = await generateProblem().get(requestId);
    console.log(problem);
};

export default function ErrorPage() {
    const [color, setColor] = useRecoilState(colorState);
    return (
        <div id="error-page">
            <h1>Oops!</h1>
            <p>Sorry, an unexpected error has occurred.</p>
            <p>
                <i>{error.statusText || error.message}</i>
            </p>
        </div>
    );
}
