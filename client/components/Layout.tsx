import React, { ReactNode } from "react";
import Head from "next/head";
import Link from "next/link";

type Props = {
  children?: ReactNode;
  title?: string;
};

const Layout = ({ children, title = "Roadmap" }: Props) => (
  <div className="navbar-grid w-full h-full">
    <Head>
      <title>{title}</title>
      <meta charSet="utf-8" />
      <meta name="viewport" content="initial-scale=1.0, width=device-width" />
    </Head>
    <header className="h-20">
      <nav className="flex items-center w-full h-20 border-b-2 pl-10">
        <Link href="/">
          <a className="text-2xl text-blueGray-700 font-bold">ROADMAP</a>
        </Link>
        <div className="grid grid-flow-col gap-8 ml-auto mr-10">
          {/* TODO: Add current pathname as redirect */}
          <Link href="/login">
            <a className="">Login</a>
          </Link>
          <Link href="/register">
            <a>Sign up</a>
          </Link>
        </div>
      </nav>
    </header>
    {children}
    <footer></footer>
  </div>
);

export default Layout;
