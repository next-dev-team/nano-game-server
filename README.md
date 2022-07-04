# nano-game-server
A Game Server build on Nano Game Server Framework

**Example of Rolling Dice with ReactJs**
```
import React, { useEffect, useState } from "react";
import "./App.css";
import { Select } from "antd";
function App() {
  const [data, setData] = useState({});
  const [ws, setWebsocket] = useState<any>();
  const [isGameStarted, setIsGameStated] = useState(false);

  var join = function (data: any) {
    console.log("Join:", data);
    setData(data);
  };


  var onGuessGame = function (data: any) {
    console.log("on Guess game:", data);
    setData(data);
  };
  useEffect(() => {
    //@ts-ignore
    starx.init({ host: "127.0.0.1", port: 3251, path: "/game" }, function () {
      console.log("initialized");
      //@ts-ignore
      setWebsocket(starx);
      //@ts-ignore
      starx.request("dice.dogamestate", {status:1}, join);
    });
  }, []);

  return (
    <div className="App">
      <h3>Get started with web socket</h3>
      <h4>Ready to guess!!</h4>
      <pre>{JSON.stringify(data)}</pre>
      <button
        onClick={() => {
          setIsGameStated(!isGameStarted);
          ws.request("dice.dogamestate", {status:isGameStarted? 1:2}, join);
        }}
      >
        {!isGameStarted ? "Start Game" : "Stop Game"}
      </button>
      {isGameStarted && (
        <>
          <button
            style={{ marginLeft: "10px" }}
            onClick={() => {
              ws.request("dice.guessgame", {text:"big"}, onGuessGame);
            }}
          >
            Big
          </button>
          <button
            style={{ marginLeft: "10px" }}
            onClick={() => {
              ws.request("dice.guessgame", {text:"small"}, onGuessGame);
            }}
          >
            Small
          </button>
        </>
      )}
    </div>
  );
}

export default App;

```
**Reference:**

[Nano Core](https://github.com/lonng/nano)
<br/>
[Nano Server](https://github.com/lonng/nanoserver)
