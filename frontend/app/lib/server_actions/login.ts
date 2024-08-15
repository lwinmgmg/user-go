"use server";
import { z } from "zod";
import { setServerAuthCookie } from "@/app/lib/data/cookies/auth_server";
import { setServerOtpCookie } from "@/app/lib/data/cookies/otp_server";

const schema = z.object({
    username: z.string().min(1, "Name is required."),
    password: z.string().min(1, "Password is required")
})

export default async function loginAction(preState: any, formData: FormData): Promise<{
    message?: string,
    response?: LoginResponse,
    success: boolean
}> {
    const data = schema.safeParse({
        username: formData.get("username"),
        password: formData.get("password"),
    })
    if (data.success) {
        const resp = await fetch(`${process.env.USER_BACKEND}/user/api/v1/func/user/login`, {
            method: "POST",
            body: formData
        });
        if (resp.status !== 200) {
            preState["message"] = "Failed to login.";
            preState["success"] = false;
        } else {
            const respData: LoginResponse = await resp.json();
            preState["message"] = "Successfully logged in.";
            preState["success"] = true;
            preState["response"] = respData
            if (respData.token_type == "Otp") {
                setServerOtpCookie(respData);
            } else if (respData.token_type == "Bearer") {
                setServerAuthCookie(respData.user_id, respData.access_token);
            }
        }
    } else if (data.error) {
        preState["success"] = false;
        const messages: Array<{ message: string }> = JSON.parse(data.error.message);
        preState["message"] = messages.map(mesg => mesg.message).join(", ");
    }
    return preState
}
