"use server";

import { getServerOtpCookie } from "@/app/lib/data/cookies/otp_server";
import { redirect } from "next/navigation";
import OtpForm from "./otp_form";

export default async function OtpParent() {
    const activeOtp = getServerOtpCookie();
    if (!activeOtp){
        redirect("/")
    }
    return (
        <OtpForm userCode={activeOtp.user_id} code={activeOtp.access_token} type={activeOtp.sotp_type || ""}  />
    );
}