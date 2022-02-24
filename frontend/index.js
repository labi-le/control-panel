setInterval(function() {
    updateStatistics();
}, 1500);
function updateStatistics() {
    let requestMemoryInfo = new XMLHttpRequest();
    requestMemoryInfo.open("POST", "http://localhost:7000/api/memory/info", true);
    requestMemoryInfo.onload = function () {
        let dataMemoryInfo = JSON.parse(requestMemoryInfo.responseText);
        document.getElementById("total-memory-size").innerHTML = formatBytes(dataMemoryInfo.data.total);
        document.getElementById("usage-memory-size").innerHTML = formatBytes(dataMemoryInfo.data.total - dataMemoryInfo.data.free);
        document.getElementById("available-memory-size").innerHTML = formatBytes(dataMemoryInfo.data.free);
        // update api/panel version
        document.getElementById("footer-version").innerHTML = "v" + dataMemoryInfo.version;
    };
    document.getElementById("cpu-load-size").innerHTML = getRandomInt(0, 100) + "%"; // <- заглушка
    requestMemoryInfo.send();
}