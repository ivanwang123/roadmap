import { gql } from "@apollo/client";

export const ROADMAPS_QUERY = gql`
  # enum SortType {
  #   NEWEST
  #   OLDEST
  #   MOST_FOLLOWERS
  #   MOST_CHECKPOINTS
  #   LEAST_CHECKPOINTS
  # }

  query Roadmaps($sort: Sort!) {
    roadmaps(input: { cursor: "2021-07-01 17:33:22.283204-04", sort: $sort }) {
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
