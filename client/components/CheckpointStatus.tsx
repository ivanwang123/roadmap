import React from "react";
import CheckCircle from "../svgs/check-circle.svg";
import Circle from "../svgs/circle.svg";
import SkipCircle from "../svgs/skip-circle.svg";
import { StatusType } from "../types/checkpointTypes";

type Props = {
  checkpoint: string;
  status?: StatusType;
  isLast?: boolean;
};

function CheckpointStatus({ checkpoint, status, isLast = false }: Props) {
  let icon = null;
  let barColor = "border-gray-200";

  switch (status) {
    case "complete":
      icon = (
        <CheckCircle
          className="col-start-1 fill-current text-emerald-400 mx-auto z-10"
          width={24}
          height={24}
        />
      );
      barColor = "border-emerald-400";
      break;
    case "skip":
      icon = (
        <SkipCircle
          className="col-start-1 fill-current text-yellow-400 mx-auto z-10"
          width={24}
          height={24}
        />
      );
      barColor = "border-yellow-400";
      break;
    case "current":
      icon = (
        <Circle
          className="col-start-1 fill-current text-gray-700 mx-auto z-10"
          width={12}
          height={12}
        />
      );
      break;
    default:
      icon = (
        <Circle
          className="col-start-1 fill-current text-gray-300 mx-auto z-10"
          width={12}
          height={12}
        />
      );
      break;
  }

  return (
    <>
      {icon}
      <p
        className={`flex items-center text-gray-400 tracking-wide ${
          status === "current" ? "font-bold" : ""
        }`}
      >
        {checkpoint}
      </p>
      {!isLast && (
        <div
          className={`w-1/2 h-8 justify-self-end border-l-2 -my-2 mr-xs ${barColor}`}
        ></div>
      )}
    </>
  );
}

export default CheckpointStatus;
