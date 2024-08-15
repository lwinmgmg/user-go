"use server";
import { z } from "zod";
import { setServerAuthCookie } from "../data/cookies/auth_server";
import { deleteServerOtpCookie, setServerOtpCookie } from "../data/cookies/otp_server";

const schema = z.object({
    passcode: z.string().min(6, "Passcode must be 6 number long.").max(6, "Passcode must be 6 number long."),
})

async function resendFormType(preState: any, formData: FormData) {
    formData.delete("passcode")
    formData.append("access_token", preState.access_token)
    const resp = await fetch(`${process.env.USER_BACKEND}/user/api/v1/func/user/resend_otp`, {
        method: "POST",
        body: formData
    });
    if (resp.status === 200) {
        const respData: LoginResponse = await resp.json();
        preState["message"] = `Resent to your [${respData.sotp_type}]`;
        preState["success"] = true;
        preState["response"] = respData
        setServerOtpCookie(respData);
    } else {
        preState["message"] = "Failed to resend.";
        preState["success"] = false;
    }
    return preState
}

async function otpFormType(preState: any, formData: FormData) {
    const data = schema.safeParse({
        passcode: formData.get("passcode"),
    })
    formData.append("access_token", preState.access_token)
    if (data.success) {
        const resp = await fetch(`${process.env.USER_BACKEND}/user/api/v1/func/user/otp_auth`, {
            method: "POST",
            body: formData
        });
        if (resp.status === 200) {
            const respData: LoginResponse = await resp.json();
            preState["message"] = "Successfully logged in.";
            preState["success"] = true;
            preState["response"] = respData
            preState["access_token"] = respData.access_token
            setServerAuthCookie(respData.user_id, respData.access_token)
            deleteServerOtpCookie();
        } else {
            preState["message"] = "Failed to login.";
            preState["success"] = false;
        }
    } else if (data.error) {
        preState["success"] = false;
        const messages: Array<{ message: string }> = JSON.parse(data.error.message);
        preState["message"] = messages.map(mesg => mesg.message).join(", ");
    }
    return preState
}

export default async function otpLoginAction(preState: any, formData: FormData): Promise<{
    message?: string,
    response?: LoginResponse,
    access_token: string,
    success: boolean
}> {
    const formType = formData.get("type");
    switch (formType){
        case "otp":
            preState = await otpFormType(preState, formData);
            break
        case "resend":
            preState = await resendFormType(preState, formData);
            break
    }
    formData.delete("type");
    preState["type"] = formType;
    return preState
}