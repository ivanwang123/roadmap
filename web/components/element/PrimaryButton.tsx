import React, { forwardRef } from "react";

type Props = {
  label?: string;
} & React.DetailedHTMLProps<
  React.ButtonHTMLAttributes<HTMLButtonElement>,
  HTMLButtonElement
>;

export const PrimaryButton = forwardRef<HTMLButtonElement, Props>(
  (props, ref) => {
    const { label, children, ...rest } = props;

    return (
      <button
        ref={ref}
        type="button"
        className="flex items-center justify-center text-white text-sm font-medium tracking-wide bg-blue-500 px-6 py-1.5 rounded-sm transition duration-300 hover:bg-blue-600 active:bg-blue-700 focus:outline-none"
        {...rest}
      >
        {children}
        {label}
      </button>
    );
  }
);
