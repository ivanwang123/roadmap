import React, { forwardRef } from "react";

type Props = {
  label: string;
} & React.DetailedHTMLProps<
  React.ButtonHTMLAttributes<HTMLButtonElement>,
  HTMLButtonElement
>;

export const LinkButton = forwardRef<HTMLButtonElement, Props>((props, ref) => {
  const { label, ...rest } = props;

  return (
    <button
      ref={ref}
      type="button"
      className="primary-font text-gray-400 text-sm font-medium tracking-wide cursor-pointer transition duration-300 hover:text-black"
      {...rest}
    >
      {label}
    </button>
  );
});
