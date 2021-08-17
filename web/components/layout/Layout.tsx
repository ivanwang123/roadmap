import Head from "next/head";
import React, { ReactNode } from "react";
import { Navbar } from "./Navbar";

type Props = {
  children?: ReactNode;
  title?: string;
};

export function Layout({ children, title = "Roadmap" }: Props) {
  return (
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
}

export default Layout;
