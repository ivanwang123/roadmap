import React, { ReactNode } from "react";
import Head from "next/head";
import Navbar from "./Navbar";

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
    <header className="h-16">
      <Navbar />
    </header>
    {children}
    <footer></footer>
  </div>
);

export default Layout;
