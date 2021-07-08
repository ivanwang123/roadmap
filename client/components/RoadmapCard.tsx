import React from "react";
import User from "../svgs/user.svg";
import Flag from "../svgs/flag-big.svg";
import Book from "../svgs/book.svg";
import Link from "next/link";
import Icon from "./Icon";
import { RoadmapInfoFieldsFragment } from "../graphql/generated/generated";

type Props = {
  roadmap: RoadmapInfoFieldsFragment;
};

function RoadmapCard({}: Props) {
  return (
    <article className="max-w-2xl bg-white p-6 my-14 rounded shadow-light">
      <div>
        <div className="flex items-center">
          <Link href="/user/1">
            <a className="flex items-center text-gray-800 text-sm tracking-wide transition duration-300 hover:text-hover">
              <div className="w-5 h-5 bg-blue-200 rounded-full mr-2"></div>
              Stuart Little
            </a>
          </Link>
          <div className="text-xs text-gray-400 font-light tracking-wide ml-auto">
            July 17, 2020
          </div>
        </div>
        <h3 className="text-xl text-gray-800 font-medium mt-1 mb-3">
          <Link href="/map/1">
            <a className="transition duration-300 hover:text-hover">
              Full stack developer
            </a>
          </Link>
        </h3>
        <p className="text-sm font-light tracking-wide text-gray-400 line-clamp-3">
          In this tutorial, you'll learn how to apply linear algebra concepts to
          practical problems, how to work with vectors and matrices using Python
          and NumPy, how to model practical problems using linear systems, and
          how to solve linear systems using scipy.linalg.
        </p>
        <div className="flex text-gray-500 text-xs tracking-wide pr-4 mt-3">
          <div className="flex items-center mr-8">
            <Icon icon={User} size={12} />
            <span className="ml-2">12 followers</span>
          </div>
          <div className="flex items-center mr-8">
            <Icon icon={Flag} size={12} />
            <span className="ml-2">7 checkpoints</span>
          </div>
          <div className="flex items-center">
            <Icon icon={Book} size={12} />
            <span className="ml-2">12 resources</span>
          </div>
        </div>
        <div className="grid grid-cols-3 mt-7">
          <div className="tracking-wide border-r-2 border-secondary px-2">
            <h6 className="text-sm text-gray-800 truncate">HTML</h6>
            <p className="text-xs font-light text-gray-400 line-clamp-3">
              HTML is the backbone of a website. These resources go over the
              basics and when you finish reading through the tutorials, take the
              quiz on w3schools before moving on.
            </p>
          </div>
          <div className="tracking-wide border-r-2 border-secondary px-2">
            <h6 className="text-sm text-gray-800 truncate">CSS</h6>
            <p className="text-xs font-light text-gray-400 line-clamp-3">
              CSS is the style of the wyle. Use it to spice up your website and
              add some color to your life. Look through the following links to
              get a better understanding of it.
            </p>
          </div>
          <div className="tracking-wide px-2">
            <h6 className="text-sm text-gray-800 truncate">Javascript</h6>
            <p className="text-xs font-light text-gray-400 line-clamp-3">
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
