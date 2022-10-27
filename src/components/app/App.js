import Header from "../header/header";



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
  return (
      <div className="App">
        {/*<Form json={forms}/>*/}
          <Header/>
      </div>
    );
}

export default App;
