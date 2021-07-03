import { useQuery } from "@apollo/client";
import dayjs from "dayjs";
import { GetServerSideProps } from "next";
import Link from "next/link";
import { useRouter } from "next/router";
import React from "react";
import Checkpoint from "../../components/Checkpoint";
import CheckpointStatus from "../../components/CheckpointStatus";
import Layout from "../../components/Layout";
import { ROADMAP_QUERY } from "../../graphql/queries/roadmapById";
import { addApolloState, getApolloClient } from "../../lib/apollo-client";
import Book from "../../svgs/book.svg";
import Check from "../../svgs/check.svg";
import Close from "../../svgs/close.svg";
import DashedArrow from "../../svgs/dashed-arrow.svg";
import Flag from "../../svgs/flag-big.svg";
import Skip from "../../svgs/skip.svg";
import User from "../../svgs/user.svg";
import { CheckpointType } from "../../types/checkpointTypes";

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

function Roadmap() {
  const router = useRouter();
  const { data, loading, error } = useQuery(ROADMAP_QUERY, {
    variables: { id: router.query.id },
  });

  if (loading) return <h1>Loading</h1>;
  if (error) return <h1>Error {error}</h1>;

  return (
    <Layout title="Map | Roadmap">
      <main className="sidebar-grid h-full max-h-full overflow-hidden bg-white">
        {/* SIDEBAR */}
        <section className="w-64 h-full">
          <div className="sticky top-0 flex flex-col h-full bg-tertiary pl-10 pt-16 shadow-inne overflow-auto">
            <h6 className="text-gray-400 font-bold mb-4">CHECKPOINTS</h6>
            <div className="checkpoints-grid gap-x-1 items-center">
              {data.roadmap.checkpoints.map(
                (checkpoint: CheckpointType, idx: number) => (
                  <CheckpointStatus
                    checkpoint={checkpoint.title}
                    status={checkpoint.status}
                    isLast={idx === data.roadmap.checkpoints.length - 1}
                  />
                )
              )}
              {/* <CheckpointStatus checkpoint="HTML" status="skip" />
              <CheckpointStatus checkpoint="HTML" status="current" />
              <CheckpointStatus checkpoint="HTML" status="incomplete" /> */}
            </div>
          </div>
        </section>

        {/* MAIN */}
        <section className="px-20 overflow-auto">
          <div className="max-w-4xl">
            {/* HEADER */}
            <div className="col-start-4 col-end-11 my-16">
              <h1 className="text-5xl text-gray-800 font-bold tracking-wide my-3">
                {data.roadmap.title}
              </h1>
              <div className="flex text-gray-400 tracking-wide mt-4 mb-3">
                <div className="flex items-center mr-10">
                  <span className="grid place-items-center bg-secondary p-1">
                    <User className="fill-current" width={20} height={20} />
                  </span>
                  <span className="font-semibold mr-1">
                    {data.roadmap.followers.length}
                  </span>{" "}
                  followers
                </div>
                <div className="flex items-center mr-10">
                  <Flag className="fill-current mr-2" width={20} height={20} />
                  <span className="font-semibold mr-1">
                    {data.roadmap.checkpoints.length}
                  </span>{" "}
                  checkpoints
                </div>
                <div className="flex items-center mr-10">
                  <Book className="fill-current mr-2" width={20} height={20} />
                  <span className="font-semibold mr-1">12</span> resources
                </div>
              </div>
              <div className="text-gray-400 tracking-wide">
                <div className="flex items-center mb-2">
                  Created by{" "}
                  <span className="mx-2">
                    <Link href="/user/1">
                      <a className="flex items-center text-gray-400 font-semibold tracking-wide transition duration-200 hover:text-gray-600">
                        <span className="w-6 h-6 bg-red-200 rounded-full mr-1"></span>
                        {data.roadmap.creator.username}
                      </a>
                    </Link>
                  </span>
                  on {dayjs(data.roadmap.createdAt).format("MMMM D, YYYY")}
                </div>
                <div>
                  Last updated:{" "}
                  {dayjs(data.roadmap.updatedAt).format("MMMM D, YYYY")}
                </div>
              </div>
            </div>

            {/* CONTENT */}
            <div className="col-start-4 col-end-11 mb-32">
              {/* <Checkpoint reference={(el) => createObserver(el)} /> */}
              {data.roadmap.checkpoints.map(
                (checkpoint: CheckpointType, idx: number) => (
                  <Checkpoint idx={idx + 1} checkpoint={checkpoint} key={idx} />
                )
              )}

              <div className="flex flex-col">
                <div className="flex items-center mb-3">
                  <div className="grid place-items-center w-10 h-10 bg-gray-400 text-white font-bold rounded-xl">
                    2
                  </div>
                  <h3 className="text-2xl text-gray-400 font-bold tracking-wide ml-4">
                    CSS
                  </h3>
                </div>
                <div className="flex flex-col">
                  <article className="w-full bg-white px-6 -ml-6 pt-4 pb-12 rounded ">
                    <p className="text-gray-500 tracking-wide leading-7">
                      HTML is the backbone of a website. These resources go over
                      the basics and when you finish reading through the
                      tutorials, take the quiz on w3schools before moving on.
                    </p>
                    <h6 className="text-gray-400 font-bold mt-4">RESOURCES</h6>
                    <div className="flex flex-wrap mt-2">
                      <div className="text-gray-500 border p-3 rounded">
                        <span>w3schools.com</span>
                      </div>
                    </div>
                  </article>
                  <div className="grid grid-flow-col gap-4 place-items-center bg-white px-6 py-4 ml-auto -mt-9 mr-12 rounded-full border-2">
                    <Check
                      className="fill-current text-green-500 rounded-full cursor-pointer hover:bg-green-100"
                      width={40}
                      height={40}
                    />
                    <Close
                      className="fill-current text-red-500 p-1 rounded-full cursor-pointer hover:bg-red-100"
                      width={40}
                      height={40}
                    />
                    <Skip
                      className="fill-current stroke-current text-yellow-500 p-1 rounded-full cursor-pointer hover:bg-yellow-100"
                      width={40}
                      height={40}
                    />
                  </div>
                </div>
              </div>

              <div>
                <DashedArrow
                  className="fill-current text-gray-400 transform ml-auto"
                  width={150}
                  height={150}
                />
              </div>

              <div className="flex flex-col">
                <div className="flex items-center mb-3">
                  <div className="grid place-items-center w-10 h-10 bg-gray-300 text-white font-bold rounded-xl">
                    2
                  </div>
                  <h3 className="text-2xl text-gray-800 ml-3">CSS</h3>
                </div>
                <article className="w-full bg-white px-6 py-4 rounded shadow">
                  <p className="text-gray-500 tracking-wide leading-7">
                    CSS is the style of the wyle. Use it to spice up your
                    website and add some color to your life. Look through the
                    following links to get a better understanding of it.
                  </p>
                  <div className="flex flex-wrap mt-3">
                    <div className="text-gray-500 border p-3 rounded">
                      <span>w3schools.com</span>
                    </div>
                  </div>
                </article>
              </div>

              <div className="">
                <DashedArrow
                  className="fill-current text-gray-400 transform ml-auto"
                  width={150}
                  height={150}
                />
                {/* <ArrowTwo
                  className="transform rotate-135"
                  width={100}
                  height={100}
                /> */}
              </div>

              <div className="flex flex-col">
                <div className="flex items-center mb-3">
                  <div className="grid place-items-center w-10 h-10 bg-gray-300 text-white font-bold rounded-xl">
                    3
                  </div>
                  <h3 className="text-2xl text-gray-800 ml-3">Javascript</h3>
                </div>
                <article className="w-full bg-white px-6 py-4 rounded shadow">
                  <p className="text-gray-500 tracking-wide leading-7">
                    CSS is the style of the wyle. Use it to spice up your
                    website and add some color to your life. Look through the
                    following links to get a better understanding of it.
                  </p>
                  <div className="flex flex-wrap mt-3">
                    <div className="text-gray-500 border p-3 rounded">
                      <span>w3schools.com</span>
                    </div>
                  </div>
                </article>
              </div>
            </div>
          </div>
        </section>
      </main>
    </Layout>
  );
}

export const getServerSideProps: GetServerSideProps = async (ctx) => {
  const id = ctx.params?.id;

  const client = getApolloClient();

  try {
    await client.query({
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

export default Roadmap;
