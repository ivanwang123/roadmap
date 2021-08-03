import { gql } from "@apollo/client";

export const LINK_FIELDS = gql`
  fragment LinkFields on Link {
    url
    title
    description
    image
  }
`;
