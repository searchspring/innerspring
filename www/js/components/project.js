let Project = {
    oninit: () => {
        var urlParams = new URLSearchParams(window.location.search);
        m.request({
            method: 'GET', url: 'api/project/' + urlParams.get('id'), headers: {
                'Authorization': GS.token
            }
        }).then((response) => {
            GS.model.currentProject = response
        })
    },
    view: function (vnode) {
        if (!GS.model.currentProject) {
            return
        }
        return [
            m('h1', { class: 'bg-gray-900 p-2' }, GS.model.currentProject.name),
            m('div', { class: 'flex' }, [
                m('iframe', { class: 'flex-1', src: 'https://stats.kube.searchspring.io/d-solo/nshAi8pmz/jmx-exporter-prometheus?refresh=30s&orgId=1&from=1590286296335&to=1590289896335&var-job=indexer-west&var-pod=All&panelId=8', width: '600px', height: '400px' }),
                m('iframe', { class: 'flex-1', src: 'https://service.us2.sumologic.com/ui/dashboard.html?k=L9QWiaqGyHhljC9OdCgNs2RHOMrqqaCpgbbil6QsqPNyCZg7koQNFc93UMlG&f=&t=r', width: '600px', height: '400px' })
            ])
        ]
    },
}