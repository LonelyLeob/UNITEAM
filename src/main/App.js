import SignIn from "../entrance/signIn";
import {Routes, Route} from "react-router-dom";
import SignUp from "../entrance/signUp";
import FormProcessing from "../dynamicForm/forms/requests/formProcessing";
import FormEdit from "../dynamicForm/formsEdit/formEdit";
import PersonalArea from "../personalArea/personalArea"
import RestorePass from "../entrance/restorePass";

function App() {

  return (
      <div className="App">
          <Routes>
              <Route path="signUp" element={<SignUp/>}></Route>
              <Route path="/" element={<SignIn/>}></Route>
              <Route path="forms" element={<FormProcessing/>}></Route>
              <Route path="/edit" element={<FormEdit/>}></Route>
              <Route path="/personalArea" element={<PersonalArea/>}></Route>
              <Route path="/restorePass" element={<RestorePass/>}></Route>
          </Routes>
      </div>
  )
}

export default App;
