import { gql } from "@apollo/client";

export const USER_INFO_FIELDS = gql`
  fragment UserInfoFields on User {
    id
    username
    email
    createdAt
    updatedAt
  }
`;
