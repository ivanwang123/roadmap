import { gql } from "@apollo/client";
import { ROADMAP_INFO_FIELDS } from "../fragments/roadmap";

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
