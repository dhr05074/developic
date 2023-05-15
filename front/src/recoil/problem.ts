import { atom, selector } from "recoil";

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
