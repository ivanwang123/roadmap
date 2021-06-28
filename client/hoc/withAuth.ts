import { graphql, ChildDataProps } from "@apollo/react-hoc";
import { ME_QUERY } from "../graphql/queries/me";

type Response = {
  me: {
    id: number;
    username: string;
    email: string;
  };
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
