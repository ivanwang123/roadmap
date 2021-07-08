import { gql } from "@apollo/client";
import { ROADMAP_INFO_FIELDS } from "../fragments/roadmap";
import { USER_INFO_FIELDS } from "../fragments/user";

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
