setInterval(function() {
    let x = new XMLHttpRequest();
    x.open("POST", "http://localhost:7000/api/memory/info", true);
    x.onload = function() {
        let json = JSON.parse(x.responseText);
        document.getElementById("usage-memory-size").innerHTML = /*formatBytes(json.data.used + json.data.cached) <- с округлением до 2х после запятой */ (json.data.used + json.data.cached) / (1000 ** 3) + " GB" /* <- без округления*/;
        document.getElementById("available-memory-size").innerHTML = formatBytes(json.data.free);
        //console.log(json);
        //console.log("Доступно: " + formatBytes(json.data.free) + "\n" + "Используется: " + formatBytes(json.data.used + json.data.cached) + "\n" + "Всего: " + formatBytes(json.data.total));
    }
    x.onerror = function () {
        console.log("error fetch data");
    }
    x.send();
}, 5000);
function formatBytes(bytes, decimals = 2) {
    if (bytes === 0) return '0 Bytes';
    const k = 1024;
    const dm = decimals < 0 ? 0 : decimals;
    const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + ' ' + sizes[i];
}