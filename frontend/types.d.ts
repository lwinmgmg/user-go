type LoginResponse = {
    token_type?: string,
    access_token: string,
    user_id: string,
    sotp_type?: string
}

interface LoginRespOtp extends LoginResponse{
    image?: string,
    key?: string
}

interface Dictionary<T> {
    [Key: string]: T;
}

interface AnyDict<T1, T2> {
    [key: T1]: T2;
}
