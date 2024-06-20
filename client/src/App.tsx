import { ChangeEvent, useState } from 'react'
import './App.css'

const ws = new WebSocket("ws://localhost:8080/ws")

type MessageWs = {
  clientAlias: string
  message: string
}
function App() {
  const [inputValue, setInputValue] = useState('')
  const [listMessages, setListMessages] = useState<MessageWs[]>([])
  

  ws.onmessage = (event: MessageEvent<string>) => {
    let message = JSON.parse(event.data) as MessageWs 
    setListMessages([...listMessages, message])
  }

  const handleInputChange = (event: ChangeEvent<HTMLInputElement>) => {
    setInputValue(event.target.value);
  };
  
  function sendMessage(){
    ws.send(inputValue)
    setInputValue("")
  }
  return (
    <>
      <input className="msg-input" onChange={handleInputChange} value={inputValue}/>
      <button onClick={sendMessage}>enviar</button>

      <div>
          {listMessages.map((item) => (
            <p>{item.clientAlias}: {item.message}</p>
          ))}
      </div>      
    </>
  )
}

export default App
