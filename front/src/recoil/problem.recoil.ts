import { atom, selector } from "recoil";
import { Problem, ProgrammingLanguage } from "api/api";
import { api } from "@/api/defaultApi";
import { useSearchParams } from "react-router-dom";
import { profileState } from "./profile.recoil";

// recoil을 model로 사용한다.
// problem 타입이 any라서 일단..


export const selectOptionState = atom({
    key: "selectOptions",
    default: {
        currentLang: "" as LanguageType,
        defaultDifficulty: "" as difficultyType,
    },
});
// get: async ({get}) => {
//     const response = await myDBQuery({
//       userID: get(currentUserIDState),
//     });
//     return response.name;
//   },
export const SelectLang = selector({
    "key":"selectOptionLang",
    get:({get}) => {
        return get(selectOptionState).currentLang
    },
    set: ({set}, newValue:LanguageType) => {
        set(selectOptionState, (prevState) => ({
            ...prevState,
            currentLang: newValue,
          }));
    }
})
export const SelectDifficulty = selector({
    "key":"selectOptionDifficulty",
    get:({get}) => {
        return get(selectOptionState).defaultDifficulty
    },
    set: ({set}, newValue:difficultyType) => {
        set(selectOptionState, (prevState) => ({
            ...prevState,
            defaultDifficulty: newValue,
          }));
    }
})
// const getCreate = await api.requestProblem(
//     {
//         language: getLanguage,
//     },
//     { headers: profile.headers },
// );
// export const problemState = atom({
//     key: "problem",
//     default: {
//         id,
//         data
//     },
// });

export const problemIdSelector = selector({ // 비동기 셀렉터를 사용하면 readOnly
    key:"selectProblemId",
    get:async ({get}) => {
       const profile = get(profileState)
        const response = await api.requestProblem({
            language: get(selectOptionState).currentLang as ProgrammingLanguage
        },{headers: profile.headers });
        return response.data.problem_id
    }
})
export const problemSelector = selector({
    key:"selectProblem",
    get:async ({get}) => {
        const profile = get(profileState)
        const id = get(problemIdSelector)
        const response = await api.getProblem(id, {headers:profile.headers})
        return response.data
    },
})
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