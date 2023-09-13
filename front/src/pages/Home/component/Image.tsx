type PropType = {
  image: string;
};

function Image({ image }: PropType) {
  return (
    <article className="flex h-full w-1/2 flex-row items-center justify-center">
      <img className="h-[40%]" src={image} alt="refactor your code" />
    </article>
  );
}

export { Image };
