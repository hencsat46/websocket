const url = "ws://localhost:3000/ws"

const ws = new WebSocket("ws://localhost:3000/ws")
console.log(ws.readyState)
ws.onopen = function() {
    console.log("Connected")
    getMessages()


}

async function getMessages() {
    const request = new Request("http://localhost:3000/getmessages", {
        method: "GET",
        mode: "cors",
    })

    const response = await (await fetch(request)).json()
    console.log(response)
    
    if (response.Payload.length != 0 && response.Payload[0].Message_text != "Internal Server Error") {
        for (let i = 0; i < response.Payload.length; ++i) {
            drawMessage(response.Payload[i].Message_text)
        }
    }
}

function drawMessage(text) {
    const message = document.createElement("div")
    message.classList.add("message")
    message.innerText = text
    document.querySelector(".messages").append(message)
}

ws.onmessage = (event) => {
    const responseMessage = JSON.parse(event.data)
    
    console.log(responseMessage)

    drawMessage(responseMessage.message)
    
    
}

function sendMessage() {
    const message = document.querySelector("input").value

    const data = {
        userId: localStorage.getItem("UserId"),
        Message: message,
    }

    console.log(data)

    ws.send(JSON.stringify(data))
}



