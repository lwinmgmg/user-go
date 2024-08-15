import { OTP_COOKIE } from "./enums";
import { deleteServerCookie, getServerCookie, setServerCookie } from "./server_cookie";

export function setServerOtpCookie(data: LoginResponse){
    setServerCookie(OTP_COOKIE, JSON.stringify(data));
}

export function getServerOtpCookie(): LoginResponse | undefined{
    const otpStr = getServerCookie(OTP_COOKIE);
    if (otpStr){
        return JSON.parse(otpStr) as LoginResponse
    }
    return undefined
}

export function deleteServerOtpCookie(){
    deleteServerCookie(OTP_COOKIE);
}
