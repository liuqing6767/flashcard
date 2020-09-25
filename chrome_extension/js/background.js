var menuID = 'word.flash_card'
var menu = {
    'type':'normal',
    'title':'将单词加入FlashCard',
    'contexts':['selection'],
    'id': menuID,
    'onclick':addWord
}

chrome.contextMenus.create({
    'type':'normal',
    'title':'将单词加入FlashCard',
    'contexts':['selection'],
    'id': menuID,
    'onclick':addWord
});

function goLogin() {
    window.open('http://fc.liuximu.com', '_blank');
}

function firstWord(word) {
    word =  word.trim().split(" ")[0]
    if(!!! word.match(/^[A-Za-z]+$/)) {
        return false
    }

    return word
}

function httpRequest(url, callback){
    var xhr = new XMLHttpRequest();
    xhr.open("GET", url, true);
    xhr.onreadystatechange = function() {
        if (xhr.readyState == 4) {
            callback(xhr.responseText);
        }
    }
    xhr.send();
}

function notification(msg){
    if (!("Notification" in window)) {
        alert(msg)
        return    
    }
    
    // check whether notification permissions have alredy been granted
    if (Notification.permission === "granted") {
        // If it's okay let's create a notification
        new Notification(msg);
    }
    
    // Otherwise, ask the user for permission
    
    if (Notification.permission !== 'denied') {
        Notification.requestPermission(function (permission) {
            // If the user accepts, let's create a notification
            if (permission === "granted") {
                 new Notification(msg);
            }else {
                alert(msg)
            }
        });
    }
}

function addWord(info, tab){
    chrome.cookies.get({
        url: 'http://fc.liuximu.com',
        name: 'token'
    }, function(cookie){
        if(!cookie || !cookie.value) {
            // 未登陆，去登录
            goLogin();
            return;
        }

        message = firstWord(info.selectionText)
        if (!!! message) { // do nothing
            return
        }

        httpRequest('http://fc.liuximu.com/api/card/ext?token=' + cookie.value + '&word=' + message, function(respone) {
            try {
                obj = JSON.parse(respone)
                if (obj.errno !== 0) {
                    notification("添加 '" +message + "'失败: " + obj.msg)
                }else {
                    notification("添加 '" +message + "'成功") 
                }
            } catch (error) {
                console.log(error);
            }   
        })
    });
}

chrome.runtime.onMessage.addListener(function(message, sender, sendResponse){
    message = firstWord(message)
    if (!!! message) {

    chrome.contextMenus.update(menuID,{
        'title':'FCW Disabled: 只能添加英文单词'
    });
        return
    }
    chrome.contextMenus.update(menuID,{
        'title':'FCW: 添加单词“'+message+'”'
    });
});
