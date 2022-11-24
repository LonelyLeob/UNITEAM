import {Routes, Route} from "react-router-dom";
import AllForms from "../form/allForms";
import SignUp from "../signUp/signUp";
import SignIn from "../signIn/signIn";
import ChangeForm from "../changeForm/change";
import Header from "../header/header";

function App() {

  return (
      <div className="App">
          <Header/>
              <Routes>
                  <Route path="Form/Change" element={<ChangeForm/>}></Route>
                  <Route path="Form" element={<AllForms/>}></Route>
                  <Route path="signUp" element={<SignUp/>}></Route>
                  <Route path="signIn" element={<SignIn/>}></Route>
              </Routes>
      </div>
    );
}

export default App;
