"use server";
import { revalidatePath } from "next/cache";
import { setServerActiveUserCookie } from "../data/cookies/auth_server";

export default async function setActiveUser(preState: any, formData: FormData){
    setServerActiveUserCookie(preState.userCode);
    revalidatePath("/profile");
    revalidatePath("/");
    preState["success"] = true
    return preState;
}