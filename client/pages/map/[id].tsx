import { useQuery } from "@apollo/client";
import dayjs from "dayjs";
import { GetServerSideProps } from "next";
import Link from "next/link";
import { useRouter } from "next/router";
import React from "react";
import Checkpoint from "../../components/Checkpoint";
import CheckpointStatus from "../../components/CheckpointStatus";
import FollowButton from "../../components/FollowButton";
import Icon from "../../components/Icon";
import Layout from "../../components/Layout";
import Loading from "../../components/Loading";
import {
  RoadmapQuery,
  RoadmapQueryVariables,
} from "../../graphql/generated/generated";
import { ROADMAP_QUERY } from "../../graphql/queries/roadmapById";
import withAuth, { AuthChildProps } from "../../hoc/withAuth";
import { addApolloState, getApolloClient } from "../../lib/apollo-client";
import Book from "../../svgs/book.svg";
import Flag from "../../svgs/flag-big.svg";

// function createObserver(el: Element | null) {
//   if (el) {
//     const options = {
//       root: null,
//       rootMargin: "-50% 0px 0px 0px",
//       threshold: 1.0,
//     };
//     const observer = new IntersectionObserver(
//       (observe) => console.log("OBSERVE", observe),
//       options
//     );
//     observer.observe(el);
//   }
// }

function Roadmap({ data: { me } }: AuthChildProps) {
  const router = useRouter();
  const { data, loading, error } = useQuery<
    RoadmapQuery,
    RoadmapQueryVariables
  >(ROADMAP_QUERY, {
    variables: { id: parseInt(router.query.id as string) },
  });

  if (loading) return <Loading />;
  if (error) return <h1>Error</h1>;
  // if (meError) return <h1>Me error</h1>;
  console.log(me, data?.roadmap);

  return (
    <Layout title="Map | Roadmap">
      <main className="sidebar-grid h-full max-h-full overflow-hidden">
        {/* SIDEBAR */}
        <section className="w-64 h-full">
          <div className="sticky top-0 flex flex-col h-full border-r-2 border-secondary pl-10 pt-10 overflow-auto">
            <div className="flex text-gray-400 text-sm font-light tracking-wide mb-2">
              <div className="flex items-center mr-10">
                <FollowButton
                  me={me}
                  followers={data!.roadmap.followers}
                  roadmapId={data!.roadmap.id}
                />
              </div>
            </div>
            <div className="text-gray-400 text-sm font-light tracking-wide mb-8">
              <div className="flex items-center mb-2">
                <Icon icon={Flag} size={12} />
                <span className="text-right font-medium ml-2 mr-1">
                  {data!.roadmap.checkpoints.length}
                </span>
                <span>checkpoints</span>
              </div>
              <div className="flex items-center">
                <Icon icon={Book} size={12} />
                <span className="text-right font-medium ml-2 mr-1">12</span>
                <span>resources</span>
              </div>
            </div>
            <h6 className="text-gray-400 font-bold tracking-wide mb-4">
              CHECKPOINTS
            </h6>
            <div className="checkpoints-grid gap-x-1 gap-y-3 items-center">
              {data!.roadmap.checkpoints.map((checkpoint, idx) => (
                <CheckpointStatus
                  id={checkpoint.id}
                  title={checkpoint.title}
                  status={checkpoint.status}
                  key={idx}
                />
              ))}
              {/* <CheckpointStatus checkpoint="HTML" status="skip" />
              <CheckpointStatus checkpoint="HTML" status="current" />
              <CheckpointStatus checkpoint="HTML" status="incomplete" /> */}
            </div>
          </div>
        </section>

        {/* MAIN */}
        <section className="px-10 overflow-auto scroll-smooth">
          <div className="max-w-4xl">
            {/* HEADER */}
            {/* TODO: Singular/plural */}
            <div className="mt-10 mb-8">
              <div className="flex text-gray-400 font-light tracking-wide text-sm mb-6">
                <Link href="/user/1">
                  <a className="flex items-center font-normal mr-2 transition duration-300 hover:text-hover">
                    <span className="w-5 h-5 bg-blue-200 rounded-full mr-2"></span>
                    {data!.roadmap.creator.username}
                  </a>
                </Link>
                ·
                <span className="ml-2">
                  {dayjs(data!.roadmap.createdAt).format("MMMM D, YYYY")}
                </span>
              </div>
              <h1 className="text-3xl text-gray-800 font-medium mb-1">
                {/* {data.roadmap.title} */}
                Visual Elements of User Interface
              </h1>
              <div className="text-sm text-gray-400 font-light tracking-wide mb-3">
                {/* {data.roadmap.description} */}A guide to learn everything
                fullstack.
              </div>
              <div className="flex flex-wrap mb-6">
                <span className="text-gray-400 text-xs font-medium bg-secondary px-2 py-1 mr-1 rounded">
                  Wedev
                </span>
                <span className="text-gray-400 text-xs font-medium bg-secondary px-2 py-1 mr-1 rounded">
                  Fullstack
                </span>
                <span className="text-gray-400 text-xs font-medium bg-secondary px-2 py-1 rounded">
                  Javascript
                </span>
              </div>
            </div>

            {/* CONTENT */}
            <div className="mb-32">
              {data!.roadmap.checkpoints.map((checkpoint, idx) => (
                <Checkpoint
                  idx={idx + 1}
                  checkpoint={checkpoint}
                  isAuth={!!me}
                  key={idx}
                />
              ))}
            </div>
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
    await client.query<RoadmapQuery, RoadmapQueryVariables>({
      query: ROADMAP_QUERY,
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

export default withAuth()(Roadmap);
