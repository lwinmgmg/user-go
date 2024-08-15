export function parseZodMesg(mesg: string){
    const messages: Array<{ message: string }> = JSON.parse(mesg);
    return messages.map(mesg => mesg.message).join(", ");
}