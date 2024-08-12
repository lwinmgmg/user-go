import { createContext, Dispatch } from 'react';


export type AlertType = "info" | "warn" | "error"

export type AlertDismiss = {
    show: boolean,
    message: string,
    type: AlertType
}

export const alertDismissInitState: AlertDismiss = {
    show: false,
    message: "",
    type: "info"
}

export const AlertDismissContext = createContext<AlertDismiss>(alertDismissInitState);
export const AlertDismisDispatchContext = createContext<Dispatch<{
    type: string,
    state: AlertDismiss
}> | null>(null);

export function alertDismissReducer(state: AlertDismiss, action: {
    type: string,
    state: AlertDismiss
}){
    return action.state
}
