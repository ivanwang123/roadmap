import Link from "next/link";
import React from "react";
import Layout from "../../components/Layout";
import Check from "../../svgs/check-two.svg";
import Close from "../../svgs/close-two.svg";
import Flag from "../../svgs/flag-two.svg";
import Circle from "../../svgs/circle.svg";
import Pin from "../../svgs/marker.svg";
import CircleTwo from "../../svgs/circle-three.svg";
import DownArrow from "../../svgs/down-arrow.svg";
import DashedArrowOne from "../../svgs/dashed-arrow-one.svg";
import DashedArrowTwo from "../../svgs/dashed-arrow-two.svg";
// TODO: Replace icons with non hand drawn

function Roadmap() {
  return (
    <Layout title="Map | Roadmap">
      <main className="sidebar-grid h-full bg-white">
        {/* NAVBAR */}
        {/* <section className="col-span-2"></section> */}

        {/* SIDEBAR */}
        <section className="w-64 h-full">
          <div className="sticky top-0 flex flex-col h-full bg-trueGray-100 pl-10 pt-16 shadow-inner overflow-auto">
            <h6 className="text-gray-400 font-bold mb-1">CHECKPOINTS</h6>
            <ol className="sidebar-grid gap-x-1 text-gray-400">
              <Check
                className="col-start-1 fill-current text-green-600"
                width={28}
                height={28}
              />
              <li className="flex items-center text-gray-700">HTML</li>
              <div className="w-1/2 h-4 justify-self-end border-l-2 border-gray-300 -my-1 mr-xs"></div>
              <Close
                className="col-start-1 fill-current text-red-500"
                width={28}
                height={28}
              />
              <li className="flex items-center">CSS</li>
              <div className="w-1/2 h-4 justify-self-end border-l-2 border-gray-300 -my-1 mr-xs"></div>
              <Flag
                className="col-start-1 fill-current text-yellow-500"
                width={28}
                height={28}
              />
              <li className="flex items-center">Javascript</li>
              <div className="w-1/2 h-4 justify-self-end border-l-2 border-gray-300 -my-1 mr-xs"></div>
              <CircleTwo
                className="col-start-1 fill-current text-gray-300"
                width={28}
                height={28}
              />
              <li className="flex items-center">Javascript</li>
            </ol>
          </div>
        </section>

        {/* MAIN */}
        <section className="pl-20 pr-32 overflow-auto">
          <div className="max-w-4xl">
            {/* HEADER */}
            <div className="col-start-4 col-end-11 my-16">
              <h1 className="text-5xl text-gray-800 font-bold tracking-wider my-3">
                Full stack developer
              </h1>
              <div className="flex flex-col text-gray-400 tracking-wide">
                <p className="mb-1">
                  <strong>12</strong> checkpoints · <strong>33</strong>{" "}
                  resources
                </p>
                <p>
                  Created by{" "}
                  <Link href="/user/id">
                    <a className="underline">Stuart Little</a>
                  </Link>{" "}
                  on June 17, 2021
                </p>
              </div>
            </div>

            {/* CONTENT */}
            <div className="col-start-4 col-end-11 mb-32">
              <div className="flex flex-col">
                <div className="flex items-center mb-3">
                  {/* <div className="relative grid place-items-center w-10 h-10 text-gray-800 font-bold px-6 rounded-xl">
                    <Circle
                      className="col-start-1 col-end-1 row-start-1 row-end-1"
                      width={40}
                      height={40}
                    />
                    <span className="col-start-1 col-end-1 row-start-1 row-end-1">
                      1
                    </span>
                  </div> */}
                  {/* <Pin width={50} height={50} /> */}
                  <div className="grid place-items-center w-10 h-10 bg-gray-800 text-white font-bold rounded-xl">
                    1
                  </div>
                  <h3 className="text-2xl text-gray-800 font-bold ml-4">
                    HTML
                  </h3>
                </div>
                <div className="flex flex-col">
                  <article className="w-full bg-white pt-4 pb-12 rounded ">
                    <p className="text-gray-700 tracking-wide leading-7">
                      HTML is the backbone of a website. These resources go over
                      the basics and when you finish reading through the
                      tutorials, take the quiz on w3schools before moving on.
                    </p>
                    <h6 className="text-gray-600 font-bold mt-4">RESOURCES</h6>
                    <div className="flex flex-wrap mt-2">
                      <div className="text-gray-500 border p-3 rounded">
                        <span>w3schools.com</span>
                      </div>
                    </div>
                  </article>
                  <div className="grid grid-flow-col gap-0 place-items-center bg-white px-6 py-2 ml-auto -mt-9 mr- rounded- border- ">
                    <Check
                      className="fill-current text-green-500 border-b-2 p-1 rounded- cursor-pointer hover:bg-green-100 hover:border-green-300 hover:border-b-4"
                      width={60}
                      height={40}
                    />
                    <Close
                      className="fill-current text-red-500 border-b-2 p-1 rounded- cursor-pointer hover:bg-red-100 hover:border-red-300 hover:border-b-4"
                      width={60}
                      height={40}
                    />
                    <Flag
                      className="fill-current text-yellow-500 border-b-2 p-1 rounded- cursor-pointer hover:bg-yellow-100 hover:border-yellow-300 hover:border-b-4"
                      width={60}
                      height={40}
                    />
                  </div>
                </div>
              </div>

              <div className="">
                <DashedArrowOne
                  className="fill-current text-gray-600 transform ml-auto"
                  width={150}
                  height={150}
                />
                {/* <DashedArrowTwo
                  className="transform rotate-6"
                  width={150}
                  height={150}
                /> */}
                {/* <ArrowOne
                  className="transform -rotate-135"
                  width={100}
                  height={100}
                /> */}
              </div>

              <div className="flex flex-col">
                <div className="flex items-center mb-3">
                  {/* <div className="relative grid place-items-center w-10 h-10 text-gray-800 font-bold px-6 rounded-xl">
                    <Circle
                      className="col-start-1 col-end-1 row-start-1 row-end-1"
                      width={40}
                      height={40}
                    />
                    <span className="col-start-1 col-end-1 row-start-1 row-end-1">
                      1
                    </span>
                  </div> */}
                  {/* <Pin width={50} height={50} /> */}
                  <div className="grid place-items-center w-10 h-10 bg-gray-400 text-white font-bold rounded-xl">
                    2
                  </div>
                  <h3 className="text-2xl text-gray-400 font-bold ml-4">CSS</h3>
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
                    <Flag
                      className="fill-current stroke-current text-yellow-500 p-1 rounded-full cursor-pointer hover:bg-yellow-100"
                      width={40}
                      height={40}
                    />
                  </div>
                </div>
              </div>

              <div className="">
                <DashedArrowOne
                  className="fill-current text-gray-400 transform ml-auto"
                  width={150}
                  height={150}
                />
                {/* <DashedArrowTwo
                  className="transform rotate-6"
                  width={150}
                  height={150}
                /> */}
                {/* <ArrowOne
                  className="transform -rotate-135"
                  width={100}
                  height={100}
                /> */}
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
                <DashedArrowOne
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

export default Roadmap;
