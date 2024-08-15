"use client";

import { useFormState } from "react-dom";
import Input from "../components/input";
import FormButton from "../components/form_button";
import FormLogo from "../icons/logo";
import { useContext, useEffect, useRef, useState } from "react";
import { AlertDismisDispatchContext } from "@/app/lib/data/contexts/alert_dismiss_context";
import SelectOtp from "./otp_type_select";
import { redirect, useRouter } from "next/navigation";

export default function OtpForm({
    loginResponse,
    otpAction,
}: {
    loginResponse: LoginRespOtp,
    otpAction: any,
}) {
    const defaultCountDown = 60;
    const alertDispatch = useContext(AlertDismisDispatchContext);
    const router = useRouter();
    const passCode = useRef<HTMLInputElement>(null);
    const [countDown, setCountDown] = useState(defaultCountDown);
    const [formState, formAction] = useFormState(otpAction, {
        message: "",
        access_token: loginResponse.access_token,
        response: loginResponse,
        type: "otp",
        success: false
    })
    const canOtpTypeChange = formState.response.sotp_type != undefined && ["email", "phone", "change_pass"].includes(formState.response.sotp_type);

    useEffect(() => {
        if (formState.type == "resend" && formState.success) {
            setCountDown(defaultCountDown);
        }
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
        if (passCode.current) {
            passCode.current.value = "";
        }
        if (formState.success && formState.type == "otp") {
            redirect("/");
        }
    }, [formState, passCode, alertDispatch])

    useEffect(() => {
        const interval = setInterval(() => {
            setCountDown((val) => val == 0 ? val : val - 1);
        }, 1000)
        return () => {
            clearInterval(interval);
        }
    }, [])

    const download = ()=>{
        const aTag = document.createElement("a")
        aTag.href = 'data:application/octet-stream;base64,' + formState.response.image;
        aTag.download = "qr.png";
        aTag.click();
    }

    return (
        <section className="flex flex-col justify-center items-center my-auto h-full">
            <div className="container border rounded-md w-full max-w-md flex flex-col p-5 space-y-2">
                {
                    formState.response.image ? (
                    <div className="flex flex-col justify-center items-center">
                        <div onClick={download} className="h-52 w-52" style={{
                            backgroundImage: `url(data:image/png;base64,${formState.response.image})`,
                            backgroundSize: "cover",
                            backgroundRepeat: "no-repeat",
                            }}>
                        </div>
                        <p>{formState.response.key}</p>
                    </div>):null
                }
                <FormLogo />
                <p className="font-semibold text-slate-600 text-sm text-center">Otp Auth Form</p>
                <form action={formAction} className="w-full max-w-md flex flex-col space-y-2">
                    <Input innerRef={passCode} name="passcode" label="PassCode" pattern="^[0-9]{6}$" minLength={6} maxLength={6} type="text" placeHolder="eg. 000000" required />
                    <input type="text" name="type" defaultValue={"otp"} hidden />
                    <FormButton className="btn-primary" message="Confirm" pendingMessage="Confirming..." />
                </form>
                <form hidden={formState.response.sotp_type == "auth"} action={formAction} className="w-full max-w-md flex flex-col space-y-2">
                    <input type="text" name="type" defaultValue={"resend"} hidden />
                    <SelectOtp hidden={!(canOtpTypeChange && countDown == 0)} name="sotp_type" defaultValue={formState.response.sotp_type || "email"} values={["email", "phone"]} />
                    <FormButton className="btn-secondary flex-grow" disabled={countDown > 0} message={countDown == 0 ? "Resend Otp" : `Resend Otp in ${countDown}s`} pendingMessage="Resending..." />
                </form>
                <a className="btn-secondary text-center flex-grow" onClick={() => router.back()}>Go Back</a>
            </div>
        </section>
    );
}