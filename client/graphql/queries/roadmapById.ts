import { gql } from "@apollo/client";
import { LINK_FIELDS } from "../fragments/link";

export const ROADMAP_QUERY = gql`
  ${LINK_FIELDS}
  query Roadmap($id: Int!) {
    roadmap(input: { id: $id }) {
      id
      title
      description
      creator {
        id
        username
      }
      checkpoints {
        id
        title
        instructions
        status
        links {
          ...LinkFields
        }
      }
      followers {
        id
      }
      createdAt
      updatedAt
    }
  }
`;
