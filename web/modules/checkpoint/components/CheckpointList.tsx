import React, { useCallback } from "react";
import { RoadmapQuery, Status } from "types/graphql-generated";
import { CheckpointStatus } from "./CheckpointStatus";

type Props = {
  checkpoints: RoadmapQuery["roadmap"]["checkpoints"];
};

export function CheckpointList({ checkpoints }: Props) {
  const numCompletedCheckpoints = useCallback(() => {
    return checkpoints.filter((c) => c.status && c.status === Status.Complete)
      .length;
  }, [checkpoints]);

  return (
    <>
      <h6 className="text-gray-300 text-sm font-semibold tracking-wide mb-4">
        CHECKPOINTS {`(${numCompletedCheckpoints()}/${checkpoints.length})`}
      </h6>
      <div className="checkpoints-grid gap-x-1 gap-y-3 items-center">
        {checkpoints.map((checkpoint, idx) => (
          <CheckpointStatus
            id={checkpoint.id}
            title={checkpoint.title}
            status={checkpoint.status}
            key={idx}
          />
        ))}
      </div>
    </>
  );
}
