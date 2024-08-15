"use server";

import getProfile from "@/app/lib/requests/get_profile";
import AccountItemName from "./account_item_name";
import AccountItemLogout from "./account_item_logout";
import { Suspense } from "react";

export default async function AccountItem({
    userCode,
    accessToken,
}:{
    userCode: string,
    accessToken: string
}){
    try{
        const profile = await getProfile(userCode, accessToken);
        return (
            <div className="flex flex-row items-center h-20">
                <Suspense>
                    <AccountItemName profile={profile} />
                </Suspense>
                <AccountItemLogout userCode={profile.userCode} />
            </div>
        );
    }catch(err){
        return null;
    }
}