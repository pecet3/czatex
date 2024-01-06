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
    const message = e.target.elements.message

    if (userName === "" || message.value === ""){
        return
    }
    
    conn.send(userName + ": " + message.value);

    message.value = "";
    message.focus()
    
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
            const chatDashboard = document.getElementById("chatDashboard")

            alert("connected to ws server")
            entryForm.classList.add("hidden")
            
            chatDashboard.classList.remove("hidden")
            chatDashboard.classList.add("flex")
        }

        conn.onclose=(e)=>{
            alert("closed connection with ws server ")
        }

        conn.onmessage = (e)=>{
            const messagesList = document.getElementById("messagesList")
            const eventData = e.data
            const elementHTML = `<li class="p-1 m-1 bg-slate-400 rounded-md">${eventData}</li>`
          
            messagesList.insertAdjacentHTML("beforeend",elementHTML)

        }
    }else{
        alert("your browser doesn't support websockets")
    }
}

