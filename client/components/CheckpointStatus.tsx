import React from "react";
import { Status } from "../graphql/generated/generated";
// import CheckCircle from "../svgs/check-circle.svg";
import Check from "../svgs/check.svg";
import Circle from "../svgs/circle.svg";
// import SkipCircle from "../svgs/skip-circle.svg";
import Skip from "../svgs/skip.svg";

type Props = {
  id: number;
  title: string;
  status?: Status | null;
};

function CheckpointStatus({ id, title, status }: Props) {
  let icon = null;

  switch (status) {
    case Status.Complete:
      icon = (
        <span className="col-start-1 w-6 h-6 p-1 bg-emerald-100 mx-auto rounded-full z-10">
          <Check
            className="fill-current text-emerald-500"
            width={16}
            height={16}
          />
        </span>
      );
      break;
    case Status.Skip:
      icon = (
        <span className="col-start-1 w-6 h-6 p-1 bg-yellow-100 mx-auto rounded-full z-10">
          <Skip
            className="fill-current text-yellow-500"
            width={16}
            height={16}
          />
        </span>
      );
      break;
    default:
      icon = (
        <span className="grid place-items-center w-6 h-6 mx-auto">
          <Circle
            className="col-start-1 fill-current text-gray-300 mx-auto z-10"
            width={10}
            height={10}
          />
        </span>
      );
      break;
  }

  return (
    <>
      {icon}
      <a
        href={"#" + title + " " + id}
        className="flex items-center text-gray-400 text-sm font-light tracking-wide"
      >
        {title}
      </a>
      {/* {!isLast && (
        <div
          className={`w-px h-10 justify-self-center -my-2 ${barColor}`}
        ></div>
      )} */}
    </>
  );
}

export default CheckpointStatus;
