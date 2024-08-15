import MenuBtn from "@/app/ui/profile/menu_btn";
import { Menus } from "@/app/ui/profile/profile_section";
import ProfileBtn from "./profile_btns";

export default function ProfileMenus({
    selectedMenu,
    setSelected
}:{
    selectedMenu: Menus,
    setSelected: any
}){
    const onClick = (selectedItem: Menus)=>{
        setSelected(selectedItem);
    }
    return (
        <div className="inline-flex overflow-auto rounded-lg shadow-sm mt-1 mx-5">
            <MenuBtn onClick={()=>onClick("about")} selected={selectedMenu === "about"}>About</MenuBtn>
            <MenuBtn onClick={()=>onClick("settings")} selected={selectedMenu === "settings"}>Settings</MenuBtn>
            <MenuBtn onClick={()=>onClick("security")} selected={selectedMenu === "security"}>Security</MenuBtn>
        </div>
    );
}