"use server";

import { getServerOtpCookie } from "@/app/lib/data/cookies/otp_server";
import { redirect } from "next/navigation";
import OtpForm from "./otp_form";
import otpLoginAction from "@/app/lib/server_actions/otp_login";

export default async function OtpParent() {
    const activeOtp = getServerOtpCookie();
    if (!activeOtp){
        redirect("/")
    }
    return (
        <OtpForm loginResponse={activeOtp} otpAction={otpLoginAction}/>
    );
}