"use client";
import Link from "next/link";
import GoogleSignup from "@/app/ui/components/google_signup";
import FormLogo from "@/app/ui/icons/logo";
import { redirect, useSearchParams } from "next/navigation";
import FormButton from "@/app/ui/components/form_button";
import Input from "@/app/ui/components/input";
import { useFormState } from "react-dom";
import signup from "@/app/lib/server_actions/signup";
import { ChangeEvent, useContext, useEffect, useRef, useState } from "react";
import { AlertDismisDispatchContext } from "@/app/lib/data/contexts/alert_dismiss_context";
import { useRouter } from "next/navigation";

export default function SignupForm() {
    const searchParams = useSearchParams();
    const [formState, formAction] = useFormState(signup, {
        checkPass: false,
        success: false,
        response: undefined,
        message: "",
    });
    const [isRightPass, setIsRightPass] = useState(true);
    const [pwd, setPwd] = useState("");
    const [cPass, setCPass] = useState("");
    const alertDispatch = useContext(AlertDismisDispatchContext);
    const router = useRouter();

    const onChangePwd = (e: ChangeEvent<HTMLInputElement>) => {
        if (formState.checkPass && e.target.value != cPass) {
            setIsRightPass(false);
        } else {
            setIsRightPass(true);
        }
        setPwd(e.target.value);
    }
    const onChangeCPass = (e: ChangeEvent<HTMLInputElement>) => {
        if (formState.checkPass && e.target.value != pwd) {
            setIsRightPass(false);
        } else {
            setIsRightPass(true);
        }
        setCPass(e.target.value);
    }

    useEffect(() => {
        if (formState.message.length > 0) {
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
                redirect("/accounts?" + searchParams.toString())
            }
        }
    }, [formState, alertDispatch, router, searchParams])

    return (
        <div className="flex flex-col justify-center items-center my-auto h-full">
            <div className="container border rounded-md w-full max-w-md flex flex-col p-5 space-y-2">
                <form className="w-full flex flex-col" action={formAction}>
                    <div className="h-5"></div>
                    <FormLogo />
                    <p className="font-semibold text-slate-600 text-sm text-center">User Sign Up Form</p>
                    <div className="h-3"></div>
                    <div className="flex flex-row justify-between space-x-1">
                        <div>
                            <Input name="firstname" placeHolder="First Name" label="First Name" required />
                        </div>
                        <div>
                            <Input name="lastname" placeHolder="Last Name" label="Last Name" />
                        </div>
                    </div>
                    <div>
                        <Input name="email" label="Email" type="email" placeHolder="example@gmail.com" required />
                    </div>
                    <div>
                        <Input name="phone" label="Phone" type="text" placeHolder="Phone" required />
                    </div>
                    <div>
                        <Input name="username" label="Username" placeHolder="username" required />
                    </div>
                    <div className="flex flex-row justify-between space-x-1">
                        <div>
                            <Input value={pwd} onChange={onChangePwd} name="password" placeHolder="Password" type="password" label="Password" autoComplete="new-password" required />
                        </div>
                        <div>
                            <Input value={cPass} onChange={onChangeCPass} placeHolder="Confirm Password" type="password" label="Confirm Password" autoComplete="new-password" required />
                            <p hidden={isRightPass} className="text-red-500 text-xs">Password are not equal!</p>
                        </div>
                    </div>
                    <div className="h-1"></div>
                    <p className="text-xs text-slate-600">By creating an account, you agree to the Terms of use and Privacy Policy</p>
                    <FormButton className="btn-primary" message="Sign up" pendingMessage="Signing Up..." />
                </form>
                <p className="text-sm text-slate-600">If you already have an account, please login <Link href={{ pathname: '/login', query: searchParams.toString() }} className="text-blue-400">here</Link></p>
                <p className="text-center text-sm font-bold">Or</p>
                <GoogleSignup />
                <Link className="btn-secondary shadow-sm hover:shadow-lg text-center" href={{ pathname: '/accounts', query: searchParams.toString() }}>Back to my accounts</Link>
                <div className="h-5"></div>
            </div>
        </div>
    );
}