import { useDropDownContext } from "../useDropDownContext";

function Selected() {
  const { selected } = useDropDownContext();
  return <p>{selected}</p>;
}

export { Selected };
