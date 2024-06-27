import { useEffect, useRef, useState } from "react";

interface MessageWS {
  clientAlias: string;
  message: string;
  timestamps: string;
}
const URL = "ws://localhost:8080/ws";

export default function ChatComponent() {
  const socket = useRef<WebSocket | null>(null);
  const [messages, setMessage] = useState<MessageWS[]>([]);
  const [input, setInput] = useState<string>("");

  useEffect(() => {
    if (socket.current === null) {
      socket.current = new WebSocket(URL);
    }

    socket.current.addEventListener(
      "message",
      (event: MessageEvent<string>) => {
        const message = JSON.parse(event.data) as MessageWS;

        setMessage([...messages, message]);
      }
    );
  });

  return (
    <>
      <div>
        {messages.map((value) => (
          <p>
            {value.clientAlias}: {value.message}
          </p>
        ))}
      </div>
      <div className="input-section">
        <input
          className="msg-input"
          onChange={(event) => {
            setInput(event.target.value);
          }}
          value={input}
        />
        <button
          onClick={() => {
            if (input.length < 2) {
              alert("Message should have at least 2 characters");
              return;
            }
            socket.current?.send(input);
            setInput("");
          }}
        >
          enviar
        </button>
      </div>
    </>
  );
}
