import { useDropDownContext } from "../useDropDownContext";

function Menu() {
  const { menu, onClickMenu, isClick } = useDropDownContext();

  if (!isClick) return null;
  return (
    <>
      {menu.map((m, index) => (
        <li
          onClick={onClickMenu}
          key={m + index}
          className="motion_basic hover:text-coco-green_500 "
        >
          {m}
        </li>
      ))}
    </>
  );
}

export { Menu };
