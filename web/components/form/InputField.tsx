import clsx from "clsx";
import { get } from "lodash";
import React from "react";
import { useFormContext } from "react-hook-form";
import { FieldWrapper } from "./FieldWrapper";

type Props = {
  id: string;
  label: string;
  type?: string;
} & React.DetailedHTMLProps<
  React.InputHTMLAttributes<HTMLInputElement>,
  HTMLInputElement
>;

export function InputField({
  id,
  label,
  type = "text",
  children,
  ...rest
}: Props) {
  const {
    register,
    formState: { errors },
  } = useFormContext();

  return (
    <FieldWrapper id={id} label={label} error={get(errors, id)}>
      <div className="flex">
        <input
          type={type}
          id={id}
          className={clsx(
            "w-full border tracking-wide px-2 py-0.5 rounded-sm focus:outline-none",
            get(errors, id) ? "border-red-500" : "focus:border-gray-500"
          )}
          {...register(id)}
          {...rest}
        />
        {children}
      </div>
    </FieldWrapper>
  );
}
