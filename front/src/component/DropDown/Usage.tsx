import { DropDown } from "./DropDown";

export default function Usage(props: ProviderValue) {
  return (
    <DropDown value={props}>
      <DropDown.Wrapper>
        <DropDown.Label>
          <DropDown.Selected />
          <DropDown.Polygon />
        </DropDown.Label>
        <DropDown.MenuWrapper>
          <DropDown.Menu />
        </DropDown.MenuWrapper>
      </DropDown.Wrapper>
    </DropDown>
  );
}
