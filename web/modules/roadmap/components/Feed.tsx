import dayjs from "dayjs";
import { Checkpoint } from "modules/checkpoint";
import Link from "next/link";
import React from "react";
import { RoadmapQuery, UserInfoFieldsFragment } from "types/graphql-generated";

type Props = {
  roadmap: RoadmapQuery["roadmap"];
  me?: UserInfoFieldsFragment;
};

export function Feed({ roadmap, me }: Props) {
  return (
    <section className="px-10 overflow-auto scroll-smooth">
      <div className="max-w-4xl">
        {/* TODO: Singular/plural */}
        <div className="mt-10 mb-8">
          <div className="flex text-gray-500 tracking-wide mb-6">
            <Link href="/user/1">
              <a className="flex items-center mr-2 transition duration-300 hover:text-hover">
                <span className="w-5 h-5 bg-blue-200 rounded-full mr-2"></span>
                {roadmap!.creator.username}
              </a>
            </Link>
            |
            <span className="ml-2">
              {dayjs(roadmap!.createdAt).format("MMMM D, YYYY")}
            </span>
          </div>
          <h1 className="text-3xl text-black font-semibold mb-2">
            {/* {data.roadmap.title} */}
            Visual Elements of User Interface
          </h1>
          <div className="text-gray-500 tracking-wide mb-4">
            {/* {data.roadmap.description} */}A guide to learn everything
            fullstack.
          </div>
          <div className="flex flex-wrap mb-6">
            <span className="text-gray-400 text-sm font-medium bg-secondary px-2 py-1 mr-2 rounded-sm">
              Wedev
            </span>
            <span className="text-gray-400 text-sm font-medium bg-secondary px-2 py-1 mr-2 rounded-sm">
              Fullstack
            </span>
            <span className="text-gray-400 text-sm font-medium bg-secondary px-2 py-1 rounded-sm">
              Javascript
            </span>
          </div>
        </div>

        {/* CONTENT */}
        <div className="mb-32">
          {roadmap!.checkpoints.map((checkpoint, idx) => (
            <Checkpoint checkpoint={checkpoint} isAuth={!!me} key={idx} />
          ))}
        </div>
      </div>
    </section>
  );
}
