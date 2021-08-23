import { zodResolver } from "@hookform/resolvers/zod";
import React, { ReactNode } from "react";
import { FormProvider, SubmitHandler, useForm } from "react-hook-form";
import { ZodType, ZodTypeDef } from "zod";

type Props<TFormValues, Schema> = {
  schema: Schema;
  onSubmit: SubmitHandler<TFormValues>;
  children: ReactNode;
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
    <FormProvider {...methods}>
      <form className="flex flex-col" onSubmit={methods.handleSubmit(onSubmit)}>
        {children}
      </form>
    </FormProvider>
  );
}
