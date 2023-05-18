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
// recoil에서 비동기를 사용하려면 suspence를 사용해야한다.
export const setProblemState = selector({
    key: "problemState/set",
    get: ({ get }, kind: "data" | "id") => {
        if (kind === "id") return get(problemIdState);
        return get(problemState);
    },
    setId: ({ set, get }, id: string) => {
        const problem = get(problemIdState);
        problem.id = id;
        set(problemState, problem);
    },
    setData: ({ set }, data) => {
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
