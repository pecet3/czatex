console.log("hello");

const entryForm = document.getElementById("entryForm")


entryForm.addEventListener("submit",(e)=>{
    e.preventDefault();
    connectWs();
})

const messageForm = document.getElementById("messageForm")

messageForm.addEventListener("submit",(e)=>{
    e.preventDefault();
    const inputMessage = document.getElementById("message").value
    conn.send(inputMessage);
})

function connectWs(){
    const name = document.getElementById("name").value
    const room = document.getElementById("room").value

    if (name === "" || room === ""){
        return
    }

    if (window.WebSocket){
        conn = new WebSocket(`ws://localhost:3000/ws?room=${room}&name=${name}`)
        conn.onopen = (e)=>{
            alert("connected to ws server")
            conn.send("hello");

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