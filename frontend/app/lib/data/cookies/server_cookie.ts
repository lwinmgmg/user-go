import { cookies } from "next/headers";

export function setServerCookie(key: string, value: any){
    const cookie = cookies();
    cookie.set(key, value)
}

export function getServerCookie(key: string): string | undefined{
    const cookie = cookies();
    return cookie.get(key)?.value
}

export function deleteServerCookie(key: string){
    const cookie = cookies();
    cookie.delete(key);
}
