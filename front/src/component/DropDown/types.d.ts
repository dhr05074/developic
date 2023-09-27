type DropDownProps = {
  menu: string[];
  disabled: string;
  location: "left-0" | "right-0";
  size: "small" | "large";
};

type ProviderValue = {
  menu: string[];
  disabled: string;
  location: "left-0" | "right-0";
  size: "small" | "large";
  selected: string;
  isClick: boolean;
  onClickDropDown: () => void;
  onClickMenu: (e: React.MouseEvent<HTMLLIElement>) => void;
};
