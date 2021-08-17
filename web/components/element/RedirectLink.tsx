import Link from "next/link";
import { useRouter } from "next/router";
import React from "react";
import { Url } from "url";

type Props = {
  label: string;
  pathname: string;
};

export function RedirectLink({ label, pathname }: Props) {
  const router = useRouter();

  let href: Partial<Url> = {
    pathname,
  };
  if (router.query.redirect !== undefined) {
    href["query"] = { redirect: router.query.redirect };
  }

  return (
    <Link href={href}>
      <a className="text-blue-500 hover:underline">{label}</a>
    </Link>
  );
}
