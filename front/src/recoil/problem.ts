import { atom, selector } from "recoil";

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

export const setProblemState = selector({
    key: "problemState/set",
    get: async ({ get }, kind: "data" | "id") => {
        if (kind === "id") return get(problemIdState);
        return get(problemState);
    },
    setId: async ({ set, get }, id: string) => {
        const problem = get(problemIdState);
        problem.id = id;
        set(problemState, problem);
    },
    setData: async ({ set }, data) => {
        const problem = get(problemIdState);
        problem.data = data;
        set(problemState, problem);
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
