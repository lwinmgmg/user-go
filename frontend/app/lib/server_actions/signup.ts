"use server";
import { z } from "zod";
import { parseZodMesg } from "../utils/parse_zod_mesg";
import { setServerOtpCookie } from "../data/cookies/otp_server";
import { setServerAuthCookie } from "../data/cookies/auth_server";

const schema = z.object({
    firstname: z.string().min(1, "Name is required."),
    lastname: z.string(),
    email: z.string().min(1, "Email is required."),
    phone: z.string().min(1, "Phone is required."),
    username: z.string().min(1, "Name is required."),
    password: z.string().min(1, "Password is required")
})

export default async function signup(preState: any, formData: FormData) {
    preState["checkPass"] = true;
    const data = schema.safeParse({
        firstname: formData.get("firstname"),
        lastname: formData.get("lastname"),
        email: formData.get("email"),
        phone: formData.get("phone"),
        username: formData.get("username"),
        password: formData.get("password"),
    });
    if (data.success){
        const resp = await fetch(`${process.env.USER_BACKEND}/user/api/v1/func/user/signup`, {
            method: "POST",
            body: formData,
        })
        if (resp.status !== 200) {
            var message = "Failed to login.";
            if (resp.status === 400){
                const errResp = await resp.json();
                console.log(errResp);
                message = errResp.message;
            }
            preState["message"] = message;
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
    }else if (data.error){
        preState["success"] = false;
        preState["message"] = parseZodMesg(data.error.message)
    }
    return preState
}