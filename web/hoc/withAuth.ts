import { graphql, ChildDataProps } from "@apollo/react-hoc";
import { UserInfoFieldsFragment } from "../graphql/generated/generated";
import { ME_QUERY } from "../graphql/queries/me";

type Response = {
  me: UserInfoFieldsFragment;
};

export type AuthChildProps<TInputProps = {}> = ChildDataProps<
  TInputProps,
  Response,
  {}
>;

function withAuth<TInputProps = {}>() {
  return graphql<TInputProps, Response, {}, AuthChildProps<TInputProps>>(
    ME_QUERY
  );
}

export default withAuth;
