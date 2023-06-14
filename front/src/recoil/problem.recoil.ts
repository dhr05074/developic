import { atom } from "recoil";
import { Problem } from "api/api";

// recoil을 model로 사용한다.
// problem 타입이 any라서 일단..

// effects를 하나의 함수로 묶을려고 시도했지만 타입이 좀 까다로움.

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
    effects: [
        ({ setSelf, onSet }) => {
            const savedData = localStorage.getItem(selectOptionState.key);
            // setSelf: atom 값을 설정 혹은 재설정
            if (savedData) setSelf(JSON.parse(savedData));

            // atom이 변화가 감지될 때 작동, Storage에 데이터 저장
            // setSelf에 의해서는 작동하지 않음
            onSet((newValue, _, isReset) => {
                isReset
                    ? localStorage.removeItem(selectOptionState.key)
                    : localStorage.setItem(selectOptionState.key, JSON.stringify(newValue));
            });
        },
    ],
});

export const problemIdState = atom<string | null>({
    key: "problemId",
    default: null,
});
export const problemState = atom<Problem | null>({
    key: "problem",
    default: null,
    effects: [
        ({ setSelf, onSet }) => {
            const savedData = localStorage.getItem(problemState.key);
            // setSelf: atom 값을 설정 혹은 재설정
            if (savedData) setSelf(JSON.parse(savedData));

            // atom이 변화가 감지될 때 작동, Storage에 데이터 저장
            // setSelf에 의해서는 작동하지 않음
            onSet((newValue, _, isReset) => {
                isReset
                    ? localStorage.removeItem(problemState.key)
                    : localStorage.setItem(problemState.key, JSON.stringify(newValue));
            });
        },
    ],
});
export const editorInCode = atom<string | undefined>({
    key: "editorInCode",
    default: undefined,
    effects: [
        ({ setSelf, onSet }) => {
            const savedData = localStorage.getItem(editorInCode.key);
            // setSelf: atom 값을 설정 혹은 재설정
            if (savedData) setSelf(JSON.parse(savedData));

            // atom이 변화가 감지될 때 작동, Storage에 데이터 저장
            // setSelf에 의해서는 작동하지 않음
            onSet((newValue, _, isReset) => {
                isReset
                    ? localStorage.removeItem(editorInCode.key)
                    : localStorage.setItem(editorInCode.key, JSON.stringify(newValue));
            });
        },
    ],
});

const effectHandler = () => {};
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
