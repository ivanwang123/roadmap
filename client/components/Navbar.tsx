import { useMutation } from "@apollo/client";
import Link from "next/link";
import { useRouter } from "next/router";
import React from "react";
import { LOGOUT_MUTATION } from "../graphql/mutations/logout";
import withAuth, { AuthChildProps } from "../hoc/withAuth";
import { getApolloClient } from "../lib/apollo-client";

function Navbar({ data: { me, loading, error } }: AuthChildProps) {
  const router = useRouter();

  const [logout] = useMutation(LOGOUT_MUTATION);

  const handleLogout = () => {
    logout()
      .then(() => {
        const client = getApolloClient();
        client.clearStore().then(() => {
          client.resetStore();
        });
      })
      .catch((e) => console.error(e));
  };

  console.log(me, error);
  let links = null;
  if (loading) {
    console.log("LOADING");
    links = null;
  } else if (!me || error) {
    let query: any = {
      redirect: router.query.redirect || router.asPath,
    };
    if (query.redirect === "/login" || query.redirect === "/register") {
      query = {};
    }
    links = (
      <>
        <Link href={{ pathname: "/login", query }}>
          <a className="text-gray-400 font-medium tracking-wide transition duration-300 hover:text-gray-800">
            Log in
          </a>
        </Link>
        <Link href={{ pathname: "/register", query }}>
          <a className="text-gray-400 font-medium tracking-wide transition duration-300 hover:text-gray-800">
            Sign up
          </a>
        </Link>
      </>
    );
  } else {
    links = (
      <>
        <button
          type="button"
          className="text-gray-400 font-medium tracking-wide transition duration-300 hover:text-gray-800 focus:outline-none"
          onClick={handleLogout}
        >
          Logout
        </button>
      </>
    );
  }

  return (
    <nav className="flex items-center w-full h-full border-b-2 border-secondary pl-10">
      <Link href="/">
        <a className="text-2xl text-gray-800 font-bold">ROADMAP</a>
      </Link>
      <div className="grid grid-flow-col gap-8 ml-auto mr-10">
        {/* TODO: Add current pathname as redirect */}
        {links}
      </div>
    </nav>
  );
}

export default withAuth()(Navbar);
