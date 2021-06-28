import { gql } from "@apollo/client";

export const ROADMAPS_QUERY = gql`
  query Roadmaps {
    roadmaps {
      id
      title
      description
      creator {
        id
        username
      }
      followers {
        id
      }
      createdAt
    }
  }
`;
