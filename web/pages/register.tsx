import { RedirectLink } from "components/element";
import { Form, InputField, SubmitButton } from "components/form";
import { Layout } from "components/layout";
import { useRegister } from "modules/auth";
import { Notification } from "modules/notification";
import React from "react";
import { z } from "zod";

type RegisterValues = {
  email: string;
  username: string;
  password: string;
};

const schema = z.object({
  email: z.string().email("Must be a valid email"),
  username: z.string().min(1, "Username is required"),
  password: z.string().min(8, "Password must be at least 8 characters long"),
});

function Register() {
  const { register: signup, loading } = useRegister();

  return (
    <Layout title="Sign up | Roadmap">
      <main className="grid grid-cols-12">
        <section className="col-start-2 col-end-7"></section>
        <section className="max-w-sm col-start-7 col-end-12">
          <h1 className="text-3xl text-gray-800 font-medium tracking-wide mt-8 mb-3">
            Sign up
          </h1>

          <Form<RegisterValues, typeof schema>
            onSubmit={signup}
            schema={schema}
          >
            {({ register, formState: { errors } }) => (
              <>
                <InputField
                  label="Email"
                  id="email"
                  error={errors.email}
                  register={register("email")}
                />
                <InputField
                  label="Username"
                  id="username"
                  error={errors.username}
                  register={register("username")}
                />
                <InputField
                  label="Password"
                  id="password"
                  type="password"
                  error={errors.password}
                  register={register("password")}
                />

                <SubmitButton label="Sign up" loading={loading} />
              </>
            )}
          </Form>
          <Notification type="text" style="mt-2" showOnly="error" />

          <div className="text-sm text-gray-500 text-center mt-6">
            Already have an account?{" "}
            <RedirectLink label="Log in" pathname="/login" />
          </div>
        </section>
      </main>
    </Layout>
  );
}

export default Register;
