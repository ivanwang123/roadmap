import { useMutation } from "@apollo/client";
import React from "react";
import { UPDATE_STATUS_MUTATION } from "../graphql/mutations/updateStatus";
import Check from "../svgs/check.svg";
// import DashedArrow from "../svgs/dashed-arrow.svg";
import Skip from "../svgs/skip.svg";
import {
  LinkFieldsFragment,
  Maybe,
  Status,
} from "../graphql/generated/generated";

type CheckpointType = {
  id: number;
  title: string;
  instructions: string;
  status?: Maybe<Status> | undefined;
  links: LinkFieldsFragment[];
};

type Props = {
  idx: number;
  checkpoint: CheckpointType;
  isAuth: boolean;
};

function Checkpoint({ idx, checkpoint, isAuth }: Props) {
  const [updateStatus] = useMutation(UPDATE_STATUS_MUTATION);

  const handleUpdateStatus = (status: Status) => {
    if (isAuth) {
      const newStatus = checkpoint.status === status ? "incomplete" : status;
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
        id={checkpoint.title + " " + checkpoint.id}
      >
        <div className="flex items-center">
          <div className="text-3xl text-gray-300 font-medium">
            {idx.toString().padStart(2, "0")}
          </div>
          <h3 className="text-xl text-gray-800 font-medium ml-4">
            {checkpoint.title}
          </h3>
        </div>
        <div className="flex flex-col">
          <article className="w-full pt-5 pb-2 rounded">
            <p className="text-gray-500 tracking-wide font-light leading-7">
              {/* {checkpoint.instructions} */}
              HTML is the backbone of a website. These resources go over the
              basics and when you finish reading through the tutorials, take the
              quiz on w3schools before moving on.
            </p>
            {/* <h6 className="text-gray-300 tracking-wide font-bold mt-4">
              RESOURCES
            </h6> */}
            <div className="flex flex-col w-full mt-8">
              {checkpoint.links.map((link) => {
                return (
                  <a
                    className="flex items-center w-max mb-3"
                    href={link.url}
                    target="_blank"
                    rel="noopener noreferrer"
                  >
                    <div className="flex w-full h-24 text-gray-500 rounded overflow-hidden transition duration-300 hover:bg-secondary">
                      <img
                        className="w-32 h-24 text-xs text-gray-400 tracking-wide border object-cover rounded overflow-hidden"
                        src={link.image}
                        alt={link.title}
                      />
                      <div className="flex flex-col w-96 tracking-wide px-3 py-2">
                        <p className="text-gray-700 text-sm font-medium truncate">
                          {link.title}
                        </p>
                        <p className="text-gray-400 text-xs font-light line-clamp-2">
                          {link.description || ""}
                        </p>
                        <p className="text-gray-700 text-xs font-light mt-auto truncate">
                          {link.url}
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
                checkpoint.status === Status.Complete
                  ? "border-b-4 bg-emerald-100 border-emerald-300 shadow"
                  : ""
              }`}
              onClick={() => handleUpdateStatus(Status.Complete)}
            >
              <Check className="fill-current" width="100%" height="100%" />
            </button>
            <button
              className={`w-16 h-10 text-yellow-500 border-b-2 p-1 cursor-pointer hover:bg-yellow-100 hover:border-yellow-300 focus:outline-none ${
                checkpoint.status === Status.Skip
                  ? "border-b-4 bg-yellow-100 border-yellow-300 shadow"
                  : ""
              }`}
              type="button"
              onClick={() => handleUpdateStatus(Status.Skip)}
            >
              <Skip className="fill-current" width="100%" height="100%" />
            </button>
          </div>
        </div>
      </div>
    </>

    // <>
    //   <div
    //     className="flex flex-col py-10"
    //     id={checkpoint.title + " " + checkpoint.id}
    //   >
    //     <div className="flex items-center">
    //       <div className="text-3xl text-gray-300 font-medium">
    //         {idx.toString().padStart(2, "0")}
    //       </div>
    //       <h3 className="text-xl text-gray-800 font-medium ml-4">
    //         {checkpoint.title}
    //       </h3>
    //     </div>
    //     <div className="flex flex-col">
    //       <article className="w-full bg-white pt-4 pb-2 rounded">
    //         <p className="text-gray-500 tracking-wide font-light leading-7">
    //           {/* {checkpoint.instructions} */}
    //           HTML is the backbone of a website. These resources go over the
    //           basics and when you finish reading through the tutorials, take the
    //           quiz on w3schools before moving on.
    //         </p>
    //         {/* <h6 className="text-gray-300 tracking-wide font-bold mt-4">
    //           RESOURCES
    //         </h6> */}
    //         <div className="flex flex-wrap mt-4">
    //           {checkpoint.links.map((link) => {
    //             return (
    //               <a
    //                 className="mr-3 mb-3"
    //                 href={link.url}
    //                 target="_blank"
    //                 rel="noopener noreferrer"
    //               >
    //                 <div className="flex h-20 text-gray-500 border rounded overflow-hidden transition duration-300 hover:bg-secondary">
    //                   <img
    //                     className="w-20 h-20 text-xs text-gray-400 tracking-wide border-r object-cover overflow-hidden"
    //                     src={link.image}
    //                     alt={link.title}
    //                   />
    //                   <div className="w-52 tracking-wide px-3 py-2">
    //                     <p className="text-gray-700 text-sm font-medium truncate">
    //                       {link.title}
    //                     </p>
    //                     <p className="text-gray-400 text-xs font-light line-clamp-2">
    //                       {link.description || ""}
    //                     </p>
    //                   </div>
    //                 </div>
    //               </a>
    //             );
    //           })}
    //         </div>
    //       </article>
    //       <div className="flex ml-aut">
    //         <button
    //           type="button"
    //           className={`w-16 h-10 text-emerald-500 border-b-2 p-1 cursor-pointer hover:bg-emerald-100 hover:border-emerald-300 focus:outline-none ${
    //             checkpoint.status === Status.Complete
    //               ? "border-b-4 bg-emerald-100 border-emerald-300 shadow"
    //               : ""
    //           }`}
    //           onClick={() => handleUpdateStatus(Status.Complete)}
    //         >
    //           <Check className="fill-current" width="100%" height="100%" />
    //         </button>
    //         <button
    //           className={`w-16 h-10 text-yellow-500 border-b-2 p-1 cursor-pointer hover:bg-yellow-100 hover:border-yellow-300 focus:outline-none ${
    //             checkpoint.status === Status.Skip
    //               ? "border-b-4 bg-yellow-100 border-yellow-300 shadow"
    //               : ""
    //           }`}
    //           type="button"
    //           onClick={() => handleUpdateStatus(Status.Skip)}
    //         >
    //           <Skip className="fill-current" width="100%" height="100%" />
    //         </button>
    //       </div>
    //     </div>
    //   </div>

    //   {/* TODO: Change color */}
    //   {/* <div>
    //     <DashedArrow
    //       className="fill-current text-gray-300"
    //       width={150}
    //       height={150}
    //     />
    //   </div> */}
    // </>
  );
}

export default Checkpoint;
