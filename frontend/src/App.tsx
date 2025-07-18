import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";

import axios from "axios";

function App() {
  const [count, setCount] = useState(0);

  axios
    .get("/api/1.0/")
    .then(function (response) {
      console.log(response?.data);
    })
    .catch(function (error) {
      console.log(error);
    });

  return (
    <>
      <div className="flex flex-row justify-center">
        <a href="https://vite.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p className="mt-8">
          Edit{" "}
          <code className="font-bold text-red-500 italic">src/App.tsx</code> and
          save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  );
}

export default App;
