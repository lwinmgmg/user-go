import { Suspense } from "react";
import SignupForm from "@/app/ui/signup/signup_form";

export default function Signup(){
    return (
        <Suspense>
            <SignupForm />
        </Suspense>
    );
}