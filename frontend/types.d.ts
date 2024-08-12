type LoginResponse = {
    token_type?: string,
    access_token: string,
    user_id: string,
    sotp_type?: string
}
interface Dictionary<T> {
    [Key: string]: T;
}
