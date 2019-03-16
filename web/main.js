var conn = new WebSocket('ws://localhost:12345/ws');
var messages = [];

conn.onmessage = function (e) {
    m = JSON.parse(e.data);
    console.log(m);
    messages.push(m);
}

var app = new Vue({
    el: '#app',
    data: {
        text: '',
        placeholder: 'Firstly, please input your name.',
        messages: messages
    },
    methods: {
        send: function () {
            conn.send(this.text);
            this.text = '';
            this.placeholder = 'Now you can chat!';
        },
    }
})