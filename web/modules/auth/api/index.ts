import { gql } from "@apollo/client";
import { USER_INFO_FIELDS } from "modules/user";

export const LOGIN_MUTATION = gql`
  ${USER_INFO_FIELDS}
  mutation Login($email: String!, $password: String!) {
    login(input: { email: $email, password: $password }) {
      ...UserInfoFields
    }
  }
`;

export const LOGOUT_MUTATION = gql`
  mutation Logout {
    logout
  }
`;

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

export const ME_QUERY = gql`
  ${USER_INFO_FIELDS}
  query Me {
    me {
      ...UserInfoFields
    }
  }
`;
