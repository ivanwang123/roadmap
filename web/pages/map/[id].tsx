import { Layout } from "components/layout";
import { Loading } from "components/placeholder";
import { addApolloState, getApolloClient } from "lib/apollo-client";
import { withAuth, WithAuthProps } from "modules/auth";
import { Feed, ROADMAP_QUERY, Sidebar, useRoadmap } from "modules/roadmap";
import { GetServerSideProps } from "next";
import { useRouter } from "next/router";
import React from "react";
import { RoadmapQuery, RoadmapQueryVariables } from "types/graphql-generated";

function Roadmap({ data: { me } }: WithAuthProps) {
  const router = useRouter();

  const { roadmap, loading, error } = useRoadmap({
    id: parseInt(router.query.id as string),
  });

  if (loading) return <Loading />;
  if (error) return <h1>Error</h1>;

  return (
    <Layout title="Map | Roadmap">
      <main className="sidebar-grid h-full max-h-full overflow-hidden">
        <Sidebar roadmap={roadmap!} me={me} />
        <Feed roadmap={roadmap!} me={me} />
      </main>
    </Layout>
  );
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const id = parseInt(ctx.params?.id as string);

  const client = getApolloClient();

  try {
    await client.query<RoadmapQuery, RoadmapQueryVariables>({
      query: ROADMAP_QUERY,
      variables: { id: id },
      context: {
        headers: {
          Cookie: ctx.req.headers.cookie,
        },
      },
    });
  } catch (err) {
    console.error(err);
  }

  return addApolloState(client);
};

export default withAuth(Roadmap);
