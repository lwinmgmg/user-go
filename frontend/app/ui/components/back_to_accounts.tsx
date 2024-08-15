"use client";

import Link from "next/link";
import { useSearchParams } from "next/navigation";

export default function BackToAccounts(){
    const searchParams = useSearchParams();
    return (
        <Link className="btn-secondary shadow-sm hover:shadow-lg text-center" href={{ pathname: '/accounts', query: searchParams.toString() }}>Back to my accounts</Link>
    );
}