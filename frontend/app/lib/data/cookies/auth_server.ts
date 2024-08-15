import { ACTIVE_USER_COOKIE, USER_LIST_COOKIE } from "./enums";
import { deleteServerCookie, getServerCookie, setServerCookie } from "./server_cookie";

export function setServerActiveUserCookie(userCode: string){
    setServerCookie(ACTIVE_USER_COOKIE, userCode);
}

export function getServerActiveUserCookie(){
    return getServerCookie(ACTIVE_USER_COOKIE);
}

export function deleteActiveUserServerCookie(){
    deleteServerCookie(ACTIVE_USER_COOKIE);
}

export function setUserListCookie(userCode: string, token: string){
    var userDictStr = getServerCookie(USER_LIST_COOKIE) || "{}";
    var userDict: Dictionary<string> = JSON.parse(userDictStr);
    userDict[userCode] = token;
    setServerCookie(USER_LIST_COOKIE, JSON.stringify(userDict));
}

export function getServerUserListCookie(){
    return getServerCookie(USER_LIST_COOKIE);
}

export function deleteUserFromListServerCookie(userCode: string){
    const activeUser = getServerActiveUserCookie();
    var userDictStr = getServerCookie(USER_LIST_COOKIE) || "{}";
    var userDict: Dictionary<string> = JSON.parse(userDictStr);
    delete userDict[userCode]
    setServerCookie(USER_LIST_COOKIE, JSON.stringify(userDict));
    if (activeUser === userCode){
        const userList = Object.keys(userDict);
        if (userList.length > 0){
            setServerActiveUserCookie(userList[0]);
        }else{
            deleteActiveUserServerCookie();
        }
    }
}

export function setServerAuthCookie(userCode: string, token: string){
    setServerActiveUserCookie(userCode);
    setUserListCookie(userCode, token);
}
