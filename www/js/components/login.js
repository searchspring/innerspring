let Login = {
    view: function (vnode) {
        return GS.user.getBasicProfile().getEmail()
    }
}
let Logout = {
    view: function (vnode) {
        return m('span', {onclick: ()=>{
            gapi.auth2.getAuthInstance().signOut()
        }}, 'sign out')
    },
}