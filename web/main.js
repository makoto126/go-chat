var conn = new WebSocket('ws://localhost:12345/ws');
var messages = [];

//Javascript's atob to decode base64 doesn't properly decode utf-8 strings, so found this at
//https://stackoverflow.com/questions/30106476/using-javascripts-atob-to-decode-base64-doesnt-properly-decode-utf-8-strings

conn.onmessage = function (e) {
    m = JSON.parse(e.data);
    console.log(m);
    messages.push(m);
}

var app = new Vue({
    el: '#app',
    data: {
        text: '',
        messages: messages
    },
    methods: {
        send: function () {
            conn.send(this.text);
            this.text = '';
        },
        b64DecodeUnicode: function (str) {
            return decodeURIComponent(atob(str).split('').map(function (c) {
                return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
            }).join(''));
        }
    }
})