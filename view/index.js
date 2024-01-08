const entryForm = document.getElementById("entryForm")

const userName = document.getElementById("name")
const room = document.getElementById("room")

const generateBtn = document.getElementById("generateBtn")

replaceInputRoom("room_1")

// 

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
    const date = getCurrentDateTimeString()
    console.log
    let data = {
        "name": userName.value,
        "message": message.value,
        "date": date
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
          
            const data = JSON.parse(e.data)
   
            const elementHTML = `
            <li class="p-1 bg-slate-400 rounded-md break-words max-w-xl">
                <a class="font-bold">[${data.name}] </a> 
                <a class="italic">${data.message}</a>
                
                ${typeof data.date !== 'undefined' ? `<a class="mono text-xs text-gray-700">${data.date}</a>` : ""}
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

    const roomDisplay = document.getElementById("roomDisplay")
    roomDisplay.textContent = room.value
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

// helpers

function addQuery(param,value){
    const url = new URL(window.location.href)
    url.searchParams.set(param,value)
    history.replaceState(null,null, url.href)
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


function getCurrentDateTimeString() {
    const currentDate = new Date();
  
    const year = currentDate.getFullYear();
    const month = (currentDate.getMonth() + 1).toString().padStart(2, '0'); // Miesiące są od 0 do 11, więc dodajemy 1
    const day = currentDate.getDate().toString().padStart(2, '0');
    const hours = currentDate.getHours().toString().padStart(2, '0');
    const minutes = currentDate.getMinutes().toString().padStart(2, '0');
  
  
    const dateTimeString = `${year}-${month}-${day} ${hours}:${minutes}`;
  
    return dateTimeString;
  }
  