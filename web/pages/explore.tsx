import { useQuery } from "@apollo/client";
import clsx from "clsx";
import { Layout } from "components/layout";
import { Loading } from "components/placeholder";
import { RoadmapCard, ROADMAPS_QUERY } from "modules/roadmap";
import React, { useState } from "react";
import {
  RoadmapsQuery,
  RoadmapsQueryVariables,
  Sort,
} from "types/graphql-generated";

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
        <section className="col-start-4 col-end-10">
          <h1 className="text-3xl text-black font-semibold mt-14 mb-6">
            Explore
          </h1>
          <div className="flex text-gray-400 text-sm tracking-wide border-b">
            <div
              className={clsx(
                "w-16 text-center border-b-3 border-white pb-2 mr-4 cursor-pointer hover:text-black",
                sort === Sort.MostFollowers && "text-black border-blue-500"
              )}
              onClick={() => setSort(Sort.MostFollowers)}
            >
              Popular
            </div>
            <div
              className={clsx(
                "w-16 text-center border-b-3 border-white pb-2 mr-4 cursor-pointer hover:text-black",
                sort === Sort.Newest && "text-black border-blue-500"
              )}
              onClick={() => setSort(Sort.Newest)}
            >
              New
            </div>
            <div
              className={clsx(
                "w-28 text-center border-b-3 border-white pb-2 mr-4 cursor-pointer hover:text-black",
                sort === Sort.MostCheckpoints && "text-black border-blue-500"
              )}
              onClick={() => setSort(Sort.MostCheckpoints)}
            >
              ^ Checkpoints
            </div>
            <div className="w-28 text-center mr-4 cursor-pointer hover:text-black">
              ^ Resources
            </div>
          </div>
          {data?.roadmaps.map((roadmap) => (
            <RoadmapCard roadmap={roadmap} key={roadmap.id} />
          ))}
        </section>
      </main>
    </Layout>
  );
}

export default Explore;
