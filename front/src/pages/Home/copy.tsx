import SelectComponent from "@/component/Select/Select";
import ButtonLink from "@/component/Button/Link.Button";
import mainImage from "@/assets/images/main_image.svg";
import useSelectComponent from "@/hook/SelectComponent.hook";
import { useEffect, useState } from "react";
import useProfile from "@/hook/Profile.hook";

//refactoring
import { Home } from "./Home";
export default function Home() {
  const [languages, setLanguages] = useState<LanguageType[]>([
    "Javascript",
    "Go",
    "Cpp",
  ]);
  const [difficultList, setDifficultList] = useState<difficultyType[]>([
    "Hard",
    "Normal",
    "Easy",
  ]);
  const {
    setLang,
    setDifficulty,
    initSelectOption,
    optionLength,
    optionDifficulty,
  } = useSelectComponent();
  const { setAuth } = useProfile();
  const buttonOption = {
    pathName: "/stepper",
    search: `?difficulty=${optionDifficulty}&language=${optionLength}`,
  };
  useEffect(() => {
    setAuth();
    initSelectOption();
  }, []);

  // Home.??
  return (
    <Home onChange={handleChangeCounter}>
      {/* Main image */}
      <article className="flex h-full w-1/2 flex-row items-center justify-center">
        <img className="h-[40%]" src={mainImage} alt="refactor your code" />
      </article>
      {/*  */}
      <article className=" h-full w-1/2 gap-4">
        <div className="flex h-full  w-[440px] flex-col items-start justify-center  gap-4 text-left text-white">
          {/* Main Title */}
          <h3 className="pretendard_extrabold_32 w-full">Code Refactoring</h3>
          {/* Main Description */}
          <p className="pretendard_medium_20 w-11/12">
            Here at Flowbite we focus on markets where technology, innovation,
            ,and capital can unlock long-term value and drive economic growth.
          </p>

          <section
            id="home_selectButtons"
            className="relative mt-10 flex h-14 w-full  flex-row"
          >
            {/* 추상화 수준이 섞여있다. 리팩토링 필요. */}
            <SelectComponent
              value={{ menu: languages, callback: setLang }}
              disabled="Language"
              location="left-0"
              size="large"
            />
            <SelectComponent
              value={{ menu: difficultList, callback: setDifficulty }}
              disabled="Difficulty"
              location="right-0"
              size="large"
            />
          </section>
          <section
            id="home_startButton"
            className="flex w-full flex-row justify-center"
          >
            <div className="mt-3 w-[80%]">
              {optionLength && optionDifficulty ? (
                <ButtonLink link={buttonOption} name="Start" />
              ) : (
                <ButtonLink name="Start" />
              )}
            </div>
          </section>
        </div>
      </article>
    </Home>
  );
}
