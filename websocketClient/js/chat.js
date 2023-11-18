const url = "ws://localhost:3000/ws"

const ws = new WebSocket("ws://localhost:3000/ws")
console.log(ws.readyState)
ws.onopen = function() {
    console.log("Connected")
}

ws.onmessage = (event) => {
    const responseMessage = JSON.parse(event.data)
    
    console.log(responseMessage)
    
    const message = document.createElement("div")
    message.classList.add("message")
    message.innerText = responseMessage.message
    document.querySelector(".messages").append(message)
}

function sendMessage() {
    const message = document.querySelector("input").value

    ws.send(message)
}



