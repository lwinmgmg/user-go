import { ACTIVE_USER_COOKIE, OTP_COOKIE, USER_LIST_COOKIE } from "./enums";
import { getServerCookie, setServerCookie } from "./server_cookie";

export function setServerActiveUserCookie(userCode: string){
    setServerCookie(ACTIVE_USER_COOKIE, userCode);
}

export function setUserListCookie(userCode: string, token: string){
    var userDictStr = getServerCookie(USER_LIST_COOKIE) || "{}";
    var userDict: Dictionary<string> = JSON.parse(userDictStr);
    userDict[userCode] = token;
    setServerCookie(USER_LIST_COOKIE, JSON.stringify(userDict));
}

export function setServerAuthCookie(userCode: string, token: string){
    setServerActiveUserCookie(userCode);
    setUserListCookie(userCode, token);
}
