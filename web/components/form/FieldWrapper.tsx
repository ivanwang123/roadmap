import React, { ReactNode } from "react";
import { FieldError } from "react-hook-form";

type Props = {
  id: string;
  label: string;
  error: FieldError | undefined;
  children: ReactNode;
};

export function FieldWrapper({ id, label, error, children }: Props) {
  return (
    <div className="flex flex-col w-full my-2">
      <label
        htmlFor={id}
        className="text-gray-500 text-sm font-medium tracking-wide mb-1"
      >
        {label}
      </label>
      {children}
      <span className="h-4 text-red-500 text-sm mt-1">{error?.message}</span>
    </div>
  );
}
