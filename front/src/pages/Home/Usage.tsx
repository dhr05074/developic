//refactoring
import Select from "@/component/DropDown/Usage";
import { Home } from "./Home";
import mainImage from "@/assets/images/main_image.svg";
function PageHome() {
  // UI하나만 볼 수 있다.
  const description =
    "Here at Flowbite we focus on markets where technology, innovation, ,and capital can unlock long-term value and drive economic growth.";
  return (
    <Home>
      <Home.Image image={mainImage} />
      <Home.Wrapper>
        <Home.Title title={"Code Refactoring"} />
        <Home.Description description={description} />
        {/* select 만든 후 리팩. */}
        <Home.SelectWrapper>
          <Select
            menu={["Javascript", "Go", "Cpp"]}
            disabled="Language"
            location="left-0"
            size="large"
          />
          <Select
            menu={["Hard", "Normal", "Easy"]}
            disabled="Difficulty"
            location="right-0"
            size="large"
          />
        </Home.SelectWrapper>
        <Home.ButtonWrapper>
          <Home.Button name="Start" />
        </Home.ButtonWrapper>
      </Home.Wrapper>
    </Home>
  );
}
export { PageHome };
