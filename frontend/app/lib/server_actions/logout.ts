"use server";

import { revalidatePath } from "next/cache";
import { deleteUserFromListServerCookie } from "../data/cookies/auth_server";

export default async function logout(userCode: string, formData: FormData){
    deleteUserFromListServerCookie(userCode);
    revalidatePath("/accounts");
    revalidatePath("/");
}