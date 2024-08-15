"use client";
import { ChangeEvent, HTMLInputTypeAttribute, LegacyRef } from "react";

export default function Input({
    label,
    placeHolder = "",
    type = "text",
    disabled = false,
    innerRef,
    onChange,
    required = false,
    autoComplete = "",
    maxLength,
    minLength,
    pattern,
    className = "",
    name,
    value
}: {
    label?: string,
    placeHolder?: string
    type?: HTMLInputTypeAttribute,
    disabled?: boolean,
    innerRef?: LegacyRef<HTMLInputElement>,
    required?: boolean,
    autoComplete?: string,
    maxLength?: number,
    minLength?: number,
    pattern?: string,
    className?: string,
    name?: string,
    value?: any,
    onChange?: (e: ChangeEvent<HTMLInputElement>)=>void,
}) {
    const onInnerChange = (e: ChangeEvent<HTMLInputElement>) => {
        if (maxLength && e.target.value.length > maxLength) {
            e.target.value = e.target.value.slice(0, maxLength)
        }
        onChange && onChange(e);
    }
    return (
        <div>
            {
                label ? <span className="block text-sm font-medium text-slate-700">{label}</span> : null
            }
            <input value={value} name={name} type={type} minLength={minLength} maxLength={maxLength} pattern={pattern} onChange={onInnerChange} required={required} disabled={disabled} placeholder={placeHolder} className={"mt-1 block w-full px-3 py-2 bg-white border border-slate-300 rounded-md text-sm shadow-sm placeholder-slate-400 focus:outline-none focus:border-sky-500 focus:ring-1 focus:ring-sky-500 disabled:bg-slate-50 disabled:text-slate-500 disabled:border-slate-200 disabled:shadow-none invalid:border-pink-500 invalid:text-pink-600 focus:invalid:border-pink-500 focus:invalid:ring-pink-500 " + className} ref={innerRef} autoComplete={autoComplete} />
        </div>
    );
}