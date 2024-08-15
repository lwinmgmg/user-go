export type UserProfile = {
    firstName: string,
    lastName?: string,
    email: string,
    imageUrl?: string,
    userCode: string,
};
export default async function getProfile(accessToken: string): Promise<UserProfile>{
    const headers = new Headers();
    headers.append("Authorization", `Bearer ${accessToken}`)
    const resp = await fetch(`${process.env.USER_BACKEND}/user/api/v1/user/profile`, {
        method: "GET",
        headers: headers
    })
    if (resp.status === 200){
        const data = await resp.json()
        return {
            firstName: data.firstname,
            lastName: data.lastname,
            email: data.email,
            imageUrl: data.image_url,
            userCode: data.user_id
        }
    }
    throw new Error(`status : ${resp.status}`)
}