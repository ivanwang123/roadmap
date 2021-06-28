import React from "react";
import { RoadmapType } from "../types/roadmapTypes";
import User from "../svgs/user.svg";
import Flag from "../svgs/flag-big.svg";
import Book from "../svgs/book.svg";
import Link from "next/link";

type Props = {
  roadmap?: RoadmapType;
};

function RoadmapCard({}: Props) {
  return (
    <article className="w-full px-5 pt-3 my-12 shadow">
      <div className="flex text-gray-400 text-sm font-semibold tracking-wide pr-4 mb-1">
        <div className="flex items-center mr-10">
          <User className="fill-current mr-2" width={20} height={20} /> 12
          followers
        </div>
        <div className="flex items-center mr-10">
          <Flag className="fill-current mr-2" width={20} height={20} />5
          checkpoints
        </div>
        <div className="flex items-center mr-10">
          <Book className="fill-current mr-2" width={20} height={20} />
          12 resources
        </div>
        <p className="text-sm font-normal ml-auto">Jul 5, 2020</p>
      </div>
      <div>
        <h3 className="text-2xl text-gray-800 font-bold tracking-wide">
          <Link href="/map/1">
            <a className="hover:underline">Full stack developer</a>
          </Link>
        </h3>
        <div className="flex flex-wrap my-2 tracking-wide">
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
        <p className="text-gray-500 line-clamp-3">
          In this tutorial, you'll learn how to apply linear algebra concepts to
          practical problems, how to work with vectors and matrices using Python
          and NumPy, how to model practical problems using linear systems, and
          how to solve linear systems using scipy.linalg.
        </p>
        <div className="flex items-center my-3">
          <Link href="/user/1">
            <a className="flex items-center text-gray-400 text-sm font-semibold tracking-wide transition duration-200 hover:text-gray-600">
              <div className="w-6 h-6 bg-red-200 rounded-full mr-1"></div>
              Stuart Little
            </a>
          </Link>
        </div>
        <div className="grid grid-cols-3 border-t-2 border-gray-100 mt-3">
          <div className="text-sm border-r-2 border-gray-100 p-3">
            <h6 className="text-gray-700 truncate">HTML</h6>
            <p className="text-gray-500 line-clamp-3">
              HTML is the backbone of a website. These resources go over the
              basics and when you finish reading through the tutorials, take the
              quiz on w3schools before moving on.
            </p>
          </div>
          <div className="text-sm border-r-2 border-gray-100 p-3">
            <h6 className="text-gray-700 truncate">CSS</h6>
            <p className="text-gray-500 line-clamp-3">
              CSS is the style of the wyle. Use it to spice up your website and
              add some color to your life. Look through the following links to
              get a better understanding of it.
            </p>
          </div>
          <div className="text-sm p-3">
            <h6 className="text-gray-700 truncate">Javascript</h6>
            <p className="text-gray-500 line-clamp-3">
              CSS is the style of the wyle. Use it to spice up your website and
              add some color to your life. Look through the following links to
              get a better understanding of it.
            </p>
          </div>
        </div>
      </div>
    </article>
  );
}

export default RoadmapCard;
