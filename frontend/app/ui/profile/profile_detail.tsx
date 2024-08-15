import { UserProfileDetail } from "@/app/lib/requests/get_profile_detail";

export default function ProfileDetail({
    profileDetail
}:{
    profileDetail: UserProfileDetail
}){
    return (
        <div className="w-full ">
            <div className="mx-5 p-5 space-y-2 flex flex-col bg-slate-100 rounded-lg mb-5">
                <p className="text-xs"><span>Email : </span>{profileDetail.email}</p>
                <p className="text-xs"><span>Phone : </span>{profileDetail.phone}</p>
            </div>
        </div>
    );
}