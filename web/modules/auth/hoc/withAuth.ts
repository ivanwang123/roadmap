import { graphql, ChildDataProps, DataProps } from "@apollo/client/react/hoc";
import { ComponentType } from "react";
import { UserInfoFieldsFragment } from "types/graphql-generated";
import { ME_QUERY } from "../api";

type Response = {
  me: UserInfoFieldsFragment;
};

export type WithAuthProps<TInputProps = {}> = ChildDataProps<
  TInputProps,
  Response,
  {}
>;

export function withAuth<TInputProps = {}>(
  component: ComponentType<TInputProps & DataProps<Response, {}>>
) {
  if (component)
    return graphql<TInputProps, Response, {}, WithAuthProps<TInputProps>>(
      ME_QUERY
    )(component);

  return component;
}
