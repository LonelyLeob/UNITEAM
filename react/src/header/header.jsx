import "./headerStyle.css";
import { useNavigate  } from "react-router-dom";

function Header(){

    const navigate = useNavigate()

    let handlerSubmit = () => {
        localStorage.clear()
        navigate("/")
    }

    return (
        <div className="Header">
            <div className="upMenu">
                <h1 className="HeaderLogo">UNIVERSITY.Inc</h1>
                <ul>
                    <li onClick={() => navigate("/forms")}>
                       Мои формы
                    </li>
                    <li onClick={() => navigate("/course")}>
                        Курсы
                    </li>
                    <li onClick={() => navigate("/personalArea")}>
                        Личный кабинет
                    </li>
                  <li onClick={() => handlerSubmit()}>
                      Выход
                  </li>
                </ul>
            </div>
        </div>
    )
}


export default Header