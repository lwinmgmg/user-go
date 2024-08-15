"use server";

import { revalidatePath } from "next/cache";
import { deleteUserFromListServerCookie } from "../data/cookies/auth_server";

export async function deleteCookieLogout(userCode:string) {
    deleteUserFromListServerCookie(userCode);
    revalidatePath("/accounts");
    revalidatePath("/");
}

export default async function logout(userCode: string, formData: FormData){
    deleteCookieLogout(userCode);
}