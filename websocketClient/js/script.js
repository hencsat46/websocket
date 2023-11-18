


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

    const response = await (await fetch(request)).json()

    if (response.Status == 200) {
        localStorage.setItem("UserId", response.Payload)
        window.location.href = "/chat";
    }

    


    console.log(response)
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

 
}