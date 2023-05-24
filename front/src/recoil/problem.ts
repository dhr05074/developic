import { DefaultValue, atom, selector } from "recoil";
import { generateProblem } from "@/api/problem.api";

// recoil을 model로 사용한다.
// problem 타입이 any라서 일단..
export const problemIdState = atom<any | undefined>({
    key: "problemIdState",
    default: "",
});
export const problemState = atom<any | undefined>({
    key: "problemState",
    default: "",
});
// recoil에서 비동기를 사용하려면 suspence를 사용해야한다.
export const setProblemId = selector({
    key: "problemId/set",
    get: ({ get }) => get(problemIdState),
    set: ({ set }, id) => set(problemIdState, id instanceof DefaultValue ? id : id),
});
export const setProblem = selector({
    key: "problemState/set",
    get: ({ get }) => {
        return get(problemState);
    },
    set: ({ set }, data) => set(problemState, data),
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
