"use client";
import clsx from "clsx";
import { useFormStatus } from "react-dom";

export default function FormButton({
    message,
    pendingMessage,
    className = "",
    disabled = false
}:{
    message: string,
    pendingMessage: string,
    className?: string,
    disabled?: boolean
}){
    const { pending } = useFormStatus();
    return (
        <button className={clsx(className, "shadow-sm hover:shadow-lg")} disabled={pending || disabled}>{pending?pendingMessage:message}</button>
    );
}
