"use client";
import Link from "next/link";
import GoogleLogin from "@/app/ui/components/google_login";
import Input from "@/app/ui/components/input";
import FormLogo from "@/app/ui/icons/logo";
import FormButton from "@/app/ui/components/form_button";
import loginAction from "@/app/lib/server_actions/login";
import { useFormState } from "react-dom";
import { Suspense, useContext, useEffect } from "react";
import { redirect, useSearchParams } from "next/navigation";
import { AlertDismisDispatchContext } from "@/app/lib/data/contexts/alert_dismiss_context";
import { useRouter } from "next/navigation";
import BackToAccounts from "../components/back_to_accounts";

export default function LoginForm() {
  const searchParams = useSearchParams();
  const [formState, formAction] = useFormState(loginAction, {
    message: '',
    success: false
  })
  const router = useRouter();
  const alertDispatch = useContext(AlertDismisDispatchContext);
  useEffect(() => {
    if (formState.message!.length > 0) {
      if (alertDispatch) {
        alertDispatch({
          type: "update",
          state: {
            show: true,
            message: formState.message || "",
            type: formState.success ? "info" : "error"
          }
        })
      }
    }
    if (formState.response) {
      if (formState.response!.token_type == "Otp") {
        router.push("/otp?" + searchParams.toString());
      } else if (formState.response!.token_type == "Bearer") {
        redirect("/profile?" + searchParams.toString())
      }
    }
  }, [formState, alertDispatch, router, searchParams])
  return (
    <section className="flex flex-col justify-center items-center my-auto h-full">
      <form className="container border rounded-md w-full max-w-md flex flex-col p-5 space-y-2" action={formAction}>
        <div className="h-5"></div>
        <FormLogo />
        <h3 className="text-md font-bold text-slate-700 text-center">Next App</h3>
        <p className="font-semibold text-slate-600 text-sm text-center">User Login Form</p>
        <div className="h-3"></div>
        <div>
          <Input name="username" label="Username" placeHolder="Username or Email" />
        </div>
        <div>
          <Input name="password" placeHolder="Password" type="password" label="Password" autoComplete="current-password" />
        </div>
        <div className="h-1"></div>
        <FormButton className="btn-primary" message="Login" pendingMessage="Logging In..." />
        <p className="text-sm text-slate-600">If you don&apos;t have an account, please signup <Link href={{ pathname: '/signup', query: searchParams.toString() }} className="text-blue-400">here</Link></p>
        <p className="text-center text-sm font-bold">Or</p>
        <GoogleLogin />
        <Suspense>
          <BackToAccounts />
        </Suspense>
        <div className="h-5"></div>
      </form>
    </section>
  );
}