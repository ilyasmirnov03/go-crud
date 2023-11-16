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