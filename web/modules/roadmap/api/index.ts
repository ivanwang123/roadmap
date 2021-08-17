import { gql } from "@apollo/client";
import { LINK_FIELDS } from "modules/checkpoint";

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

export const ROADMAPS_QUERY = gql`
  ${ROADMAP_INFO_FIELDS}
  query Roadmaps($cursorId: Int!, $cursorValue: String!, $sort: Sort!) {
    roadmaps(
      input: { cursorId: $cursorId, cursorValue: $cursorValue, sort: $sort }
    ) {
      ...RoadmapInfoFields
    }
  }
`;

export const ROADMAP_QUERY = gql`
  ${ROADMAP_INFO_FIELDS}
  ${LINK_FIELDS}
  query Roadmap($id: Int!) {
    roadmap(input: { id: $id }) {
      ...RoadmapInfoFields
      checkpoints {
        id
        title
        instructions
        status
        links {
          ...LinkFields
        }
      }
    }
  }
`;

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

export const UPDATE_STATUS_MUTATION = gql`
  mutation UpdateCheckpointStatus($checkpointId: Int!, $status: Status!) {
    updateCheckpointStatus(
      input: { checkpointId: $checkpointId, status: $status }
    ) {
      id
      status
    }
  }
`;
