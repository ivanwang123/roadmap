import { useQuery } from "@apollo/client";
import React, { useState } from "react";
import Layout from "../components/Layout";
import Loading from "../components/Loading";
import RoadmapCard from "../components/RoadmapCard";
import { ROADMAPS_QUERY } from "../graphql/queries/roadmaps";
import { SortType } from "../types/roadmapTypes";

function Explore() {
  const [sort, setSort] = useState<SortType>(SortType.MOST_FOLLOWERS);

  const { data, loading, error } = useQuery(ROADMAPS_QUERY, {
    variables: {
      cursorId: 1,
      cursorValue: "1",
      sort: sort,
    },
  });
  console.log(data);

  if (loading) return <Loading />;
  if (error) return <h1>Error</h1>;

  // TODO: Pagination
  return (
    <Layout title="Roadmap">
      <main className="grid grid-cols-12">
        <section className="col-start-4 col-end-11">
          <h1 className="text-3xl text-gray-800 font-medium mt-14 mb-6">
            Explore
          </h1>
          <div className="flex text-gray-400 text-sm font-light tracking-wide">
            <div
              className={`w-16 text-center mr-4 cursor-pointer hover:text-gray-800 ${
                sort === SortType.MOST_FOLLOWERS
                  ? "text-gray-800 font-medium"
                  : ""
              }`}
              onClick={() => setSort(SortType.MOST_FOLLOWERS)}
            >
              Popular
            </div>
            <div
              className={`w-16 text-center mr-4 cursor-pointer hover:text-gray-800 ${
                sort === SortType.NEWEST ? "text-gray-800 font-medium" : ""
              }`}
              onClick={() => setSort(SortType.NEWEST)}
            >
              New
            </div>
            <div
              className={`w-28 text-center mr-4 cursor-pointer hover:text-gray-800 ${
                sort === SortType.MOST_CHECKPOINTS
                  ? "text-gray-800 font-medium"
                  : ""
              }`}
              onClick={() => setSort(SortType.MOST_CHECKPOINTS)}
            >
              ^ Checkpoints
            </div>
            <div className="w-28 text-center mr-4 cursor-pointer hover:text-gray-800">
              ^ Resources
            </div>
          </div>
          {data.roadmaps.map(() => (
            <RoadmapCard />
          ))}
        </section>
      </main>
    </Layout>
  );
}

export default Explore;
