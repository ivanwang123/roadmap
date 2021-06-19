import { gql } from "@apollo/client";
import { USER_INFO_FIELDS } from "../fragments/user";

export const REGISTER_MUTATION = gql`
  ${USER_INFO_FIELDS}
  mutation Register($email: String!, $username: String!, $password: String!) {
    createUser(
      input: { email: $email, username: $username, password: $password }
    ) {
      ...UserInfoFields
    }
  }
`;
