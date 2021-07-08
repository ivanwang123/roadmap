import { gql } from "@apollo/client";

export const ROADMAP_INFO_FIELDS = gql`
  fragment RoadmapInfoFields on Roadmap {
    id
    title
    description
    createdAt
    updatedAt
    creator {
      id
      username
    }
    followers {
      id
    }
    checkpoints {
      id
    }
  }
`;
