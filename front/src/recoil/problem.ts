import { atom, selector } from "recoil";

import { generateProblem } from "@/api/problem";

// const createProblem = async () => {
//     const problem = await generateProblem().create(language, difficulty);
//     // problem.request_id;
// };

// const getProblem = async (requestId: string) => {
//     const problem = await generateProblem().get(requestId);
//     console.log(problem);
// };

// recoil을 model로 사용한다.
// problem 타입이 any라서 일단..
export const problemState = atom<any | undefined>({
    key: "problemState",
    default: null,
});

export const fontSizeLabel = selector({
    key: "updatedFontSize",
    get: ({ get }) => {
        const fontSize = get(fontSizeState);
        const unit = "px";

        return `${fontSize}${unit}`;
    },
});

// use case
// const [color, setColor] = useRecoilState(colorState)

// const onChange = (e: React.MouseEvent<HTMLDivElement>) => {
//   if (color === '#FBFBFE') {
//     setColor('#ffa000')
//   } else {
//     setColor('#FBFBFE')
//   }
// }
