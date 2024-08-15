import { redirect } from "next/navigation";
import { getServerActiveUserCookie, getServerUserListCookie } from "../lib/data/cookies/auth_server";
import ProfilePage from "../ui/profile/profile_page";
import { Suspense } from "react";

export default function Profile(){
    const activeUser = getServerActiveUserCookie();
    if (!activeUser){
        redirect("/accounts");
    }
    const userDictStr = getServerUserListCookie()
    const userDict = JSON.parse(userDictStr || "{}")
    const userToken = userDict[activeUser];
    if (!userToken){
        redirect("/accounts");
    }
    return (
        <Suspense>
            <ProfilePage accessToken={userToken} userCode={activeUser} />
        </Suspense>
    );
}