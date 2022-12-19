import SignIn from "../entrance/signIn";
import {Routes, Route} from "react-router-dom";
import SignUp from "../entrance/signUp";
import FormProcessing from "../dynamicForm/forms/requests/formProcessing";
import FormEdit from "../dynamicForm/formsEdit/formEdit";
import PersonalArea from "../personalArea/personalArea"
import RestorePass from "../entrance/restorePass";
import Course from "../course/course";
import ViewCourse from "../course/viewCourse";

function App() {

  return (
      <div className="App">
          <Routes>
              <Route path="signUp" element={<SignUp/>}></Route>
              <Route path="/" element={<SignIn/>}></Route>
              <Route path="forms" element={<FormProcessing/>}></Route>
              <Route path="/edit/:uuid" element={<FormEdit/>}></Route>
              <Route path="/personalArea" element={<PersonalArea/>}></Route>
              <Route path="/restorePass" element={<RestorePass/>}></Route>
              <Route path="/course" element={<Course/>}></Route>
              <Route path="/viewCourse/:id" element={<ViewCourse/>}></Route>
          </Routes>
      </div>
  )
}

export default App;
