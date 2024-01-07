const entryForm = document.getElementById("entryForm")

const userName = document.getElementById("name")
const room = document.getElementById("room")

const generateBtn = document.getElementById("generateBtn")

replaceInputRoom("room_1")


generateBtn.addEventListener("click",()=>{
    room.value = generateRoomName(8)

})

entryForm.addEventListener("submit",(e)=>{
    e.preventDefault();

    connectWs();
})

const messageForm = document.getElementById("messageForm")

messageForm.addEventListener("submit",(e)=>{
    e.preventDefault();
    const message = e.target.elements.message

    if (userName.value === "" || message.value === ""){
        return
    }

    let data = {
        "name": userName.value,
        "message": message.value
    }
    
    conn.send(JSON.stringify(data));

    message.value = "";
    message.focus()
    
})

///// functions

function connectWs(){
   
    if (userName.value === "" || room.value === ""){
        return
    }

    if (window.WebSocket){
        conn = new WebSocket(`ws://localhost:3000/ws?room=${room.value}&name=${userName.value}`)
        conn.onopen = (e)=>{
            showDashboard()
            addQuery("room",room.value)
      
        }

        conn.onclose=(e)=>{
            alert("closed connection with ws server ")
        }

        conn.onmessage = (e)=>{
            const messagesList = document.getElementById("messagesList")
            console.log(e.data)
            const data = JSON.parse(e.data)
            
            const elementHTML = `
            <li class="p-1 bg-slate-400 rounded-md break-words max-w-xl">
                <a class="font-bold">[${data.name}] </a> 
                <a class="italic">${data.message}</a<>
            </li>`
          
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

function addQuery(param,value){
    const url = new URL(window.location.href)
    url.searchParams.set(param,value)
    history.replaceState(null,null, url.href)
}

function replaceInputRoom(value){
    const url = new URL(window.location.href)
    const params = new URLSearchParams(url.search)
    const queryRoom = params.get("room")
    if (queryRoom ==="" || queryRoom === null){
        room.value = value
        return
    }
    room.value = queryRoom
}


function generateRoomName(length) {
    const characters = '0123456789ABCDE';
    let randomId = '0x';
  
    for (let i = 0; i < length; i++) {
      const randomIndex = Math.floor(Math.random() * characters.length);
      randomId += characters.charAt(randomIndex);
    }
  
    return randomId;
}