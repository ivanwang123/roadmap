import { graphql, ChildDataProps } from "@apollo/react-hoc";
import { ME_QUERY } from "../graphql/queries/me";

type Response = {
  me: {
    id: number;
    username: string;
    email: string;
  };
};

export type ChildProps<TInputProps> = ChildDataProps<TInputProps, Response, {}>;

function withAuth<TInputProps>() {
  return graphql<TInputProps, Response, {}, ChildProps<TInputProps>>(ME_QUERY);
}

export default withAuth;
