import { useLogout, withAuth, WithAuthProps } from "modules/auth";
import Link from "next/link";
import { useRouter } from "next/router";
import React from "react";

function Navbar({ data: { me, loading, error } }: WithAuthProps) {
  const router = useRouter();

  const { logout } = useLogout();

  // TODO: Clean this up
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
          onClick={logout}
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

const hoc = withAuth()(Navbar);
export { hoc as Navbar };
