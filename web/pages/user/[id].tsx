import { useQuery } from "@apollo/client";
import { Layout } from "components/layout";
import { Loading } from "components/placeholder";
import { addApolloState, getApolloClient } from "lib/apollo-client";
import { RoadmapCard } from "modules/roadmap";
import { USER_QUERY } from "modules/user";
import { GetServerSideProps } from "next";
import { useRouter } from "next/router";
import React from "react";
import { UserQuery, UserQueryVariables } from "types/graphql-generated";

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
        <section className="sticky top-0 flex flex-col h-full border-r-2 border-secondary px-10 pt-10 overflow-auto">
          <div className="w-40 h-40 bg-blue-100 rounded-full"></div>
          <div className="text-2xl text-gray-800 font-medium mt-4">
            {data!.user.username}
          </div>
          <div className="flex text-gray-400 text-sm font-light tracking-wide mt-4">
            <p className="mr-1">
              <span className="font-medium">
                {data!.user.createdRoadmaps.length}
              </span>{" "}
              created
            </p>
            ·
            <p className="ml-1">
              <span className="font-medium">
                {data!.user.followingRoadmaps.length}
              </span>{" "}
              following
            </p>
          </div>
        </section>
        <section className="px-10 pt-14 overflow-auto scroll-smooth">
          <div className="flex text-gray-400 text-sm font-light tracking-wide">
            <p className="mr-8 hover:text-gray-800">
              Created ({data!.user.createdRoadmaps.length})
            </p>
            <p className="hover:text-gray-800">
              Following ({data!.user.followingRoadmaps.length})
            </p>
          </div>
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
