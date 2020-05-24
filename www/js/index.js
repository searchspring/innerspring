let GS = {
    model: {
        projects: {
            list: [],
            reorder: () => {
                GS.model.projects.list.sort((a, b) => {
                    if (a.pinned === b.pinned) {
                        return a.name.localeCompare(b.name)
                    }
                    return a.pinned ? -1 : 1
                })
            }
        }
    }
}
let $ = (id) => { return document.querySelector(id); }
require([,
    '//apis.google.com/js/api.js',
    '//unpkg.com/mithril/mithril.js',
    'components/login',
    'components/projects',
    'misc/bootstrap'], () => {
        bootstrap((isSignedIn) => {
            if (isSignedIn) {
                GS.user = gapi.auth2.getAuthInstance().currentUser.get()
                GS.token = GS.user.getAuthResponse()['access_token']
                m.mount($('#email'), Login)
                m.mount($('#signout'), Logout)
                m.mount($('#projects'), Projects)
            } else {
                gapi.auth2.getAuthInstance().signIn()
            }
        })
    })