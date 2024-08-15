import LoginForm from "@/app/ui/login/login_form";
import { Suspense } from "react";

export default function Login(){
    return (
        <Suspense>
            <LoginForm />
        </Suspense>
    );
}