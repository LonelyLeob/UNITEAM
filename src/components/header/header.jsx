import "./headerStyle.css";

function Header(){
    return(
        <div className="Header">
            <div className="upMenu">
                <h1 className="HeaderLogo">UNIVERSITY.Inc</h1>
                <ul>
                    <li>Портфолио</li>
                    <li>Курсы</li>
                    <li>Вход</li>
                    <li>Регистрация</li>
                </ul>
            </div>
        </div>
    )
}

export default Header