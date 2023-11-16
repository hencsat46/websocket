
const url = "ws://localhost:3000/ws"

const ws = new WebSocket("ws://localhost:3000/ws")
console.log(ws.readyState)
ws.onopen = function() {
    console.log("Connected")
}

ws.onmessage = (event) => {
    console.log(event)
}

function sendMessage() {
    const message = document.querySelector("input").value
    console.log(message)
    ws.send(message)
}