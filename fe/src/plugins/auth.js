function GetUser() {
    return window.localStorage.getItem("email")
}

function SetUser(user) {
    window.localStorage.setItem("email", user)
}

function ClearUser() {
    window.localStorage.clear("email")
}

function CheckUser() {
    return GetUser()
}

export {
    GetUser,
    SetUser,
    CheckUser,
    ClearUser
}