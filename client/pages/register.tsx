import { useMutation } from "@apollo/client";
import Link from "next/link";
import Router, { useRouter } from "next/router";
import React, { useState } from "react";
import { useForm } from "react-hook-form";
import Alert from "../components/Alert";
import { REGISTER_MUTATION } from "../graphql/mutations/register";
import { ME_QUERY } from "../graphql/queries/me";

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

        console.log("REGISTER DATA", data);
        cache.writeQuery({
          query: ME_QUERY,
          data: {
            __typename: "Query",
            me: data.createUser,
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
      <h1 className="font-bold">Sign up</h1>
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

        <label htmlFor="username">Username</label>
        <input
          type="text"
          id="username"
          {...register("username", {
            required: { value: true, message: "Username is required" },
            maxLength: {
              value: 20,
              message: "Username must be less than 20 characters long",
            },
          })}
        />
        {errors.username && <span>{errors.username?.message}</span>}

        <label htmlFor="password">Password</label>
        <input
          type="password"
          id="password"
          {...register("password", {
            required: { value: true, message: "Password is required" },
          })}
        />
        {errors.password && <span>{errors.password?.message}</span>}

        <button type="submit">Sign up</button>
      </form>
      <div>
        Already have an account? <RedirectToLogin />
      </div>
    </div>
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
      <a>Log in</a>
    </Link>
  );
}

export default Register;
