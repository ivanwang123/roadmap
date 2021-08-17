import { gql } from "@apollo/client";
import { ROADMAP_INFO_FIELDS } from "modules/roadmap";

export const USER_INFO_FIELDS = gql`
  fragment UserInfoFields on User {
    id
    username
    email
    createdAt
    updatedAt
  }
`;

export const USER_QUERY = gql`
  ${USER_INFO_FIELDS}
  ${ROADMAP_INFO_FIELDS}
  query User($id: Int!) {
    user(input: { id: $id }) {
      ...UserInfoFields
      followingRoadmaps {
        ...RoadmapInfoFields
      }
      createdRoadmaps {
        ...RoadmapInfoFields
      }
    }
  }
`;
