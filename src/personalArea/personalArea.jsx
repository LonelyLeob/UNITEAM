import "./personalAreaStyle.css"
import Header from "../header/header";
import GetUser from "./requests/getUser";
import {useEffect, useState} from "react";
import ConvertUnix from "./requests/convertUnix";
import DeleteUser from "./requests/deleteUser";
import {useNavigate} from "react-router-dom";

function PersonalArea(){

    const[user, setUser] = useState([])
    const navigate = useNavigate()

    useEffect(() => {
        GetUser(setUser)
    }, [])

    

    const handleDelete = async(e) => {
        e.preventDefault()
        // await DeleteUser(user, pass)
        localStorage.clear()
        navigate("/")
    }

    return (
        <div>
            <Header/>
                <div className="btnDelContainer">
                    <label> Удалить пользователя
                    <button className="btnDel" onClick={(e) => {handleDelete(e)}}>-</button>
                    </label>
                </div>
            <div className="profileContainer">
                <div className="containerWrap">
                    <div className="profileWrap">
                        <div className="profileImg"></div>
                            <div className="profile">
                                <p> Имя пользователя: {user.name}</p>
                                <p>E-mail: {user.email}</p>
                                <p className="changePsw">Изменить пароль</p>
                            </div>
                    </div>

                        <div className="sessionsWrap">
                                {user.meta ? user.meta.map((item,idx)=>{
                                    return(
                                        <div key={idx} className="sessions">
                                            <button className="answerDelBtn">x</button>
                                            <p>Сессия открыта в: {ConvertUnix(item.lv)}</p>
                                            <p>Браузер: {item.bwr}</p>
                                            <p>Операционная система: {item.os}</p>
                                        </div>
                                    )
                                }): <p></p> }
                    </div>
                </div>
            </div>

        </div>
    )
}

export default PersonalArea