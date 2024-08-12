export default function OtpForm({
    userCode,
    code,
    type,
}:{
    userCode: string,
    code: string,
    type: string
}){
    return (
        <form>
            <h1>User : {userCode}</h1>
            <h1>Code : {code}</h1>
            <h1>Type : {type}</h1>
        </form>
    );
}