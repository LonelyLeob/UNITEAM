import "./personalAreaStyle.css"
import Header from "../header/header";
import GetUser from "./requests/getUser";
import {useEffect, useState} from "react";
import ConvertUnix from "./requests/convertUnix";
import DeleteUser from "./requests/deleteUser";
import {useNavigate} from "react-router-dom";
import Modal from "../modal/modalWin";

function PersonalArea(){

    const[user, setUser] = useState([])
    const [isModal, setModal] = useState(false);
    const [password, setPassword] = useState('');
    const navigate = useNavigate()
    let content =
        <form action="" className="modalForm">
            <input placeholder="Пароль" value={password} onChange={event => setPassword(event.target.value)} className="modalFormName" type="password" autoComplete="on"/>
            <button type="submit" className="modalBtn" onClick={(e) =>  handleDelete(e)}>Удалить</button>
            <p>Для удаления аккаунта введите пароль</p>
        </form>

    useEffect(() => {
        GetUser(setUser)
    }, [])

    

    const handleDelete = async(e) => {
        e.preventDefault()
        await DeleteUser(user.name, password)
        localStorage.clear()
        navigate("/")
    }

    return (
        <div>
            <Header/>
            <div className="profileContainer">
                <div className="containerWrap">
                    <div className="profileWrap">
                        <div className="profileImg"></div>
                            <div className="profile">
                                <p> Имя пользователя: {user.name}</p>
                                <p>E-mail: {user.email}</p>
                                <p className="changePsw" onClick={() => {navigate("/restorePass")}}>Изменить пароль</p>
                                <p className="delProfile" onClick={() => setModal(true)}>Удалить аккаунт</p>
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

            <Modal
                isVisible={isModal}
                title="Удаление пользователя"
                content={content}
                footer={<p></p>}
                onClose={() => setModal(false)}/>

        </div>
    )
}

export default PersonalArea