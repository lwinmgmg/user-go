import clsx from "clsx";
import React from "react";

export default function MenuBtn({children, onClick, selected=false}:{
    children: React.ReactNode,
    selected?: boolean,
    onClick?: any,
}){

    return (
        <a onClick={onClick} href="#" className={clsx(
            "px-4 py-2 text-sm font-medium text-gray-900 bg-slate-50 border-t border-b border-gray-200 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-2 focus:ring-blue-700 focus:text-blue-700 dark:bg-gray-800 dark:border-gray-700 dark:text-white dark:hover:text-white dark:hover:bg-gray-700 dark:focus:ring-blue-500 dark:focus:text-white",
            {
                "bg-gray-100 text-blue-700": selected
            }
        )}>
            {children}
        </a>
    );
}