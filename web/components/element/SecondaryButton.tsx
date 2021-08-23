import React, { forwardRef } from "react";

type Props = {
  label?: string;
} & React.DetailedHTMLProps<
  React.ButtonHTMLAttributes<HTMLButtonElement>,
  HTMLButtonElement
>;

export const SecondaryButton = forwardRef<HTMLButtonElement, Props>(
  (props, ref) => {
    const { label, children, ...rest } = props;

    return (
      <button
        ref={ref}
        type="button"
        className="flex items-center justify-center text-black text-sm font-medium tracking-wide border border-gray-200 px-6 py-2 rounded-sm transition duration-300 hover:bg-gray-200 active:bg-gray-300 focus:outline-none"
        {...rest}
      >
        {children}
        {label}
      </button>
    );
  }
);
