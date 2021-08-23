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
  React.TextareaHTMLAttributes<HTMLTextAreaElement>,
  HTMLTextAreaElement
>;

export function TextareaField({ id, label, ...rest }: Props) {
  const {
    register,
    formState: { errors },
  } = useFormContext();

  return (
    <FieldWrapper id={id} label={label} error={get(errors, id)}>
      <textarea
        id={id}
        rows={4}
        className={clsx(
          "border tracking-wide px-2 py-0.5 rounded-sm resize-none focus:outline-none",
          get(errors, id) ? "border-red-500" : "focus:border-gray-500"
        )}
        {...register(id)}
        {...rest}
      />
    </FieldWrapper>
  );
}
