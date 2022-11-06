import "./headerStyle.css";
import {BrowserRouter, Routes, Route, Link} from "react-router-dom";
import SignUp from "../signUp/signUp";
import SignIn from "../signIn/signIn";
import Form from "../form/Form";
import axios from "axios";
import {useEffect, useState} from "react";





function Header(){

    const [data, setData] = useState([])

    useEffect(() => {
        axios.get("http://localhost:8080/forms/get/forms")
            .then(res => {
                setData(res.data)
            }).catch(err =>
        console.log(err))
    })

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
                <Route path="Form" element={<Form getJson={data} />}></Route>
                <Route path="signUp" element={<SignUp/>}></Route>
                <Route path="signIn" element={<SignIn/>}></Route>
            </Routes>
        </BrowserRouter>
    )
}

export default Header