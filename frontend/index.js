window.onstorage = event => {
    if (event.oldValue === null) return;
    localStorage.setItem(event.key, event.oldValue);
};

setInterval(function () {
    let date = new Date();
    if (localStorage.getItem("changeTheme") === "auto")
        (date.getHours() > 5 && date.getHours() < 19)
            ? changeTheme("white")
            : changeTheme("black");

}, 60000);

function initPanel() {
    if (localStorage.getItem("firstLoad") === null) {
        localStorage.setItem("firstLoad", "0");
        localStorage.setItem("nowTheme", "black");
        localStorage.setItem("nowPage", "dashboard");
        localStorage.setItem("nowLanguage", "ru")
        localStorage.setItem("changeTheme", "auto");
    }
    changeTheme(localStorage.getItem("nowTheme"));
    changeLanguage(localStorage.getItem("nowLanguage"));
    initDashboard();
    setContent("dashboard");
}

function initDashboard() {
    let requestDashboard = new XMLHttpRequest();
    let webSocket = new WebSocket(`ws://${document.location.hostname}:${document.location.port}/ws/dashboard`);
    webSocket.addEventListener("open", function () {
        webSocket.send(JSON.stringify({"path": "/"}));
    });

    webSocket.addEventListener("message", function (event) {
        if (localStorage.getItem("nowPage") !== "dashboard") return;
        let dataDashboard = JSON.parse(event.data)[0];
        // memory info
        document.getElementById("totalMemoryValue").innerHTML = `Всего: ${formatBytes(dataDashboard.mem.total)}`;
        document.getElementById("usedMemoryValue").innerHTML = `Используется: ${formatBytes(dataDashboard.mem.used)}`;
        document.getElementById("freeMemoryValue").innerHTML = `Свободно: ${formatBytes(dataDashboard.mem.free)}`;
        document.getElementById("cachedMemoryValue").innerHTML = `Кэш: ${formatBytes(dataDashboard.mem.cached)}`;
        // CPU info
        document.getElementById("cpuLoadValue").innerHTML = `Нагрузка: ${parseFloat(dataDashboard.cpu_load.load).toFixed(2)}%`;
        // disk info
        document.getElementById("partitionType").innerHTML = `- Тип раздела (${dataDashboard.io.path}): ${dataDashboard.io.fstype}`;
        document.getElementById("totalDiskSizeValue").innerHTML = `-- Всего: ${formatBytes(dataDashboard.io.total)}`;
        document.getElementById("freeDiskSizeValue").innerHTML = `-- Свободно: ${formatBytes(dataDashboard.io.free)}`;
        document.getElementById("usedDiskSizeValue").innerHTML = `-- Используется: ${formatBytes(dataDashboard.io.used)}`;
        document.getElementById("usedDiskSizeInPercentValue").innerHTML = `-- Используется: ${parseFloat(dataDashboard.io.usedPercent).toFixed(2)}%`;
    });

    webSocket.addEventListener("close", function () {
        showWindow("default", "Возникла ошибка при обработке запроса, попробуйте перезагрузить страницу");
        webSocket.close();
    });

    requestDashboard.open("GET", `${document.location.protocol}//${document.location.hostname}:${document.location.port}/api/version`, true);
    requestDashboard.onload = function () {
        document.getElementById("footer-version").innerHTML = `v${JSON.parse(requestDashboard.responseText).version}`;
    };
    requestDashboard.send();
}