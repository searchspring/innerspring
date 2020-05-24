let Projects = {
    oninit: () => {
        m.request({
            method: 'GET', url: 'api/projects', headers: {
                'Authorization': GS.token
            }
        }).then((response) => {
            GS.model.projects.list = response
            if (!localStorage.getItem('pins')) {
                localStorage.setItem('pins', JSON.stringify({}))
            }
            let pins = JSON.parse(localStorage.getItem('pins'))
            for (let project of GS.model.projects.list) {
                project.pinned = pins[project.id] ? true : false
            }
            GS.model.projects.reorder()
        })
    },
    view: function (vnode) {
        let k = 0
        let projectList = GS.model.projects.list.map(function (project) {
            return m('.user-list-item', {
                key: project.id,
                class: 'px-2 py-1 ' + (k++ % 2 === 0 ? 'bg-gray-500' : ''),
                onclick: () => {
                    self.location = 'project.html?id=' + project.id
                }
            }, [m('span', {
                onclick: (e) => {
                    let pins = JSON.parse(localStorage.getItem('pins'))
                    project.pinned = !project.pinned
                    pins[project.id] = project.pinned
                    GS.model.projects.reorder()
                    localStorage.setItem('pins', JSON.stringify(pins))
                    return false
                }, class: 'mr-4'
            }, (project.pinned ? 'unpin' : 'pin')),
            project.name
            ])
        })
        return [
            m('h1', { class: 'bg-gray-900 p-2' }, 'Projects'),
            projectList
        ]
    },
}