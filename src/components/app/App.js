import Header from "../header/header";
import axios from "axios";


// const src = "http://localhost:8080/forms/get/forms"
// const src2 = "http://localhost:7000/token"
// const json = JSON.stringify({
//   "name": "abc",
//   "password": "abc"
// })
//
// const req = axios.post(src2, json, {withCredentials:true})


function App() {
  // const [forms, setForms] = useState([])
  //
  // useEffect(() => {
  //   axios
  //     .get(src, {withCredentials: true})
  //     .then(data => {
  //       setForms(data.data)
  //     })
  // })

    let json = JSON.stringify({"name": "name", "password": "password"})
    axios.post("http://localhost:7000/auth/token", json).then(data => console.log(data))

  return (
      <div className="App">
        {/*<Form json={forms}/>*/}
          <Header/>
      </div>
    );
}

export default App;
