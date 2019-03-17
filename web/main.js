var conn = new WebSocket('ws://localhost:12345/ws');
var messages = [];
conn.onerror = function () {
    alert('connection failed!');
}
conn.onmessage = function (e) {
    m = JSON.parse(e.data);
    messages.push(m);
}

var app = new Vue({
    el: '#app',

    data: {
        messages: messages,
        text: '',
        placeholder: 'Firstly, please input your name.',
    },

    methods: {
        send: function () {
            conn.send(this.text);
            this.text = '';
            this.placeholder = 'Now you can chat!';
        },
    }
})


var input = document.getElementById("input");

// Execute a function when the user releases a key on the keyboard
input.addEventListener("keyup", function(event) {
  // Number 13 is the "Enter" key on the keyboard
  if (event.keyCode === 13) {
    // Cancel the default action, if needed
    event.preventDefault();
    // Trigger the button element with a click
    document.getElementById("button").click();
  }
});