"use server";

import getProfileDetail from "@/app/lib/requests/get_profile_detail";
import ProfilePicture from "./profile_picture";
import ProfileName from "./profile_name";
import ProfileBtn from "./profile_btns";
import ProfileMenus from "./profile_menus";
import ProfileDetail from "./profile_detail";
import { Suspense } from "react";
import BackToAccounts from "../components/back_to_accounts";
import ProfileSection from "./profile_section";

export default async function ProfilePage({
    userCode,
    accessToken
}:{
    userCode: string,
    accessToken: string
}) {
    const data = await getProfileDetail(accessToken);
    return (
        <section className="flex flex-col justify-center items-center my-auto h-full">
            <div className="container border rounded-lg w-full max-w-md flex flex-col overflow-hidden pb-5">
                <ProfilePicture />
                <ProfileName firstName={data.firstName} lastName={data.lastName} />
                <ProfileSection profileDetail={data} />
                <div className="px-5 flex flex-col">
                    <Suspense>
                        <BackToAccounts />
                    </Suspense>
                </div>
            </div>
        </section>
    );
}