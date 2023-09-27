import styled from "styled-components";

const inactive = " button_large coco_button_disabled disabled_button  ";
// .coco_button_disabled {
//     @apply inline-flex w-full cursor-not-allowed items-center justify-center bg-Navy-800 text-center text-Navy-500 hover:text-Navy-500;
//   }

const ButtonStyle = styled.button`
  display: inline-flex;
  width: 100%;
  align-items: center;
  justify-content: center;
  background-color: #b9ff47;
  color: black;
  transition: all 0.2;
  transition-timing-function: ease-in-out;
  border-radius: 9999px;
  &:hover {
    background-color: #4c495b;
    color: white;
  }
`;
const DisabledButton = styled.button``;

export { ButtonStyle, DisabledButton };
