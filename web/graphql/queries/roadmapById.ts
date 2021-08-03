import { gql } from "@apollo/client";
import { LINK_FIELDS } from "../fragments/link";
import { ROADMAP_INFO_FIELDS } from "../fragments/roadmap";

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
