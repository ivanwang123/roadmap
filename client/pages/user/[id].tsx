import { useQuery } from "@apollo/client";
import { GetServerSideProps } from "next";
import { useRouter } from "next/router";
import React from "react";
import Layout from "../../components/Layout";
import Loading from "../../components/Loading";
import RoadmapCard from "../../components/RoadmapCard";
import {
  UserQuery,
  UserQueryVariables,
} from "../../graphql/generated/generated";
import { USER_QUERY } from "../../graphql/queries/userById";
import { addApolloState, getApolloClient } from "../../lib/apollo-client";

function UserProfile() {
  const router = useRouter();

  const { data, loading, error } = useQuery<UserQuery, UserQueryVariables>(
    USER_QUERY,
    {
      variables: { id: parseInt(router.query.id as string) },
      fetchPolicy: "cache-only",
    }
  );

  if (loading) return <Loading />;
  if (error) return <h1>Error</h1>;
  console.log(data);

  return (
    <Layout title="User | Roadmap">
      <main className="sidebar-grid h-full max-h-full overflow-hidden bg-white">
        <section className="sticky top-0 flex flex-col h-full border-r-2 border-secondary pl-10 pt-10 overflow-auto">
          <div className="w-40 h-40 bg-blue-100 rounded-full"></div>
          <div>{data!.user.username}</div>
          <div>Created {data!.user.createdRoadmaps.length} roadmaps</div>
          <div>Following {data!.user.followingRoadmaps.length} roadmaps</div>
        </section>
        <section className="px-10 overflow-auto scroll-smooth">
          <div className="max-w-4xl">
            {data!.user.createdRoadmaps.map((roadmap, idx) => (
              <RoadmapCard roadmap={roadmap} key={idx} />
            ))}
          </div>
        </section>
      </main>
    </Layout>
  );
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const id = parseInt(ctx.params?.id as string);

  const client = getApolloClient();

  try {
    await client.query<UserQuery, UserQueryVariables>({
      query: USER_QUERY,
      variables: { id: id },
      context: {
        headers: {
          Cookie: ctx.req.headers.cookie,
        },
      },
    });
  } catch (e) {
    console.error(e);
  }

  return addApolloState(client, {
    props: {},
  });
};

export default UserProfile;
