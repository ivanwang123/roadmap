import { gql } from "@apollo/client";

export const ROADMAPS_QUERY = gql`
  query Roadmaps($cursorId: Int!, $cursorValue: String!, $sort: Sort!) {
    roadmaps(
      input: { cursorId: $cursorId, cursorValue: $cursorValue, sort: $sort }
    ) {
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
