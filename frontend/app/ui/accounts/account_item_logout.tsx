"use server";

import logout from "@/app/lib/server_actions/logout";

export default async function AccountItemLogout({
    userCode
}:{
    userCode: string
}){
    const logoutWithCode = logout.bind(null, userCode);
    return (
        <form className="h-full" action={logoutWithCode}>
            <button className="w-10 h-full rounded-e-md flex flex-col justify-center items-center cursor-pointer hover:bg-slate-200 focus:bg-slate-300">
                <div className=""><svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
                <path strokeLinecap="round" strokeLinejoin="round" d="M15.75 9V5.25A2.25 2.25 0 0 0 13.5 3h-6a2.25 2.25 0 0 0-2.25 2.25v13.5A2.25 2.25 0 0 0 7.5 21h6a2.25 2.25 0 0 0 2.25-2.25V15m3 0 3-3m0 0-3-3m3 3H9" />
                </svg></div>
            </button>
        </form>
    );
}