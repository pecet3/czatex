const entryForm = document.getElementById("entryForm")
const userName = document.getElementById("name")
const room = document.getElementById("room")

replaceInputValue()

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

    let data = {
        "name": userName,
        "message": message.value
    }
    
    conn.send(JSON.stringify(data));

    message.value = "";
    message.focus()
    
})

function connectWs(){
   
    if (userName.value === "" || room.value === ""){
        return
    }

    if (window.WebSocket){
        conn = new WebSocket(`ws://localhost:3000/ws?room=${room.value}`)
        conn.onopen = (e)=>{
            showDashboard()
            addQueryParams()
      
        }

        conn.onclose=(e)=>{
            alert("closed connection with ws server ")
        }

        conn.onmessage = (e)=>{
            const messagesList = document.getElementById("messagesList")
            const data = JSON.parse(e.data)
            
            const elementHTML = `<li class="p-1 bg-slate-400 rounded-md break-words w-96">${data.name} ___ ${data.message}</li>`
          
            messagesList.insertAdjacentHTML("beforeend",elementHTML)

        }
    }else{
        alert("your browser doesn't support websockets")
    }
}

function showDashboard(){
    const chatDashboard = document.getElementById("chatDashboard")
    entryForm.classList.add("hidden")
            
    chatDashboard.classList.remove("hidden")
    chatDashboard.classList.add("flex")
}

function addQueryParams(){
    const url = new URL(window.location.href)
    url.searchParams.set("room",room.value)
    history.replaceState(null,null, url.href)
}

function replaceInputValue(){
    const url = new URL(window.location.href)
    const params = new URLSearchParams(url.search)
    room.value = params.get("room")

}