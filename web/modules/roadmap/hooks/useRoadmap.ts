import { useQuery } from "@apollo/client";
import { useNotificationStore } from "modules/notification";
import { RoadmapQuery, RoadmapQueryVariables } from "types/graphql-generated";
import { ROADMAP_QUERY } from "../api";

type Props = {
  id: number;
};

export function useRoadmap({ id }: Props) {
  const { setNotification } = useNotificationStore();
  const { data, loading, error } = useQuery<
    RoadmapQuery,
    RoadmapQueryVariables
  >(ROADMAP_QUERY, {
    variables: { id },
  });

  if (error) {
    setNotification({
      type: "error",
      message: "Unable to get roadmap",
    });
  }

  return { roadmap: data?.roadmap, loading, error };
}
