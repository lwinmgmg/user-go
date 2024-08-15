import React from "react";

export default function IconButton({
    icon,
    children
}:{
    icon?: React.ReactNode,
    children?: React.ReactNode
}){
    return (
        <button type="button" className="px-3 py-2 text-xs font-medium text-center inline-flex items-center text-white bg-blue-700 rounded-lg hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
        {icon}
        {children}
        </button>
    );
}