let pageBackground = document.querySelector(".background");

function formatBytes(bytes, decimals = 2) {
    if (bytes === 0) return "0 B";

    const k = 1024;
    const dm = decimals < 0 ? 0 : decimals;
    const sizes = ["B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];

    const i = Math.floor(Math.log(bytes) / Math.log(k));

    return parseFloat((bytes / Math.pow(k, i)).toFixed(dm)) + " " + sizes[i];
}

function showWindow(type, content) {
    pageBackground.classList.add("active");

    let windowTemplate = document.querySelector("#specialTemplate-window").content.cloneNode(true);

    windowTemplate.querySelector(".blockContent").innerHTML = content;

    switch (type) {

        case "showTask":
            let div = document.createElement("div");
            div.classList.add("window", "footer", "loader");

            windowTemplate.querySelector("#titleWindow").innerHTML = getTextByTemplatePhrase(localStorage.getItem("nowLanguage"), "window-title-showTask");
            windowTemplate.querySelector(".window").appendChild(div);
            break;

        default:
            windowTemplate.querySelector("#titleWindow").innerHTML = getTextByTemplatePhrase(localStorage.getItem("nowLanguage"), "window-title-defaultNotification");
    }

    windowTemplate.querySelector(".window").classList.add("active");

    pageBackground.append(windowTemplate);
}

function hideWindow() {
    pageBackground.classList.remove("active");
    pageBackground.innerHTML = "";
}

function setContent(page) {
    if (document.getElementById(`template-${page}`) === null) {
        showWindow(null, getTextByTemplatePhrase(localStorage.getItem("nowLanguage"), "window-content-undefinedPage"));
    } else {
        localStorage.setItem("nowPage", page);
        document.getElementById("mainApp").innerHTML = "";
        document.getElementById("mainApp").append(document.querySelector(`#template-${page}`).content.cloneNode(true));
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
        document.documentElement.style.setProperty("--themeBlockLighterColor", "#BCBCBC");
        document.documentElement.style.setProperty("--themeBlockDefaultColor", "#B8B8B8");
        document.documentElement.style.setProperty("--themeBlockLittleShadowyColor", "#AEAEAE");
        document.documentElement.style.setProperty("--themeBlockShadowyColor", "#B4B4B4");
        document.documentElement.style.setProperty("--themeListBlockColor", "#B6B6B6");
        document.documentElement.style.setProperty("--themeListElementShadow", "#B0B0B0");
        document.documentElement.style.setProperty("--themeButtonBackground", "#B2B2B2");
        document.documentElement.style.setProperty("--themeInputBackground", "#C0C0C0");
        document.documentElement.style.setProperty("--themeTextColor", "#000000");
        document.documentElement.style.setProperty("--themePlaceholderColor", "#5E5E5E");
    } else {
        localStorage.setItem("nowTheme", "black");
        document.documentElement.style.setProperty("--themeBlockLighterColor", "#2E2E2E");
        document.documentElement.style.setProperty("--themeBlockDefaultColor", "#2A2A2A");
        document.documentElement.style.setProperty("--themeBlockLittleShadowyColor", "#262626");
        document.documentElement.style.setProperty("--themeBlockShadowyColor", "#222222");
        document.documentElement.style.setProperty("--themeListBlockColor", "#2C2C2C");
        document.documentElement.style.setProperty("--themeListElementShadow", "#282828");
        document.documentElement.style.setProperty("--themeButtonBackground", "#242424");
        document.documentElement.style.setProperty("--themeInputBackground", "#323232");
        document.documentElement.style.setProperty("--themeTextColor", "#FFFFFF");
        document.documentElement.style.setProperty("--themePlaceholderColor", "#A0A0A0");
    }
}

function createList(elementId, elementType, parentElementId, list) {
    if (document.querySelector(`#${parentElementId}`).querySelector(".block.droppedList") !== null) return;
    document.querySelector(`#${parentElementId}`).onmouseleave = special_eventHandler;
    let listTemplate = document.querySelector("#specialTemplate-droppedList").content.cloneNode(true);
    for (let i = 0; i < list.length; i++) {
        let div = document.createElement("div");
        div.classList.add("button", "text", "listElement");
        div.dataset.key = list[i][0];
        div.dataset.value = list[i][1];
        div.dataset.elementId = elementId;
        div.dataset.elementType = elementType;
        div.onclick = pickElement
        div.innerHTML = list[i][1];
        listTemplate.querySelector(".blockContent").append(div);
    }
    document.querySelector(`#${parentElementId}`).append(listTemplate);
}

function pickElement() {
    let element = document.getElementById(this.dataset.elementId);
    element.dataset.value = this.dataset.key;
    switch (this.dataset.elementType) {
        case "input":
            element.value = this.dataset.value;
            break;
        case "button":
            element.innerHTML = this.dataset.value;
            break;
    }
    removeList();
}

function removeList() {
    //document.querySelector(".droppedList").remove();
}

function changeLanguage(lang) {
    let languages = document.querySelector("#specialTemplate-languageTemplates").content.cloneNode(true);
    let elements = document.querySelectorAll("title, template, div.text, input");
    if (languages.querySelector(`#${lang}-language`) === null) return;
    elements.forEach((element) => {
        if (element.content !== undefined) {
            element.content.querySelectorAll("div.text, input").forEach((templateElement) => {
                if (templateElement.dataset.templatePhrase !== undefined) setTextInElement(templateElement, templateElement.dataset.typeTemplatePhrase);
            });
        } else {
            if (element.dataset.templatePhrase !== undefined) setTextInElement(element, element.dataset.typeTemplatePhrase);
        }
    });

    localStorage.setItem("nowLanguage", lang);

    function setTextInElement(element, type) {
        let elementPhraseTemplate = element.dataset.templatePhrase.match(/__(.*)__/)[1];

        switch (type) {
            case "0":
                element.innerHTML = languages.querySelector(`#${lang}-language`).querySelector(`#${lang}-${languages.querySelector(`#${lang}-language`).querySelector(`#${lang}-${elementPhraseTemplate}`) !== null ? elementPhraseTemplate : "bug"}`).innerHTML;
                break;
            case "1":
                element.placeholder = languages.querySelector(`#${lang}-language`).querySelector(`#${lang}-${languages.querySelector(`#${lang}-language`).querySelector(`#${lang}-${elementPhraseTemplate}`) !== null ? elementPhraseTemplate : "bug"}`).innerHTML;
        }

    }
}

function getTextByTemplatePhrase(lang, phrase) {
    let languages = document.querySelector("#specialTemplate-languageTemplates").content.cloneNode(true);

    return languages.querySelector(`#${lang}-language`).querySelector(`#${lang}-${languages.querySelector(`#${lang}-language`).querySelector(`#${lang}-${phrase}`) !== null ? phrase : "bug"}`).innerHTML;
}

function updateSettings(listOfElementsId) {
    listOfElementsId.forEach((elementId) => {
        if (document.getElementById(elementId).dataset.value === undefined) return;
        switch (document.getElementById(elementId).dataset.parameter) {
            case "port":
                // ...
                break;
            case "language":
                changeLanguage(document.getElementById(elementId).dataset.value);
                break;
        }
        // TODO: update settings in server-side
    });
}

function replaceClass(element, replacement, replace) {
    let containsReplacement = element.classList.contains(replacement)
    element.className.replace(containsReplacement ? replacement : replace, containsReplacement ? replace : replacement)
}

function special_eventHandler(event) {
    switch (event.type) {
        case "mouseleave":
            removeList();
            replaceClass(event.path[1], "opened", "closed");
            break;
    }
}