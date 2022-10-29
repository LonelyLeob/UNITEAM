import Header from "../header/header";
import axios from "axios";

function App() {
    let json = JSON.stringify({"name": "name", "password": "password"})
    axios.post("http://localhost:7000/auth/token", json, {withCredentials: true}).then(data => console.log(data))

  return (
      <div className="App">
          <Header/>
      </div>
    );
}

export default App;
