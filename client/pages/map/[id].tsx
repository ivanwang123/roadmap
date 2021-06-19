import Link from "next/link";
import React from "react";
import Layout from "../../components/Layout";

function Roadmap() {
  return (
    <Layout title="Map | Roadmap">
      <main className="sidebar-grid h-full px-32">
        {/* SIDEBAR */}
        <section className="w-40 h-full">
          <div className="sticky top-0 flex flex-col justify-center h-screen">
            <h6 className="text-gray-400 font-bold mb-1">CHECKPOINTS</h6>
            <ol className="list-decimal text-gray-400">
              <li>HTML</li>
              <li>CSS</li>
              <li>Javascript</li>
            </ol>
          </div>
        </section>

        {/* MAIN */}
        <div className="max-w-3xl">
          {/* HEADER */}
          <section className="col-start-4 col-end-11 my-16">
            <h1 className="text-4xl text-gray-800 font-bold tracking-wider my-3">
              Full stack developer
            </h1>
            <div className="flex flex-col text-gray-400">
              <p className="mb-1">
                <strong>12</strong> checkpoints · <strong>33</strong> resources
              </p>
              <p>
                Created by{" "}
                <Link href="/user/id">
                  <a className="underline">Stuart Little</a>
                </Link>{" "}
                on June 17, 2021
              </p>
            </div>
          </section>

          {/* CONTENT */}
          <section className="col-start-4 col-end-11 mb-32">
            <div className="flex flex-col">
              <div className="flex items-center mb-3">
                <div className="grid place-items-center w-10 h-10 bg-gray-500 text-white font-bold rounded-xl">
                  1
                </div>
                <h3 className="text-2xl text-gray-500 ml-3">HTML</h3>
              </div>
              <article className="w-full px-6 py-4 shadow">
                <p className="text-gray-700">
                  HTML is the backbone of a website. These resources go over the
                  basics and when you finish reading through the tutorials, take
                  the quiz on w3schools before moving on.
                </p>
                <div className="flex flex-wrap mt-3">
                  <div className="text-gray-500 border p-3 rounded">
                    <span>w3schools.com</span>
                  </div>
                </div>
              </article>
            </div>

            <div className="h-32 border-l-2 border-gray-500 ml-5 my-4"></div>

            <div className="flex flex-col">
              <div className="flex items-center mb-3">
                <div className="grid place-items-center w-10 h-10 bg-gray-300 text-white font-bold rounded-xl">
                  2
                </div>
                <h3 className="text-2xl text-gray-500 ml-3">CSS</h3>
              </div>
              <article className="w-full px-6 py-4 shadow">
                <p className="text-gray-700">
                  CSS is the style of the wyle. Use it to spice up your website
                  and add some color to your life. Look through the following
                  links to get a better understanding of it.
                </p>
                <div className="flex flex-wrap mt-3">
                  <div className="text-gray-500 border p-3 rounded">
                    <span>w3schools.com</span>
                  </div>
                </div>
              </article>
            </div>

            <div className="h-32 border-l-2 border-gray-300 ml-5 my-4"></div>

            <div className="flex flex-col">
              <div className="flex items-center mb-3">
                <div className="grid place-items-center w-10 h-10 bg-gray-300 text-white font-bold rounded-xl">
                  3
                </div>
                <h3 className="text-2xl text-gray-500 ml-3">Javascript</h3>
              </div>
              <article className="w-full px-6 py-4 shadow">
                <p className="text-gray-700">
                  CSS is the style of the wyle. Use it to spice up your website
                  and add some color to your life. Look through the following
                  links to get a better understanding of it.
                </p>
                <div className="flex flex-wrap mt-3">
                  <div className="text-gray-500 border p-3 rounded">
                    <span>w3schools.com</span>
                  </div>
                </div>
              </article>
            </div>
          </section>
        </div>
      </main>
    </Layout>
  );
}

export default Roadmap;
