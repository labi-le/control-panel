let windowBackground = document.querySelector('.window-background');

function formatBytes(bytes, decimals = 2) {
    if (bytes === 0) return '0 Bytes';
    const k = 1024;
    const dm = decimals < 0 ? 0 : decimals;
    const sizes = ['байт', 'КБ', 'МБ', 'ГБ', 'ТБ', 'ПБ'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / (1024 ** i)).toFixed(dm)) + ' ' + sizes[i];
}

function generateWindow(type, content) {
    let window_ = document.createElement('div');
    window_.className = "window";
    let window_header = document.createElement('div');
    window_header.className = "window-header";
    let window_title = document.createElement('div');
    window_title.className = "window-title";
    let window_close = document.createElement('div');
    window_close.className = "window-close";
    window_close.innerHTML = "Закрыть";
    window_close.onclick = hideWindow;
    let window_body = document.createElement('div');
    window_body.className = "window-body";
    let window_content = document.createElement('div');
    window_content.className = "window-content";
    if (type === "default") {
        window_title.innerHTML = "Уведомление";
        window_content.innerHTML = content;
    } else if (type === "task-viewer") {
        window_title.innerHTML = "Просмотр выполнения задачи";
        let window_content_text = document.createElement('div');
        window_content_text.className = "window-content-text";
        window_content_text.innerHTML = content;
        let window_content_loader = document.createElement('div');
        window_content_loader.className = "window-loader";
        window_content.append(window_content_text, window_content_loader);
    } else {
        window_title.innerHTML = "Уведомление неизвестного типа";
        window_content.innerHTML = content;
    }
    window_header.append(window_title, window_close);
    window_body.append(window_content)
    window_.append(window_header, window_body);
    window_.classList.add('active');
    return window_;
}

function showWindow(type, content) {
    windowBackground.classList.add('active');
    windowBackground.appendChild(generateWindow(type, content));
}

function hideWindow() {
    windowBackground.classList.remove('active');
    this.parentElement.parentElement.remove();
}