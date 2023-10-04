//refactoring
import DropDown from "@/component/DropDown/Usage";
import { Home } from "./Home";
import mainImage from "@/assets/images/main_image.svg";
import { Button } from "@/component/Button/Button";

function PageHome({
  onNext,
}: {
  onNext: React.Dispatch<React.SetStateAction<string>>;
}) {
  // UI하나만 볼 수 있다.
  // to={{ pathname: props.link?.pathName, search: props.link?.search }}

  const description =
    "Here at Flowbite we focus on markets where technology, innovation, ,and capital can unlock long-term value and drive economic growth.";
  return (
    <Home>
      <Home.Image image={mainImage} />
      <Home.Wrapper>
        <Home.Title title={"Code Refactoring"} />
        <Home.Description description={description} />
        <Home.DropDownWrapper>
          <DropDown
            menu={["Javascript", "Go", "Cpp"]}
            disabled="Language"
            location="left-0"
            size="large"
          />
          <DropDown
            menu={["Hard", "Normal", "Easy"]}
            disabled="Difficulty"
            location="right-0"
            size="large"
          />
        </Home.DropDownWrapper>
        <Home.ButtonWrapper>
          {/* Button 리팩터링하기 */}
          <Button name="Start" action={onNext} />
          {/* <Home.Button name="Start" /> */}
        </Home.ButtonWrapper>
      </Home.Wrapper>
    </Home>
  );
}
export { PageHome };
