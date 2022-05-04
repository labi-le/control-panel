window.onstorage = event => {
    if (event.oldValue === null) return;
    localStorage.setItem(event.key, event.oldValue);
};

setInterval(function() {
    let date = new Date();
    if (localStorage.getItem("changeTheme") === "auto") (date.getHours() >= 19 && date.getHours() <= 5) ? changeTheme("black") : changeTheme("white");
}, 60000);

let requests = setInterval(function() {
    if (localStorage.getItem("nowPage") === "dashboard") updateStatistics();
}, 1000);

function initLocalStorage() {
    if (localStorage.getItem("firstLoad") === null) {
        localStorage.setItem("firstLoad", "0");
        localStorage.setItem("nowTheme", "black");
        localStorage.setItem("nowPage", "dashboard");
        localStorage.setItem("changeTheme", "auto");
    }
    changeTheme(localStorage.getItem("nowTheme"));
}

function updateStatistics() {
    let requestDashboard = new XMLHttpRequest();
    requestDashboard.open("POST", document.location.protocol + "//" + document.location.hostname + ":" + document.location.port + "/api/dashboard", true);
    requestDashboard.onload = function () {
        let dataDashboard = JSON.parse(requestDashboard.responseText);
        if (dataDashboard.success === false) {
            showNotify("default", "Возникла ошибка при обработке запроса, попробуйте перезагрузить страницу");
            clearInterval(requests);
        }

        // memory info
        document.getElementById("totalMemVal").innerHTML = formatBytes(dataDashboard.mem.total);
        document.getElementById("usedMemVal").innerHTML = formatBytes(dataDashboard.mem.used);
        document.getElementById("freeMemVal").innerHTML = formatBytes(dataDashboard.mem.free);
        document.getElementById("cachedMemVal").innerHTML = formatBytes(dataDashboard.mem.cached);
        // CPU info
        document.getElementById("cpuLoadValue").innerHTML = parseFloat(dataDashboard.cpu_load.load).toFixed(2) + "%";
        // disk info
        document.getElementById("partitionType").innerHTML = "- Тип раздела ('" + dataDashboard.io.path + "'): " + dataDashboard.io.fstype;
        document.getElementById("totalDiskSizeValue").innerHTML = "-- Всего: " + formatBytes(dataDashboard.io.total);
        document.getElementById("freeDiskSizeValue").innerHTML = "-- Свободно: " + formatBytes(dataDashboard.io.free);
        document.getElementById("usedDiskSizeValue").innerHTML = "-- Используется: " + formatBytes(dataDashboard.io.used);
        document.getElementById("usedDiskSizeInPercentValue").innerHTML = "-- Используется (%): " + parseFloat(dataDashboard.io.usedPercent).toFixed(2) + "%";
        // update api/panel version
        document.getElementById("footer-version").innerHTML = "v" + requestDashboard.getResponseHeader("Version");
    };
    requestDashboard.onerror = function () {
        showNotify("default", "Возникла ошибка при обработке запроса, попробуйте перезагрузить страницу");
        clearInterval(requests);
    }
    requestDashboard.send(JSON.stringify({"path": "/"}));
}