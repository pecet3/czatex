console.log("hello client");

const entryForm = document.getElementById("entryForm")

let userName;

entryForm.addEventListener("submit",(e)=>{
    e.preventDefault();
    connectWs();
})

const messageForm = document.getElementById("messageForm")

messageForm.addEventListener("submit",(e)=>{
    e.preventDefault();
    const message = document.getElementById("message")

    if (userName === "" || message === ""){
        return
    }

    conn.send(userName + ": " + message.value);

    message.value = "";
})

function connectWs(){
    userName = document.getElementById("name").value
    const room = document.getElementById("room").value

    if (userName === "" || room === ""){
        return
    }

    if (window.WebSocket){
        conn = new WebSocket(`ws://localhost:3000/ws?room=${room}`)
        conn.onopen = (e)=>{
            alert("connected to ws server")
            entryForm.classList.add("hidden")
            messageForm.classList.remove("hidden")
            messageForm.classList.add("flex")
        }

        conn.onclose=(e)=>{
            alert("closed connection with ws server ")
        }

        conn.onmessage = (e)=>{
            const eventData = e.data

            console.log(eventData)
        }
    }else{
        alert("your browser doesn't support websockets")
    }
}