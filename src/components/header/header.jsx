import "./headerStyle.css";
import {Link} from "react-router-dom";
import { useNavigate  } from "react-router-dom";



function Header(){

    let navigate = useNavigate()
    const outHandler = () => {
        localStorage.clear()
        navigate("/")
    }

        if (localStorage.getItem("access") !== null) {
            return (
                <div className="Header">
                    <div className="upMenu">
                        <h1 className="HeaderLogo">UNIVERSITY.Inc</h1>
                        <ul>
                            <li>
                                <Link to="Form">Формы</Link>
                            </li>
                            <li>
                                <button onClick={() => {outHandler()}}>Выход</button> 
                            </li>
                            
                        </ul>
                    </div>
                </div>
            )}

    return(
        <div className="Header">
            <div className="upMenu">
                <h1 className="HeaderLogo">UNIVERSITY.Inc</h1>
                <ul>
                    <li>
                        <Link to="signIn">Вход</Link>
                    </li>
                    <li>
                        <Link to="signUp">Регистрация</Link>
                    </li>
                </ul>
            </div>
        </div>

    )
}

export default Header