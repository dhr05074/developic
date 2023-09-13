import React, { useEffect, useMemo, useState } from "react";
import {
  Title,
  Description,
  Image,
  Wrapper,
  Select,
  SelectWrapper,
  Button,
  ButtonWrapper,
} from "./component";
import { HomeProvider } from "./useHomeContext";
import useSelectComponent from "@/hook/SelectComponent.hook";
import useProfile from "@/hook/Profile.hook";

function Home({ children }: { children: React.ReactNode }) {
  // 로직은 전부 여기서 관리.
  // const [clapState, setClapState] = useState()
  const { setLang, initSelectOption } = useSelectComponent();
  const { setAuth } = useProfile();
  useEffect(() => {
    setAuth();
    initSelectOption();
  }, []);

  // const memoizedValue = useMemo( // 다시 확인 필요
  //   () => ({
  //     ...clapState,

  //   }),
  //   [clapState]
  // )
  // value가 자식컴포넌트에 보내는 요소.
  // 여기서 로직을 작성하여 value로 값을 보내면된다.
  // 자식컴포넌트는 useContext로 로직의 값을 변경할 수 있다.
  return (
    <HomeProvider value={{ setLang }}>
      <section className="flex h-full w-full flex-row  bg-Navy-900">
        {children}
      </section>
    </HomeProvider>
  );
}

// const StyledCounter = styled.div`
//   display: inline-flex;
//   border: 1px solid #17a2b8;
//   line-height: 1.5;
//   border-radius: 0.25rem;
//   overflow: hidden;
// `;
Home.Image = Image;
Home.Wrapper = Wrapper;
Home.Title = Title;
Home.Description = Description;
Home.Select = Select;
Home.SelectWrapper = SelectWrapper;
Home.Button = Button;
Home.ButtonWrapper = ButtonWrapper;

export { Home };
