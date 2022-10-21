import Form from "./components/form/Form";
import stat from "./static.json"
import axios from "axios";

function App() {
  let json = JSON.stringify({"name": "usrr2", "password": "felani"})
  let json2 = JSON.stringify({
    "name":"usrr6",
    "desc": "felani55",
    "anon": true
})

  let req = axios.post("http://localhost:7000/forms/create", json2, {withCredentials: true}).then(res => console.log(res))

  return (
    <div className="App">
      <Form json={stat}/>
    </div>
  );
}

export default App;
