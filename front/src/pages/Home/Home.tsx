import React, { useEffect, useMemo, useState } from "react";
import {
  Title,
  Description,
  Image,
  Wrapper,
  SelectWrapper,
  Button,
  ButtonWrapper,
} from "./component";
import { HomeProvider } from "./useHomeContext";
import useProfile from "@/hook/Profile.hook";

function Home({ children }: { children: React.ReactNode }) {
  const { setAuth } = useProfile();
  useEffect(() => {
    setAuth();
  }, []);

  return (
    <HomeProvider value={undefined}>
      <section className="flex h-full w-full flex-row  bg-Navy-900">
        {children}
      </section>
    </HomeProvider>
  );
}

Home.Image = Image;
Home.Wrapper = Wrapper;
Home.Title = Title;
Home.Description = Description;
Home.SelectWrapper = SelectWrapper;
Home.Button = Button;
Home.ButtonWrapper = ButtonWrapper;

export { Home };
