"use client";
import { useFormStatus } from "react-dom";

export default function FormButton(){
    const { pending } = useFormStatus();
    return (
        <button className="btn-primary shadow-sm hover:shadow-lg" disabled={pending}>{pending?"Logging in...":"Login"}</button>
    );
}
