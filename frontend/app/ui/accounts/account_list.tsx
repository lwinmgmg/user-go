import { getServerUserListCookie } from "@/app/lib/data/cookies/auth_server";
import AccountItem from "./account_item";

export default async function AccountList(){
    const userListStr = getServerUserListCookie();
    const userList: Dictionary<string> = JSON.parse(userListStr || "{}");
    return (
        <div className="container flex flex-col p-5 divide-y-2 bg-slate-50">
            {
                Object.keys(userList).map(key=>(
                    <AccountItem key={key} userCode={key} accessToken={userList[key]} />
                ))
            }
        </div>
    );
}