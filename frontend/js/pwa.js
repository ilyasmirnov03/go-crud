async function registerServiceWorker() {
    if ("serviceWorker" in navigator) {
        const registration = await navigator.serviceWorker.register("/assets/sw.js")
        console.log(registration, 'service worker registration');

        let subscription = await registration.pushManager.getSubscription();
        if (subscription) {
            console.log(JSON.stringify(subscription));
            return;
        }
        subscription = await registration.pushManager.subscribe(({
            userVisibleOnly: true,
            applicationServerKey: 'BL0TX8utP5uYL8Zzqcf6xk7ALK9aTe8RiTqmsjatwbg5tzx651rTXmigD8dLiuv1LO9acQd-83r2M08P5uKHZdQ',
        }))
        console.log(JSON.stringify(subscription));
    }
}
registerServiceWorker()

let installPrompt = null;
const pwaInstallButton = document.getElementById('pwa-install');

window.addEventListener('beforeinstallprompt', (e) => {
    e.preventDefault();
    installPrompt = e;
    pwaInstallButton.removeAttribute('hidden');
    console.log("I can be installed");
});

pwaInstallButton.addEventListener('click', async () => {
    if (!installPrompt) {
        return;
    }
    const result = await installPrompt.prompt();
    console.log('Install prompt result', result);

    Notification.requestPermission().then((res) => {
        console.log(res);
        if (res === "granted") {
            console.log("Notifications granted");
        } else {
            console.log("Notifications refused");
        }
    });

    installPrompt = null;
    pwaInstallButton.setAttribute('hidden', '');
});