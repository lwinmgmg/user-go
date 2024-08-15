import { deleteCookieLogout } from "../server_actions/logout";

export type UserProfileDetail = {
    firstName: string,
    lastName?: string,
    email: string,
    phone: string,
    imageUrl?: string,
    userCode: string,
    isEmail: boolean,
    isPhone: boolean,
    isAuth: boolean,
    is2Fa: boolean,
};

export default async function getProfileDetail(userCode: string, accessToken: string): Promise<UserProfileDetail>{
    const headers = new Headers();
    headers.append("Authorization", `Bearer ${accessToken}`)
    const resp = await fetch(`${process.env.USER_BACKEND}/user/api/v1/user/profile_detail`, {
        method: "GET",
        headers: headers
    })
    if (resp.status === 200){
        const data = await resp.json()
        return {
            firstName: data.firstname,
            lastName: data.lastname,
            email: data.email,
            phone: data.phone,
            imageUrl: data.image_url,
            userCode: data.user_id,
            isEmail: data.is_email,
            isPhone: data.is_phone,
            isAuth: data.is_auth,
            is2Fa: data.is_2fa
        }
    }else if (resp.status === 401){
        deleteCookieLogout(userCode);
    }
    throw new Error(`status : ${resp.status}`)
}