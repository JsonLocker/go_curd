
// 基于Canvas浏览器指纹
function webFinger(txt = location.host) {
    var canvas = document.createElement('canvas');
    var ctx = canvas.getContext("2d");
    ctx.textBaseline = "top";
    ctx.font = "14px 'Arial'";
    ctx.textBaseline = "tencent";
    ctx.fillStyle = "#f60";
    ctx.fillRect(125, 1, 62, 20);
    ctx.fillStyle = "#069";
    ctx.fillText(txt, 2, 15);
    ctx.fillStyle = "rgba(102, 204, 0, 0.7)";
    ctx.fillText(txt, 4, 17);
    var b64 = canvas.toDataURL().replace("data:image/png;base64,", "");
    var bin = atob(b64);
    let s = bin.slice(-16, -12)
    var i, l, crc = '', n;
    s += '';
    for (i = 0, l = s.length; i < l; i++) {
        n = s.charCodeAt(i).toString(16);
        crc += n.length < 2 ? '0' + n : n;
    }
    return crc;
}

// 获取url 数据
function param(id) {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get(id);
}

// 缓存, 默认1小时过期
function cache(key, value, hour) {
    var expiryDate = new Date()
    let expire = expiryDate.getTime() + hour * 1000 * 60 * 60
    // 存
    if (value !== undefined) {
        localStorage.setItem(key, JSON.stringify({ value: value, expiryTime: expire }));
        return value
    }
    // 读
    var storedItem = JSON.parse(localStorage.getItem(key));
    if (!storedItem) { return null; }
    if (new Date().getTime() > storedItem.expiryTime) {
        localStorage.removeItem(key);
        return null;
    }
    return storedItem.value;
}

// post 数据 postData : {name: "lucy"}
async function curl(url = "", postData = {}, isJsonType = false) {
    let formBody
    let header = {}
    if (isJsonType == false) {
        formBody = [];
        for (var property in postData) {
            let encodedKey = encodeURIComponent(property);
            let encodedValue = encodeURIComponent(postData[property]);
            formBody.push(encodedKey + "=" + encodedValue);
        }
        formBody = formBody.join("&");
        header["Content-Type"] = "application/x-www-form-urlencoded"
        header["X-Requested-With"] = "XMLHttpRequest"
    } else {
        formBody = JSON.stringify(postData)
        header["Content-Type"] = "application/json"
    }
    const response = await fetch(url, {
        method: "POST", // *GET, POST, PUT, DELETE, etc.
        mode: "cors", // no-cors, *cors, same-origin
        cache: "no-cache", // *default, no-cache, reload, force-cache, only-if-cached
        credentials: "same-origin", // include, *same-origin, omit
        headers: header,
        redirect: "follow", // manual, *follow, error
        referrerPolicy: "no-referrer", // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
        body: formBody, // body data type must match "Content-Type" header
    });
    return response.json();
}

// 监听系统滚动到底部自动执行函数
function windowScroll(callBackfunc = "nextPage") {
    window.onscroll = function () {
        var scrollHeight = document.documentElement.scrollHeight;
        var totalScroll = window.scrollY + window.innerHeight;
        if (totalScroll >= scrollHeight) {
            eval(callBackfunc + "()");
        }
    };
}

// 把页面的table导出excel
function exportTableToExcel(tableID, filename = '') {
    var downloadLink;
    var dataType = 'application/vnd.ms-excel';
    var tableSelect = document.getElementById(tableID);
    var tableHTML = tableSelect.outerHTML.replace(/ /g, '%20');
    filename = filename ? filename + '.xls' : 'excel_data.xls';
    downloadLink = document.createElement("a");
    document.body.appendChild(downloadLink);
    if (navigator.msSaveOrOpenBlob) {
        var blob = new Blob(['\ufeff', tableHTML], {
            type: dataType
        });
        navigator.msSaveOrOpenBlob(blob, filename);
    } else {
        downloadLink.href = 'data:' + dataType + ', ' + tableHTML;
        downloadLink.download = filename;
        downloadLink.click();
    }
}

// 修改主题颜色，可覆盖bootstrap默认主题颜色
function themeColor(newColor = null, themeColor = "--bs-guangai") {
    newColor = newColor === null ? '#36bc24' : newColor
    const root = document.querySelector(':root');
    const color = getComputedStyle(root).getPropertyValue('--bs-guangai').trim();
    root.style.setProperty('--bs-guangai', newColor);
}

// 判断是否在微信浏览器
function isWeChat(){
    var ua = window.navigator.userAgent.toLowerCase();
    if(ua.match(/MicroMessenger/i) == 'micromessenger'){
        return true;
    }else{
        return false;
    }
}

// 手机端监听回车,默认关闭焦点
document.addEventListener('keyup', function(event) {
    if (event.keyCode === 13) {
      let focusedElement = document.activeElement;
      if (focusedElement.tagName === 'INPUT') {
        focusedElement.blur();
      }
    }
  });