import Form from "./components/form/Form";
import stat from "./static.json"


function App() {
  return (
    <div className="App">
      <Form json={stat}/>
    </div>
  );
}

export default App;
