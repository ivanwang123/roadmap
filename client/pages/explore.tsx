import { useQuery } from "@apollo/client";
import React from "react";
import Layout from "../components/Layout";
import Loading from "../components/Loading";
import RoadmapCard from "../components/RoadmapCard";
import { ROADMAPS_QUERY } from "../graphql/queries/roadmaps";

function Explore() {
  const { data, loading, error } = useQuery(ROADMAPS_QUERY, {
    variables: {
      sort: "NEWEST",
    },
  });

  if (loading) return <Loading />;
  if (error) return <h1>Error</h1>;
  console.log("DATA", data);

  // TODO: Pagination
  return (
    <Layout title="Roadmap">
      <main className="grid grid-cols-12">
        <section className="col-start-4 col-end-11">
          <h1 className="text-5xl text-gray-800 font-bold tracking-wider mt-8 mb-3">
            Explore
          </h1>
          {data.roadmaps.map(() => (
            <RoadmapCard />
          ))}
        </section>
      </main>
    </Layout>
  );
}

export default Explore;
