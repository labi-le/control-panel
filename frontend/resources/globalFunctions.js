let pageBackground = document.querySelector('.page-background');

function formatBytes(bytes, decimals = 2) {
    if (bytes === 0) return '0 байт';
    const k = 1024;
    const dm = decimals < 0 ? 0 : decimals;
    const sizes = ['байт', 'КБ', 'МБ', 'ГБ', 'ТБ', 'ПБ'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / (1024 ** i)).toFixed(dm)) + ' ' + sizes[i];
}

function showNotify(type, content) {
    pageBackground.classList.add('active');

    let windowTemplate = document.querySelector('#specialTemplate-window').content.cloneNode(true);

    switch (type) {
        case "default":
            windowTemplate.querySelector('#titleWindow').innerHTML = "Уведомление";
            break;

        case "task-viewer":
            windowTemplate.querySelector('#titleWindow').innerHTML = "Просмотр выполнения задачи";
            break;

        default:
            windowTemplate.querySelector('#titleWindow').innerHTML = "Уведомление неизвестного типа";
    }

    windowTemplate.querySelector('#contentWindow').innerHTML = content;
    windowTemplate.querySelector('.window').classList.add('active');

    pageBackground.append(windowTemplate);
}

function hideNotify() {
    pageBackground.classList.remove('active');
    pageBackground.innerHTML = '';
}

function setContent(page) {
    if (document.getElementById("template-"+page) === null) {
        showNotify('default', "Неизвестная страница");
    } else {
        document.querySelector(".content").innerHTML = '';
        localStorage.setItem("nowPage", page);
        document.querySelector(".content").append(document.querySelector('#template-'+page).content.cloneNode(true));
    }
}

function changeTheme(toTheme) {
    /*
     * 1st layer - backgrounds
     * 2st layer - any active blocks and contents of him
     * 3st layer - panels (header, footer, etc.)
     * 4st layer - special window (notifications, etc.)
     */
    if (toTheme === null) {
        if (localStorage.getItem("nowTheme") === "black") {
            toTheme = "white";
            localStorage.setItem("nowTheme", "white")
        } else {
            toTheme = "black";
            localStorage.setItem("nowTheme", "black");
        }
    }

    if (toTheme === "white") {
        localStorage.setItem("nowTheme", "white");
        document.documentElement.style.setProperty('--themeFirstLayerColor', '#B4B4B4');
        document.documentElement.style.setProperty('--themeSecondLayerColor', '#B8B8B8');
        document.documentElement.style.setProperty('--themeSecondLayerBlocksColor', '#BCBCBC');
        document.documentElement.style.setProperty('--themeSecondLayerBlocksFieldColor', '#C0C0C0');
        document.documentElement.style.setProperty('--themeThirdLayerColor', '#AEAEAE');
        document.documentElement.style.setProperty('--themeFourthLayerColor', '#B2B2B2');
        document.documentElement.style.setProperty('--themeTextColor', '#000000');
    } else {
        localStorage.setItem("nowTheme", "black");
        document.documentElement.style.setProperty('--themeFirstLayerColor', '#262626');
        document.documentElement.style.setProperty('--themeSecondLayerColor', '#2A2A2A');
        document.documentElement.style.setProperty('--themeSecondLayerBlocksColor', '#2E2E2E');
        document.documentElement.style.setProperty('--themeSecondLayerBlocksFieldColor', '#323232');
        document.documentElement.style.setProperty('--themeThirdLayerColor', '#202020');
        document.documentElement.style.setProperty('--themeFourthLayerColor', '#242424');
        document.documentElement.style.setProperty('--themeTextColor', '#FFFFFF');
    }
}