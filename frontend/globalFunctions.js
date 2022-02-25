let windowBackground = document.querySelector('.window-background');
let window_ = document.querySelector('.window');

function formatBytes(bytes, decimals = 2) {
    if (bytes === 0) return '0 Bytes';
    const k = 1024;
    const dm = decimals < 0 ? 0 : decimals;
    const sizes = ['байт', 'КБ', 'МБ', 'ГБ', 'ТБ', 'ПБ'];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / (1024 ** i)).toFixed(dm)) + ' ' + sizes[i];
}
function showWindow(content) {
    windowBackground.classList.add('active');
    window_.classList.add('active');
    document.querySelector('.window-content').innerHTML = content;
}
function hideWindow() {
    windowBackground.classList.remove('active');
    window_.classList.remove('active');
    document.querySelector('.window-content').innerHTML = "";
}