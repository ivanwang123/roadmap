import { zodResolver } from "@hookform/resolvers/zod";
import React, { ReactNode } from "react";
import { SubmitHandler, useForm, UseFormReturn } from "react-hook-form";
import { ZodType, ZodTypeDef } from "zod";

type Props<TFormValues, Schema> = {
  schema: Schema;
  onSubmit: SubmitHandler<TFormValues>;
  children: (method: UseFormReturn<TFormValues>) => ReactNode;
};

export function Form<
  TFormValues extends Record<string, unknown> = Record<string, unknown>,
  Schema extends ZodType<unknown, ZodTypeDef, unknown> = ZodType<
    unknown,
    ZodTypeDef,
    unknown
  >
>({ schema, onSubmit, children }: Props<TFormValues, Schema>) {
  const methods = useForm<TFormValues>({
    resolver: zodResolver(schema),
  });

  return (
    <form className="flex flex-col" onSubmit={methods.handleSubmit(onSubmit)}>
      {children(methods)}
    </form>
  );
}
