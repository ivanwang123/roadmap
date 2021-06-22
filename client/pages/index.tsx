import React from "react";
import Layout from "../components/Layout";

// TODO: Display roadmaps
function Home() {
  return (
    <Layout title="Roadmap">
      <main className="grid grid-cols-12">
        <section className="col-start-4 col-end-11">
          <h1 className="text-5xl text-gray-800 font-bold tracking-wider mt-8 mb-3">
            Explore
          </h1>
          <article className="w-full px-5 pt-3 shadow">
            <div className="flex text-gray-800 font-semibold">
              <p className="text-gray-700">5 checkpoints · 12 resources</p>
              <p className="text-gray-400 text-sm font-normal ml-auto">
                Jul 5, 2020
              </p>
            </div>
            <h3 className="text-2xl text-gray-800 font-bold">
              Full stack developer
            </h3>
            <div className="flex flex-wrap my-2">
              <div className="bg-gray-100 text-gray-400 text-xs font-bold px-1.5 py-0.5 mr-1 rounded">
                Webdev
              </div>
              <div className="bg-gray-100 text-gray-400 text-xs font-bold px-1.5 py-0.5 mr-1 rounded">
                Fullstack
              </div>
              <div className="bg-gray-100 text-gray-400 text-xs font-bold px-1.5 py-0.5 mr-1 rounded">
                Frontend
              </div>
              <div className="bg-gray-100 text-gray-400 text-xs font-bold px-1.5 py-0.5 mr-1 rounded">
                Backend
              </div>
            </div>
            <p className="text-gray-500">
              In this tutorial, you'll learn how to apply linear algebra
              concepts to practical problems, how to work with vectors and
              matrices using Python and NumPy, how to model practical problems
              using linear systems, and how to solve linear systems using
              scipy.linalg.
            </p>
            <div className="flex items-center mt-2">
              <span className="w-6 h-6 bg-red-200 rounded-full"></span>
              <span className="text-gray-400 font-semibold ml-2">
                Stuart Little
              </span>
            </div>
            <div className="grid grid-cols-3 border-t-2 border-gray-100 mt-3">
              <div className="text-sm border-r-2 border-gray-100 p-3">
                <h6 className="text-gray-700 truncate">HTML</h6>
                <p className="text-gray-500 line-clamp-3">
                  HTML is the backbone of a website. These resources go over the
                  basics and when you finish reading through the tutorials, take
                  the quiz on w3schools before moving on.
                </p>
              </div>
              <div className="text-sm border-r-2 border-gray-100 p-3">
                <h6 className="text-gray-700 truncate">CSS</h6>
                <p className="text-gray-500 line-clamp-3">
                  CSS is the style of the wyle. Use it to spice up your website
                  and add some color to your life. Look through the following
                  links to get a better understanding of it.
                </p>
              </div>
              <div className="text-sm p-3">
                <h6 className="text-gray-700 truncate">Javascript</h6>
                <p className="text-gray-500 line-clamp-3">
                  CSS is the style of the wyle. Use it to spice up your website
                  and add some color to your life. Look through the following
                  links to get a better understanding of it.
                </p>
              </div>
            </div>
          </article>
        </section>
      </main>
    </Layout>
  );
}

export default Home;
