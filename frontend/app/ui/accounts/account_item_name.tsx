"use client";

import { UserProfile } from "@/app/lib/requests/get_profile";
import setActiveUser from "@/app/lib/server_actions/active_user";
import { useSearchParams, useRouter } from "next/navigation";
import { useEffect } from "react";
import { useFormState } from "react-dom";

const defaultAvatar = "/images/profile-avatar.png";

export default function AccountItemName({
    profile
}:{
    profile: UserProfile
}){
    const searchParams = useSearchParams();
    const router = useRouter()
    const onSelect = ()=>{

    }
    const [formState, formAction] = useFormState(setActiveUser, {
        userCode: profile.userCode,
        success: false
    });
    useEffect(()=>{
        if (formState.success){
            router.push("/profile?" + searchParams.toString())
        }
    }, [formState, router, searchParams]);
    return  (
        <form className="h-full flex-grow" action={formAction}>
            <button className="h-full w-full rounded-s-lg flex flex-row items-center cursor-pointer hover:bg-slate-200 focus:bg-slate-300">
                <div className="h-16 w-16 ring-1 rounded-full" style={{
                    backgroundImage: `url(${profile.imageUrl ? profile.imageUrl : defaultAvatar})`,
                    backgroundSize: "cover",
                    backgroundRepeat: "no-repeat",
                    }}>
                </div>
                <div className="w-3"></div>
                <div className="grow space-y-1 text-left">
                    <div className="h-full max-w-full">
                        <p className="text-md font-semibold truncate">{profile.firstName} {profile.lastName}</p>
                        <p className="text-sm truncate w-full">{profile.email}</p>
                    </div>
                </div>
            </button>
        </form>
    );
}