import { useMutation } from "@apollo/client";
import Link from "next/link";
import Router, { useRouter } from "next/router";
import React, { useState } from "react";
import { useForm } from "react-hook-form";
import Alert from "../components/Alert";
import Layout from "../components/Layout";
import { REGISTER_MUTATION } from "../graphql/mutations/register";
import { ME_QUERY } from "../graphql/queries/me";
import { getApolloClient } from "../lib/apollo-client";

function Register() {
  const [signup] = useMutation(REGISTER_MUTATION);
  const [error, setError] = useState<boolean>(false);
  const [message, setMessage] = useState<string>("");

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const onSubmit = (data: any) => {
    console.log("SUBMIT", data);
    signup({
      variables: data,
      update: (cache, { data }) => {
        if (!data || !data.createUser) {
          return;
        }

        cache.writeQuery({
          query: ME_QUERY,
          data: {
            me: data.createUser,
          },
        });
      },
    })
      .then(() => {
        const client = getApolloClient();
        client.resetStore();
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
    <Layout title="Sign up | Roadmap">
      <main className="grid grid-cols-12">
        <section className="col-start-2 col-end-7"></section>
        <section className="max-w-sm col-start-7 col-end-12">
          <h1 className="text-3xl text-gray-800 font-medium tracking-wide mt-8 mb-3">
            Sign up
          </h1>
          {/* <Alert message={message} error={error} /> */}

          <form className="flex flex-col" onSubmit={handleSubmit(onSubmit)}>
            <div className="flex flex-col my-8">
              <label
                htmlFor="email"
                className="text-gray-800 font-medium tracking-wide"
              >
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
              <label
                htmlFor="username"
                className="text-gray-800 font-medium tracking-wide"
              >
                Username
              </label>
              <input
                type="text"
                id="username"
                className={`border-b-2 pt-1 focus:border-gray-800 focus:outline-none ${
                  errors.email && "border-red-500"
                }`}
                {...register("username", {
                  required: { value: true, message: "Username is required" },
                  maxLength: {
                    value: 20,
                    message: "Username must be less than 20 characters long",
                  },
                })}
              />
              <span className="h-6 text-red-500 mt-1">
                {errors.username?.message}
              </span>
            </div>

            <div className="flex flex-col mb-8">
              <label
                htmlFor="password"
                className="text-gray-800 font-medium tracking-wide"
              >
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
                })}
              />
              <span className="h-6 text-red-500 mt-1">
                {errors.password?.message}
              </span>
            </div>

            <button
              type="submit"
              className="bg-green-500 text-white font-medium tracking-wide py-2 rounded"
            >
              Sign up
            </button>
          </form>
          <div className="text-sm text-gray-500 text-center mt-6">
            Already have an account? <RedirectToLogin />
          </div>
        </section>
      </main>
    </Layout>
  );
}

function RedirectToLogin() {
  const router = useRouter();
  let href: any = {
    pathname: "/login",
  };
  if (router.query.redirect !== undefined) {
    href["query"] = { redirect: router.query.redirect };
  }
  return (
    <Link href={href}>
      <a className="text-blue-500 hover:underline">Log in</a>
    </Link>
  );
}

export default Register;
