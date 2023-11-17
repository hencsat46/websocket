
// const url = "ws://localhost:3000/ws"

// const ws = new WebSocket("ws://localhost:3000/ws")
// console.log(ws.readyState)
// ws.onopen = function() {
//     console.log("Connected")
// }

// ws.onmessage = (event) => {
//     console.log(event.data)
// }

// function sendMessage() {
//     const message = document.querySelector("input").value
//     console.log(message)
//     ws.send(message)
// }

async function signIn() {
    const user = document.querySelector(".uname").value
    const pass = document.querySelector(".passw").value

    const json = {Username: user, Password: pass}

    console.log(JSON.stringify(json))

    const request = new Request("http://localhost:3000/signin", {
        method: "POST",
        mode: "cors",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(json)
    })

    const response = await fetch(request)

    console.log(await response.json())
}

async function signUp() {
    const username = document.querySelector(".uname").value
    const password = document.querySelector(".passw").value

    const data = {
        Username: username,
        Password: password,
    }

    

    const request = new Request("http://localhost:3000/signup", {
        method: "POST",
        mode: "no-cors",
        body: JSON.stringify(data),
    })

    const response = await (await fetch(request)).json()

    console.log(response)
}