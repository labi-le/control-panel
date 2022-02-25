setInterval(function() {
    updateStatistics();
}, 1500);
function updateStatistics() {
    let requestMemoryInfo = new XMLHttpRequest();
    let requestCPULoad = new XMLHttpRequest();
    requestMemoryInfo.open("POST", "http://localhost:7000/api/memory/info", true);
    requestMemoryInfo.onload = function () {
        let dataMemoryInfo = JSON.parse(requestMemoryInfo.responseText);
        document.getElementById("total-memory-size").innerHTML = formatBytes(dataMemoryInfo.data.total);
        document.getElementById("usage-memory-size").innerHTML = formatBytes(dataMemoryInfo.data.total - dataMemoryInfo.data.free);
        document.getElementById("available-memory-size").innerHTML = formatBytes(dataMemoryInfo.data.free);
        // update api/panel version
        document.getElementById("footer-version").innerHTML = "v" + dataMemoryInfo.version;
    };
    requestCPULoad.open("POST", "http://localhost:7000/api/cpu/load", true);
    requestCPULoad.onload = function () {
        let dataCPULoad = JSON.parse(requestCPULoad.responseText);
        document.getElementById("cpu-load-size").innerHTML = parseFloat(dataCPULoad.data.load).toFixed(2) + "%";
    }
    requestMemoryInfo.send();
    requestCPULoad.send();
}