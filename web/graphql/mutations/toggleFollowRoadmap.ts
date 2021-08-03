import { gql } from "@apollo/client";

export const TOGGLE_FOLLOW_ROADMAP_MUTATION = gql`
  mutation ToggleFollowRoadmap($roadmapId: Int!) {
    toggleFollowRoadmap(input: { roadmapId: $roadmapId }) {
      id
      followers {
        id
      }
    }
  }
`;
