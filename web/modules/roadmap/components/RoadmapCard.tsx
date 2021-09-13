import { Icon } from "components/element";
import Link from "next/link";
import React from "react";
import Book from "svgs/book.svg";
import Flag from "svgs/flag-big.svg";
import User from "svgs/user.svg";
import { RoadmapInfoFieldsFragment } from "types/graphql-generated";
import { pluralize } from "utils";

type Props = {
  roadmap: RoadmapInfoFieldsFragment;
};

export function RoadmapCard({}: Props) {
  return (
    <article className="max-w-2xl bg-white p-6 my-14 rounded shadow-light">
      <div>
        <div className="flex items-center mb-2">
          <Link href="/user/1">
            <a className="flex items-center text-sm tracking-wide transition duration-300 hover:text-hover">
              <div className="w-5 h-5 bg-blue-200 rounded-full mr-2"></div>
              Stuart Little
            </a>
          </Link>
          <div className="text-sm text-gray-400 tracking-wide ml-auto">
            JUL 17, 2020
          </div>
        </div>
        <h3 className="text-xl font-semibold mb-1">
          <Link href="/map/1">
            <a className="transition duration-300 hover:text-hover">
              Full stack developer
            </a>
          </Link>
        </h3>
        <p className="text-sm text-gray-400 tracking-wide line-clamp-3 mb-4">
          In this tutorial, you'll learn how to apply linear algebra concepts to
          practical problems, how to work with vectors and matrices using Python
          and NumPy, how to model practical problems using linear systems, and
          how to solve linear systems using scipy.linalg.
        </p>
        <div className="flex text-gray-500 text-sm tracking-wide pr-4 mb-5">
          <div className="flex items-center mr-8">
            <Icon icon={User} size={12} />
            <span className="ml-2">
              <span className="font-semibold">12</span>{" "}
              {pluralize("follower", 12)}
            </span>
          </div>
          <div className="flex items-center mr-8">
            <Icon icon={Flag} size={12} />
            <span className="ml-2">
              <span className="font-semibold">7</span>{" "}
              {pluralize("checkpoint", 7)}
            </span>
          </div>
          <div className="flex items-center">
            <Icon icon={Book} size={12} />
            <span className="ml-2">
              <span className="font-semibold">12</span>{" "}
              {pluralize("resource", 12)}
            </span>
          </div>
        </div>
        <div className="grid grid-cols-3">
          <div className="text-sm tracking-wide border-r-2 border-secondary px-2">
            <h6 className="truncate mb-1">HTML</h6>
            <p className="text-gray-400 font-light line-clamp-3">
              HTML is the backbone of a website. These resources go over the
              basics and when you finish reading through the tutorials, take the
              quiz on w3schools before moving on.
            </p>
          </div>
          <div className="text-sm tracking-wide border-r-2 border-secondary px-2">
            <h6 className="truncate mb-1">CSS</h6>
            <p className="text-gray-400 font-light line-clamp-3">
              CSS is the style of the wyle. Use it to spice up your website and
              add some color to your life. Look through the following links to
              get a better understanding of it.
            </p>
          </div>
          <div className="text-sm tracking-wide px-2">
            <h6 className="truncate mb-1">Javascript</h6>
            <p className="text-gray-400 font-light line-clamp-3">
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
