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
    <div className="flex flex-col my-8">
      <label htmlFor={id} className="text-gray-800 font-medium tracking-wide">
        {label}
      </label>
      {children}
      <span className="h-6 text-red-500 mt-1">{error?.message}</span>
    </div>
  );
}
