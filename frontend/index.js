let requests = setInterval(function() {
    updateStatistics();
}, 1500);
function updateStatistics() {
    let requestDashboard = new XMLHttpRequest();
    requestDashboard.open("POST", "http://localhost:7000/api/dashboard", true);
    requestDashboard.onload = function () {
        let dataDashboard = JSON.parse(requestDashboard.responseText);
        if (dataDashboard.success === false) {
            showWindow("default", "Возникла ошибка при обработке запроса, попробуйте перезагрузить страницу");
            clearInterval(requests);
        }
        // memory info
        document.getElementById("total-memory-size").innerHTML = formatBytes(dataDashboard.data.mem.total);
        document.getElementById("usage-memory-size").innerHTML = formatBytes(dataDashboard.data.mem.total - dataDashboard.data.mem.free);
        document.getElementById("available-memory-size").innerHTML = formatBytes(dataDashboard.data.mem.free);
        // CPU info
        document.getElementById("cpu-load-size").innerHTML = parseFloat(dataDashboard.data.cpu_load.load).toFixed(2) + "%";
        // disk info
        document.querySelector(".disk-info-partitionType").innerHTML = "- Тип раздела ('" + dataDashboard.data.io.path + "'): " + dataDashboard.data.io.fstype;
        document.querySelector(".disk-infoPartitionSize-total").innerHTML = "-- Всего: " + formatBytes(dataDashboard.data.io.total);
        document.querySelector(".disk-infoPartitionSize-free").innerHTML = "-- Свободно: " + formatBytes(dataDashboard.data.io.free);
        document.querySelector(".disk-infoPartitionSize-used").innerHTML = "-- Используется: " + formatBytes(dataDashboard.data.io.used);
        document.querySelector(".disk-infoPartitionSize-usedPercent").innerHTML = "-- Используется (%): " + parseFloat(dataDashboard.data.io.usedPercent).toFixed(2) + "%";
        // update api/panel version
        document.getElementById("footer-version").innerHTML = "v" + dataDashboard.version;
    };
    requestDashboard.onerror = function () {
        showWindow("default", "Возникла ошибка при обработке запроса, попробуйте перезагрузить страницу");
        clearInterval(requests);
    }
    requestDashboard.send(JSON.stringify({"path": "/"}));
}