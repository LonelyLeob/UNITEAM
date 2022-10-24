import Form from "./../form/Form";
import axios from "axios";
import { useEffect, useState } from "react";

const src = "http://localhost:7000/forms/get/forms"
const src2 = "http://localhost:9000/token"
const json = JSON.stringify({
  "name": "usrr222",
  "password": "felani"
})

const req = axios.post(src2, json, {withCredentials:true})

function App() {
  const [forms, setForms] = useState([])

  useEffect(() => {
    axios
      .get(src, {withCredentials: true})
      .then(data => {
        setForms(data.data)
      })
  })
  return (
      <div className="App">
        <Form json={forms}/>
      </div>
    );
}

export default App;
