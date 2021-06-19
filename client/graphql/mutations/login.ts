import { gql } from "@apollo/client";
import { USER_INFO_FIELDS } from "../fragments/user";

export const LOGIN_MUTATION = gql`
  ${USER_INFO_FIELDS}
  mutation Login($email: String!, $password: String!) {
    login(input: { email: $email, password: $password }) {
      ...UserInfoFields
    }
  }
`;
