import { LinkButton, PrimaryButton, SecondaryButton } from "components/element";
import { useLogout, withAuth, WithAuthProps } from "modules/auth";
import Link from "next/link";
import { useRouter } from "next/router";
import React from "react";

function Navbar({ data: { me, loading, error } }: WithAuthProps) {
  console.log(me, error);

  let links = null;
  if (!loading) {
    if (!me || error) {
      links = <UnAuthenticatedLinks />;
    } else {
      links = <AuthenticatedLinks />;
    }
  }

  return (
    <nav className="flex items-center w-full h-full border-b border-gray-100 pl-10">
      <Link href="/">
        <a className="text-2xl text-gray-800 font-bold">ROADMAP</a>
      </Link>
      <div className="grid grid-flow-col place-items-center gap-8 ml-auto mr-10">
        {/* TODO: Add current pathname as redirect */}
        {links}
      </div>
    </nav>
  );
}

function AuthenticatedLinks() {
  const { logout } = useLogout();

  return (
    <>
      {/* <button
        type="button"
        className="text-gray-400 font-medium tracking-wide transition duration-300 hover:text-gray-800 focus:outline-none"
        onClick={logout}
      >
        Logout
      </button> */}
      <SecondaryButton label="Logout" onClick={logout} />
    </>
  );
}

function UnAuthenticatedLinks() {
  const router = useRouter();

  let query: any = {
    redirect: router.query.redirect || router.asPath,
  };
  if (query.redirect === "/login" || query.redirect === "/register") {
    query = {};
  }

  return (
    <>
      <Link href={{ pathname: "/register", query }}>
        <PrimaryButton label="Sign up" />
        {/* <a className="text-gray-400 font-medium tracking-wide transition duration-300 hover:text-gray-800">
          Sign up
        </a> */}
      </Link>
      <Link href={{ pathname: "/login", query }}>
        <LinkButton label="Login" />
        {/* <a className="text-gray-400 font-medium tracking-wide transition duration-300 hover:text-gray-800">
          Log in
        </a> */}
      </Link>
    </>
  );
}

const hoc = withAuth(Navbar);
export { hoc as Navbar };
