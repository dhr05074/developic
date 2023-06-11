import { atom } from "recoil";
import { Problem } from "api/api";

// recoil을 model로 사용한다.
// problem 타입이 any라서 일단..

export const languageState = atom<LanguageType[]>({
    key: "languageList",
    default: ["Javascript", "Go", "Cpp"],
});
export const difficultState = atom<difficultyType[]>({
    key: "difficultList",
    default: ["Hard", "Normal", "Easy"],
});
export const selectOptionState = atom({
    key: "selectOptions",
    default: {
        currentLang: "" as LanguageType,
        defaultDifficulty: "" as difficultyType,
    },
});

export const problemIdState = atom<string | null>({
    key: "problemId",
    default: null,
});
export const problemState = atom<Problem | null>({
    key: "problem",
    default: null,
});

// recoil에서 비동기를 사용하려면 suspence를 사용해야한다.
// export const setProblemId = selector({
//     key: "problemId/set",
//     get: ({ get }) => get(problemIdState),
//     set: ({ get, set }, id) => {
//         set(problemIdState, id);
//         console.log(get(problemIdState));
//     },
// });
// export const setProblem = selector({
//     key: "problemState/set",
//     get: ({ get }) => {
//         return get(problemState);
//     },
//     set: ({ set }, data) => set(problemState, data),
// });

// use case
// const [color, setColor] = useRecoilState(colorState)

// const onChange = (e: React.MouseEvent<HTMLDivElement>) => {
//   if (color === '#FBFBFE') {
//     setColor('#ffa000')
//   } else {
//     setColor('#FBFBFE')
//   }
// }
