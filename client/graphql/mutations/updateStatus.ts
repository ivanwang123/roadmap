import { gql } from "@apollo/client";

export const UPDATE_STATUS_MUTATION = gql`
  mutation UpdateCheckpointStatus($checkpointId: Int!, $status: String!) {
    updateCheckpointStatus(
      input: { checkpointId: $checkpointId, status: $status }
    ) {
      id
      status
    }
  }
`;
