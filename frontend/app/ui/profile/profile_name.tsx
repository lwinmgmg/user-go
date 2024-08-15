export default function ProfileName({
    firstName,
    lastName
}:{
    firstName: string,
    lastName?: string
}){
    return (
        <div className="w-full mx-5">
            <h2 className="text-left text-lg font-semibold">{firstName} {lastName}</h2>
        </div>
    );
}