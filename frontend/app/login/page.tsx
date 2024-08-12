import LoginForm from "@/app/ui/login/LoginForm";

export default function Login({searchParams}:{
    searchParams: any
}){
    return (
        <>
        <LoginForm searchParams={searchParams} />
        </>
    );
}