import "./headerStyle.css";
import {BrowserRouter, Routes, Route, Link} from "react-router-dom";
import SignUp from "../signUp/signUp";
import SignIn from "../signIn/signIn";
import AllForms from "../form/allForms";


function Header(){

    return(
        <BrowserRouter>
        <div className="Header">
            <div className="upMenu">
                <h1 className="HeaderLogo">UNIVERSITY.Inc</h1>
                <ul>
                    <li>
                        <Link to="Form">Формы</Link>
                    </li>
                    <li>
                        <Link to="signIn">Вход</Link>
                    </li>
                    <li>
                        <Link to="signUp">Регистрация</Link>
                    </li>
                </ul>
            </div>
        </div>
            <Routes>
                <Route path="Form" element={<AllForms/>}></Route>
                <Route path="signUp" element={<SignUp/>}></Route>
                <Route path="signIn" element={<SignIn/>}></Route>
            </Routes>
        </BrowserRouter>
    )
}

export default Header