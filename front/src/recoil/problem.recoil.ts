import { atom, selector } from "recoil";
import { Problem } from "api/api";

// recoil을 model로 사용한다.
// problem 타입이 any라서 일단..

// effects를 하나의 함수로 묶을려고 시도했지만 타입이 좀 까다로움.

// export const languageState = atom<LanguageType[]>({
//     key: "languages",
//     default: ["Javascript", "Go", "Cpp"],
// });

// export const difficultState = atom<difficultyType[]>({
//     key: "difficultList",
//     default: ["Hard", "Normal", "Easy"],
// });
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
export const optionSelector = selector({
    "key":"selectOptionsSelector",
    get:({get}) => {
    return get(selectOptionState)
    }
})
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
            console.log("🚀 ~ file: problem.recoil.ts:65 ~ savedData:", savedData)
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

// const currentUserNameQuery = selector({
//     key: 'CurrentUserName',
//     get: async ({get}) => {
//       const response = await myDBQuery({
//         userID: get(currentUserIDState),
//       });
//       return response.name;
//     },
//   });
  
//   function CurrentUserInfo() {
//     const userName = useRecoilValue(currentUserNameQuery);
//     return <div>{userName}</div>;
//   }