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
}