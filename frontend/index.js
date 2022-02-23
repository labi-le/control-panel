setInterval(function() {
    let x = new XMLHttpRequest();
    x.open("POST", "http://localhost:7000/api/memory/info", true);
    x.onload = function() {
        let json = JSON.parse(x.responseText);
        console.log(json);
        var i = parseInt(Math.floor(Math.log(json.data.available) / Math.log(1024)));
        console.log(Math.round(json.data.available / Math.pow(1024, i), 2));
    }
    x.onerror = function () {
        console.log("error fetch data");
    }
    x.send();
}, 5000);