import { useMutation } from "@apollo/client";
import clsx from "clsx";
import { UPDATE_STATUS_MUTATION } from "modules/roadmap";
import React from "react";
import Check from "svgs/check.svg";
import { LinkFieldsFragment, Maybe, Status } from "types/graphql-generated";
import { generateCheckpointId } from "../utils";

type CheckpointType = {
  id: number;
  title: string;
  instructions: string;
  status?: Maybe<Status> | undefined;
  links: LinkFieldsFragment[];
};

type Props = {
  checkpoint: CheckpointType;
  isAuth: boolean;
};

export function Checkpoint({ checkpoint, isAuth }: Props) {
  const [updateStatus] = useMutation(UPDATE_STATUS_MUTATION);

  const handleUpdateStatus = (status: Status) => {
    if (isAuth) {
      console.log("STATUESS", checkpoint.status, status);
      const newStatus =
        checkpoint.status === status ? Status.Incomplete : status;
      console.log("CHECKPOINT ID", checkpoint.id, newStatus);
      updateStatus({
        variables: {
          checkpointId: checkpoint.id,
          status: newStatus,
        },
        optimisticResponse: {
          updateCheckpointStatus: {
            __typename: "Checkpoint",
            id: checkpoint.id,
            status: newStatus,
          },
        },
      }).catch((e) => console.error(e));
    }
  };

  return (
    <>
      <div
        className="flex flex-col py-8"
        id={generateCheckpointId(checkpoint.title, checkpoint.id)}
      >
        <div className="flex items-center">
          {isAuth && (
            <div className="flex mr-2">
              <button
                type="button"
                className={clsx(
                  "w-7 h-7 text-gray-300 rounded-full cursor-pointer focus:outline-none transition duration-300",
                  checkpoint.status === Status.Complete
                    ? "text-emerald-600 animate__animated animate__heartBeat"
                    : "hover:text-gray-400 hover:bg-gray-100"
                )}
                onClick={() => handleUpdateStatus(Status.Complete)}
              >
                <Check className="fill-current" width="100%" height="100%" />
              </button>

              {/* <button
              className={`w-8 h-8 text-yellow-500 cursor-pointer hover:bg-yellow-100 focus:outline-none ${
                checkpoint.status === Status.Skip ? "bg-yellow-100 shadow" : ""
              }`}
              type="button"
              onClick={() => handleUpdateStatus(Status.Skip)}
            >
              <Skip className="fill-current" width="100%" height="100%" />
            </button> */}
            </div>
          )}
          <h3 className="text-xl text-black font-semibold">
            {/* <span className="text-gray-300 mr-4">
              {idx.toString().padStart(2, "0")}
            </span> */}
            {checkpoint.title}
          </h3>
        </div>
        <div className="flex flex-col">
          <article className="mt-3">
            <p className="text-gray-500 tracking-wide leading-8">
              {/* {checkpoint.instructions} */}
              HTML is the backbone of a website. These resources go over the
              basics and when you finish reading through the tutorials, take the
              quiz on w3schools before moving on.
            </p>
            <div className="flex flex-col w-full mt-6">
              {checkpoint.links.map((link, idx) => (
                <div className="flex items-center w-max mb-3" key={idx}>
                  <a
                    href={link.url}
                    target="_blank"
                    rel="noopener noreferrer"
                    key={idx}
                  >
                    <div className="flex w-full h-24 rounded-sm overflow-hidden transition duration-300 hover:bg-secondary">
                      <img
                        className="w-32 h-24 text-xs text-gray-400 tracking-wide border object-cover rounded-sm overflow-hidden"
                        src={link.image}
                        alt={link.title}
                      />
                      <div className="flex flex-col w-96 tracking-wide px-3 py-1">
                        <p className="font-semibold truncate">{link.title}</p>
                        <p className="text-gray-500 text-sm line-clamp-2">
                          {link.description || ""}
                        </p>
                        <p className="text-sm mt-auto truncate">{link.url}</p>
                      </div>
                    </div>
                  </a>
                </div>
              ))}
            </div>
          </article>
        </div>
      </div>
    </>
  );
}
