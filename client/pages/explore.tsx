import { useQuery } from "@apollo/client";
import React, { useState } from "react";
import Layout from "../components/Layout";
import Loading from "../components/Loading";
import RoadmapCard from "../components/RoadmapCard";
import {
  RoadmapsQuery,
  RoadmapsQueryVariables,
  Sort,
} from "../graphql/generated/generated";
import { ROADMAPS_QUERY } from "../graphql/queries/roadmaps";

function Explore() {
  const [sort, setSort] = useState<Sort>(Sort.MostFollowers);

  const { data, loading, error } = useQuery<
    RoadmapsQuery,
    RoadmapsQueryVariables
  >(ROADMAPS_QUERY, {
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
    <Layout title="Explore | Roadmap">
      <main className="grid grid-cols-12">
        <section className="col-start-4 col-end-11">
          <h1 className="text-3xl text-gray-800 font-medium mt-14 mb-6">
            Explore
          </h1>
          <div className="flex text-gray-400 text-sm font-light tracking-wide">
            <div
              className={`w-16 text-center mr-4 cursor-pointer hover:text-gray-800 ${
                sort === Sort.MostFollowers ? "text-gray-800 font-medium" : ""
              }`}
              onClick={() => setSort(Sort.MostFollowers)}
            >
              Popular
            </div>
            <div
              className={`w-16 text-center mr-4 cursor-pointer hover:text-gray-800 ${
                sort === Sort.Newest ? "text-gray-800 font-medium" : ""
              }`}
              onClick={() => setSort(Sort.Newest)}
            >
              New
            </div>
            <div
              className={`w-28 text-center mr-4 cursor-pointer hover:text-gray-800 ${
                sort === Sort.MostCheckpoints ? "text-gray-800 font-medium" : ""
              }`}
              onClick={() => setSort(Sort.MostCheckpoints)}
            >
              ^ Checkpoints
            </div>
            <div className="w-28 text-center mr-4 cursor-pointer hover:text-gray-800">
              ^ Resources
            </div>
          </div>
          {data?.roadmaps.map((roadmap) => (
            <RoadmapCard roadmap={roadmap} />
          ))}
        </section>
      </main>
    </Layout>
  );
}

export default Explore;
