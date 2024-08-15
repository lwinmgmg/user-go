"use client";

import { UserProfileDetail } from "@/app/lib/requests/get_profile_detail";
import ProfileDetail from "./profile_detail";
import ProfileMenus from "./profile_menus";
import { useState } from "react";
import ProfileBtn from "./profile_btns";

export type Menus = "about" | "settings" | "security"


export default function ProfileSection({
    profileDetail
}:{
    profileDetail: UserProfileDetail
}){
    const [selected, setSelected] = useState<Menus>("about")
    return (
        <>
            <ProfileBtn />
            <ProfileMenus selectedMenu={selected} setSelected={setSelected} />
            <ProfileDetail profileDetail={profileDetail} />
        </>
    );
}