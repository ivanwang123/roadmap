import { gql } from "@apollo/client";
import { USER_INFO_FIELDS } from "../fragments/user";

export const ME_QUERY = gql`
  ${USER_INFO_FIELDS}
  query Me {
    me {
      ...UserInfoFields
    }
  }
`;
