import { graphql, ChildDataProps } from "@apollo/react-hoc";
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

export function withAuth<TInputProps = {}>() {
  return graphql<TInputProps, Response, {}, WithAuthProps<TInputProps>>(
    ME_QUERY
  );
}
