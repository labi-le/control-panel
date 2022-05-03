let pageBackground = document.querySelector('.background');

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

    windowTemplate.querySelector('.window.content').innerHTML = content;

    switch (type) {
        case "default":
            windowTemplate.querySelector('#titleWindow').innerHTML = "Уведомление";
            break;

        case "task-viewer":
            let div = document.createElement('div');
            div.classList.add("window", "footer", "loader");

            windowTemplate.querySelector('#titleWindow').innerHTML = "Задача";
            windowTemplate.querySelector('.window').appendChild(div);
            break;

        default:
            windowTemplate.querySelector('#titleWindow').innerHTML = "Уведомление неизвестного типа";
    }

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
        document.querySelector(".mainView.content").innerHTML = '';
        localStorage.setItem("nowPage", page);
        document.querySelector(".mainView.content").append(document.querySelector('#template-'+page).content.cloneNode(true));
    }
}

function changeTheme(toTheme) {
    /*
     * 1st layer - backgrounds
     * 2nd layer - any active blocks and contents of him
     * 3rd layer - panels (header, footer, etc.)
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
        document.documentElement.style.setProperty('--themeFirstLayerColor', '#AEAEAE');
        document.documentElement.style.setProperty('--themeSecondLayerColor', '#B8B8B8');
        document.documentElement.style.setProperty('--themeSecondLayerBlocksColor', '#BCBCBC');
        document.documentElement.style.setProperty('--themeSecondLayerBlocksFieldColor', '#C0C0C0');
        document.documentElement.style.setProperty('--themeThirdLayerColor', '#B4B4B4');
        document.documentElement.style.setProperty('--themeFourthLayerColor', '#B2B2B2');
        document.documentElement.style.setProperty('--themeFourthLayerColorShadow', '#B6B6B6');
        document.documentElement.style.setProperty('--themeTextColor', '#000000');
    } else {
        localStorage.setItem("nowTheme", "black");
        document.documentElement.style.setProperty('--themeFirstLayerColor', '#262626');
        document.documentElement.style.setProperty('--themeSecondLayerColor', '#2A2A2A');
        document.documentElement.style.setProperty('--themeSecondLayerBlocksColor', '#2E2E2E');
        document.documentElement.style.setProperty('--themeSecondLayerBlocksFieldColor', '#323232');
        document.documentElement.style.setProperty('--themeThirdLayerColor', '#202020');
        document.documentElement.style.setProperty('--themeFourthLayerColor', '#242424');
        document.documentElement.style.setProperty('--themeFourthLayerColorShadow', '#2C2C2C');
        document.documentElement.style.setProperty('--themeTextColor', '#FFFFFF');
    }
}

function createList(inputElementId, idElement, list) {
    if (document.querySelector('#' + idElement).querySelector('.droppedList') !== null) return;
    let listTemplate = document.querySelector('#specialTemplate-droppedList').content.cloneNode(true);
    for (let i = 0; i < list.length; i++) {
        let div = document.createElement('div');
        div.classList.add("button", "text", "listElement");
        div.dataset.key = list[i][0];
        div.dataset.value = list[i][1];
        div.dataset.inputElementId = inputElementId;
        div.onclick = pickElement
        div.innerHTML = list[i][1];
        listTemplate.querySelector('.droppedList.content').appendChild(div);
    }
    document.querySelector('#' + idElement).appendChild(listTemplate);
}

function pickElement() {
    document.getElementById(this.dataset.inputElementId).dataset.value = this.dataset.key;
}

function removeList(parentElement, idElement) {
    if (document.querySelector('#' + idElement) == null) return;
    document.querySelector('#'+parentElement).querySelector("#"+idElement).remove();
}