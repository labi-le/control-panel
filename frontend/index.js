setInterval(function() {
    let x = new XMLHttpRequest();
    x.open("POST", "http://localhost:7000/api/mem/abc", true);
    x.onload = function() {
        let json = JSON.parse(x.responseText);
        console.log(json);
    }
    x.onerror = function () {
        console.log("error fetch data");
    }
    x.send();
}, 5000);