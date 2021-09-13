import React from "react";
import Check from "svgs/check.svg";
import Circle from "svgs/circle.svg";
import Skip from "svgs/skip.svg";
import { Status } from "types/graphql-generated";
import { generateCheckpointId } from "../utils";

type Props = {
  id: number;
  title: string;
  status?: Status | null;
};

export function CheckpointStatus({ id, title, status }: Props) {
  let icon = null;

  switch (status) {
    case Status.Complete:
      icon = (
        <Check
          className="w-6 h-6 fill-current text-emerald-600 mx-auto"
          width={16}
          height={16}
          data-testid="check-svg"
        />
      );
      break;
    case Status.Skip:
      icon = (
        <Skip
          className="w-6 h-6 fill-current text-yellow-500 mx-auto"
          width={16}
          height={16}
          data-testid="skip-svg"
        />
      );
      break;
    default:
      icon = (
        <span className="flex h-6 mx-auto">
          <Circle
            className="fill-current text-gray-200 my-auto"
            width={8}
            height={8}
            data-testid="circle-svg"
          />
        </span>
      );
      break;
  }

  return (
    <>
      {icon}
      <a
        href={`#${generateCheckpointId(title, id)}`}
        className="flex items-center w-max text-gray-400 tracking-wide"
      >
        {title}
      </a>
    </>
  );
}
