import React, { useState } from "react";
import Router, { useRouter } from "next/router";
import Alert from "../components/Alert";
import { useMutation } from "@apollo/client";
import { useForm } from "react-hook-form";
import { ME_QUERY } from "../graphql/queries/me";
import { LOGIN_MUTATION } from "../graphql/mutations/login";
import Link from "next/link";
import Layout from "../components/Layout";

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
    <Layout title="Login | Roadmap">
      <main className="grid grid-cols-12">
        {/* TODO: Add side image */}
        <section className="col-start-2 col-end-7"></section>
        <section className="max-w-sm col-start-7 col-end-12">
          <h1 className="text-5xl text-gray-800 font-bold tracking-wider mt-8 mb-3">
            Log in
          </h1>
          {/* TODO: Add error alerts */}
          {/* <Alert message={"message"} error={error} /> */}

          <form className="flex flex-col" onSubmit={handleSubmit(onSubmit)}>
            <div className="flex flex-col my-8">
              <label htmlFor="email" className="text-gray-800 font-semibold">
                Email
              </label>
              <input
                type="text"
                id="email"
                className={`border-b-2 pt-1 focus:border-gray-800 focus:outline-none ${
                  errors.email && "border-red-500"
                }`}
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
              <span className="h-6 text-red-500 mt-1">
                {errors.email?.message}
              </span>
            </div>

            <div className="flex flex-col mb-8">
              <label htmlFor="password" className="text-gray-800 font-semibold">
                Password
              </label>
              <input
                type="password"
                id="password"
                className={`border-b-2 pt-1 focus:border-gray-800 focus:outline-none ${
                  errors.email && "border-red-500"
                }`}
                {...register("password", {
                  required: { value: true, message: "Password is required" },
                  minLength: {
                    value: 8,
                    message: "Password must be at least 8 characters long",
                  },
                })}
              />
              <span className="h-6 text-red-500 mt-1">
                {errors.password?.message}
              </span>
            </div>

            <button
              type="submit"
              className="bg-green-500 text-white font-bold py-2 rounded"
            >
              Log in
            </button>
          </form>
          <div className="text-sm text-gray-500 text-center mt-6">
            Don't have an account? <RedirectToSignup />
          </div>
        </section>
      </main>
    </Layout>
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
      <a className="text-blue-500 hover:underline">Sign up</a>
    </Link>
  );
}

export default Login;
