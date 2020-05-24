async function bootstrap(cb) {
    document.body.style.backgroundImage = `url('images/dark${parseInt(Math.random() * 12)}.jpg')`
    gapi.load('client:auth2', () => {
        var SCOPE = 'https://www.googleapis.com/auth/userinfo.email https://www.googleapis.com/auth/userinfo.profile'
        gapi.client.init({
            'clientId': '706222091746-puj8il6rsncato7iepmargl2pdo6hm22.apps.googleusercontent.com',
            'scope': SCOPE
        }).then(function () {
            gapi.auth2.getAuthInstance().isSignedIn.listen(cb);
            cb(gapi.auth2.getAuthInstance().isSignedIn.get());
        })
    })
}
