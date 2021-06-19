import React, { useState } from "react";
import Router, { useRouter } from "next/router";
import Alert from "../components/Alert";
import { useMutation } from "@apollo/client";
import { useForm } from "react-hook-form";
import { ME_QUERY } from "../graphql/queries/me";
import { LOGIN_MUTATION } from "../graphql/mutations/login";
import Link from "next/link";

function Login() {
  const [error, setError] = useState<boolean>(false);
  const [message, setMessage] = useState<string>("");

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const [login] = useMutation(LOGIN_MUTATION);

  const onSubmit = (data: any) => {
    console.log("SUBMIT", data);
    login({
      variables: data,
      update: (cache, { data }) => {
        if (!data || !data.login) {
          return;
        }

        console.log("LOGIN DATA", data);
        cache.writeQuery({
          query: ME_QUERY,
          data: {
            __typename: "Query",
            me: data.login,
          },
        });
      },
    })
      .then(() => {
        if (Router.query.redirect !== undefined) {
          Router.push(Router.query.redirect as string);
        } else {
          Router.push("/");
        }
      })
      .catch((err) => {
        setError(true);
        setMessage(err.message);
      });
  };

  return (
    <div>
      <h1 className="font-bold">Login</h1>
      <Alert message={message} error={error} />

      <form className="flex flex-col" onSubmit={handleSubmit(onSubmit)}>
        <label htmlFor="email">Email</label>
        <input
          type="text"
          id="email"
          {...register("email", {
            required: { value: true, message: "Email is required" },
            pattern: {
              value: new RegExp(
                "^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$"
              ),
              message: "Must be a valid email",
            },
          })}
        />
        {errors.email && <span>{errors.email?.message}</span>}

        <label htmlFor="password">Password</label>
        <input
          type="password"
          id="password"
          {...register("password", {
            required: { value: true, message: "Password is required" },
            minLength: {
              value: 8,
              message: "Password must be at least 8 characters long",
            },
          })}
        />
        {errors.password && <span>{errors.password?.message}</span>}

        <button type="submit">Log in</button>
      </form>
      <div>
        Don't have an account? <RedirectToSignup />
      </div>
    </div>
  );
}

function RedirectToSignup() {
  const router = useRouter();
  let href: any = {
    pathname: "/register",
  };
  if (router.query.redirect !== undefined) {
    href["query"] = { redirect: router.query.redirect };
  }
  return (
    <Link href={href}>
      <a>Sign up</a>
    </Link>
  );
}

export default Login;
