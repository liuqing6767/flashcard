chrome.contextMenus.create({
    'type':'normal',
    'title':'将单词加入FlashCard',
    'contexts':['selection'],
    'id':'word.flash_card',
    'onclick':addWord
});

function goLogin() {
    window.open('http://fc.liuximu.com', '_blank');
}

function firstWord(word) {
    return word.trim().split(" ")[0]
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
                    window.alert("添加 '" +message + "'失败: " + obj.msg)
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
        return
    }
    chrome.contextMenus.update('word.flash_card',{
        'title':'将单词“'+message+'”加入FCW'
    });
});
