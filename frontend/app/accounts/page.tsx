import FormLogo from "@/app/ui/icons/logo";
import AddAccount from "../ui/accounts/add_account";
import AccountList from "../ui/accounts/account_list";

export default function Accounts(){
    return (
        <div className="flex flex-col justify-center items-center my-auto h-full">
            <div className="w-full max-w-md text-center rounded-lg border flex flex-col items-center p-5">
                <div className="h-5"></div>
                <FormLogo />
                <p className="font-semibold text-slate-600 text-sm text-center">Account Center</p>
                <AccountList />
                <div className="h-2"></div>
                <AddAccount />
                <div className="h-5"></div>
            </div>
        </div>
    );
}