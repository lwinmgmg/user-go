function capitalizeFirstLetter(str: string) {
    return str.charAt(0).toUpperCase() + str.slice(1);
}

export default function SelectOtp({
    defaultValue,
    name,
    values,
    hidden=false,
    disabled=false,
}:{
    name: string,
    defaultValue: string,
    values: Array<string>,
    hidden?: boolean,
    disabled?: boolean,
}){
    return (
        <div className="text-center" hidden={hidden}>
            <p className="block mb-2 text-xs font-semibold text-gray-900">Choose where to resend otp</p>
            <select name={name} defaultValue={defaultValue} disabled={disabled} className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5">
            {
                values.map(value=>(
                    <option key={value} value={value}>{capitalizeFirstLetter(value)}</option>
                ))
            }
            </select>
        </div>
    );
}