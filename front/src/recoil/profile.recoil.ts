import { atom } from "recoil";
import { Problem } from "api/api";

// recoil을 model로 사용한다.
// problem 타입이 any라서 일단..

// effects를 하나의 함수로 묶을려고 시도했지만 타입이 좀 까다로움.
export const profileState = atom({
    key: "profile",
    default: {
        nickname: "", //key로 사용.
        elo_score: 0,
    },
    effects: [
        ({ setSelf, onSet }) => {
            const savedData = localStorage.getItem(profileState.key);
            // setSelf: atom 값을 설정 혹은 재설정
            if (savedData) setSelf(JSON.parse(savedData));

            // atom이 변화가 감지될 때 작동, Storage에 데이터 저장
            // setSelf에 의해서는 작동하지 않음
            onSet((newValue, _, isReset) => {
                isReset
                    ? localStorage.removeItem(profileState.key)
                    : localStorage.setItem(profileState.key, JSON.stringify(newValue));
            });
        },
    ],
});
