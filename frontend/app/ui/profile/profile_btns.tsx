import IconButton from "../components/icon_btn";
import PencilIcon from "../components/icons/pencil_icon";

export default function ProfileBtn(){
    return (
        <div className="mx-5">
            <IconButton icon={<PencilIcon className="size-3.5" />}>Edit Profile</IconButton>
        </div>
    );
}