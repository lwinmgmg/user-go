export default function ProfilePicture({
    imageUrl
}:{
    imageUrl?: string
}){
    return (
        <div className="flex flex-col justify-start relative mb-10">
            <div className="bg-blue-300 h-44 w-full">

            </div>
            <div className="w-24 h-24 overflow-auto border-white rounded-full border-4 absolute -bottom-10 left-5 bg-slate-200">
                <div className="w-full h-full flex flex-row justify-center items-center " style={{
                    backgroundImage: imageUrl? `url(${imageUrl})` : `url(/images/profile-avatar.png)`,
                    backgroundSize: "cover",
                    backgroundRepeat: "no-repeat",
                    }}>
                </div>
            </div>
        </div>
    );
}