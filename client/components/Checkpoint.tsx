import { useMutation } from "@apollo/client";
import React from "react";
import { UPDATE_STATUS_MUTATION } from "../graphql/mutations/updateStatus";
import Check from "../svgs/check.svg";
import DashedArrow from "../svgs/dashed-arrow.svg";
import Skip from "../svgs/skip.svg";
import { CheckpointType, StatusType } from "../types/checkpointTypes";

type Props = {
  idx: number;
  checkpoint: CheckpointType;
  // reference: React.LegacyRef<HTMLDivElement>;
};

function Checkpoint({ idx, checkpoint }: Props) {
  const [updateStatus] = useMutation(UPDATE_STATUS_MUTATION);

  const handleUpdateStatus = (status: StatusType) => {
    const newStatus = checkpoint.status === status ? "incomplete" : status;
    updateStatus({
      variables: {
        checkpointId: checkpoint.id,
        status: newStatus,
      },
      optimisticResponse: {
        updateCheckpointStatus: {
          id: checkpoint.id,
          __typename: "Checkpoint",
          status: newStatus,
        },
      },
    }).catch((e) => console.error(e));
  };

  return (
    <>
      <div className="flex flex-col">
        <div className="flex items-center">
          <div className="grid place-items-center w-10 h-10 bg-gray-800 text-white font-bold rounded-xl">
            {idx}
          </div>
          <h3 className="text-2xl text-gray-800 font-bold tracking-wide ml-4">
            {checkpoint.title}
          </h3>
        </div>
        <div className="flex flex-col">
          <article className="w-full bg-white py-4 rounded">
            <p className="text-gray-500 tracking-wide leading-7">
              {checkpoint.instructions}
            </p>
            <h6 className="text-gray-700 font-bold mt-4">RESOURCES</h6>
            <div className="flex flex-wrap mt-2">
              {checkpoint.links.map((link) => {
                return (
                  <a
                    className="mr-3 mb-3"
                    href={link.url}
                    target="_blank"
                    rel="noopener noreferrer"
                  >
                    <div className="flex h-20 text-gray-500 border rounded overflow-hidden">
                      <img
                        className="w-20 h-20 border-r object-cover overflow-hidden"
                        src={link.image}
                        alt={link.title}
                      />
                      <div className="w-52 px-3 py-2">
                        <p className="text-gray-700 font-semibold truncate">
                          {link.title}
                        </p>
                        <p className="text-gray-500 text-sm line-clamp-2">
                          {link.description || ""}
                        </p>
                      </div>
                    </div>
                  </a>
                );
              })}
            </div>
          </article>
          <div className="flex ml-aut">
            <button
              type="button"
              className={`w-16 h-10 text-emerald-500 border-b-2 p-1 cursor-pointer hover:bg-emerald-100 hover:border-emerald-300 focus:outline-none ${
                checkpoint.status === "complete"
                  ? "border-b-4 bg-emerald-100 border-emerald-300 shadow"
                  : ""
              }`}
              onClick={() => handleUpdateStatus("complete")}
            >
              <Check className="fill-current" width="100%" height="100%" />
            </button>
            <button
              className={`w-16 h-10 text-yellow-500 border-b-2 p-1 cursor-pointer hover:bg-yellow-100 hover:border-yellow-300 focus:outline-none ${
                checkpoint.status === "skip"
                  ? "border-b-4 bg-yellow-100 border-yellow-300 shadow"
                  : ""
              }`}
              type="button"
              onClick={() => handleUpdateStatus("skip")}
            >
              <Skip className="fill-current" width="100%" height="100%" />
            </button>
          </div>
        </div>
      </div>

      <div>
        <DashedArrow
          className="fill-current text-gray-700 transform ml-aut"
          width={150}
          height={150}
        />
      </div>
    </>
  );
}

export default Checkpoint;
