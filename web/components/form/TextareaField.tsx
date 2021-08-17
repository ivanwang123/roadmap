import clsx from "clsx";
import React from "react";
import { FieldError, UseFormRegisterReturn } from "react-hook-form";
import { FieldWrapper } from "./FieldWrapper";

type Props = {
  id: string;
  label: string;
  register: UseFormRegisterReturn;
  error: FieldError | undefined;
  type?: string;
};

export function TextareaField({ id, label, register, error }: Props) {
  return (
    <FieldWrapper id={id} label={label} error={error}>
      <textarea
        id={id}
        rows={4}
        className={clsx(
          "border-b-2 pt-1 resize-none focus:border-gray-800 focus:outline-none",
          error && "border-red-500"
        )}
        {...register}
      />
    </FieldWrapper>
  );
}
