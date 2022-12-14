import "./courseStyle.css"
import Header from "../header/header";
import GetShort from "./requests/getShortCourses";
import {useEffect, useState} from "react";
import {Link} from "react-router-dom";
import AddCourse from "./requests/addCourse";
import Modal from "../modal/modalWin";

function Course() {

    const [course, setCourse] = useState([])
    const [title, setTitle] = useState('')
    const [desc, setDesc] = useState('')
    const [isModal, setModal] = useState(false);
    const [range, setRange] = useState('6')
    const[count,setCount] = useState(0)
    let content =
        <form action="" className="modalForm">
            <input placeholder="Название курса" value={title} onChange={event => setTitle(event.target.value)} className="modalFormName" type="text" />
            <input placeholder="Описание" value={desc} onChange={event => setDesc(event.target.value)} className="modalFormDesc" type="text" maxLength="150"/>
            <button type="submit" className="modalBtn" onClick={(e) =>  addHandler(e)}>Добавить</button>
        </form>



    useEffect(() => {
        request()
    }, [count])

   const request = async () => {
      await GetShort(setCourse)
   }

    let addHandler = async(e) => {
        e.preventDefault()
        await AddCourse(title, desc)
        setCount(+ 1)
        setModal(false)
    }

    return (
        <div className="course">
            <Header/>
                <div className="titleAdd">
                    <p onClick={() => setModal(true)}>Добавить курс</p>
                </div>
                    <div className="mainCourseWrap">
                                <div className="menuContainer">
                                    <h2>Тип обучения:</h2>
                                    <div className="radioForm">
                                        <label htmlFor="another">Любой <input type="radio" id="another"/></label>
                                        <label htmlFor="profession">Профессия  <input type="radio" id="profession"/></label>
                                        <label htmlFor="courses">Курсы <input type="radio" id="courses"/></label>
                                    </div>
                                    <p>Длительность: {range} месяцев</p>
                                    <input type="range" min="0" max="12" step="1" value={range} onChange={(e) => setRange(e.target.value)}/>
                                </div>
                        <div className="courseContainer">
                            {course.map((item, idx) => {
                                return(
                                    <div className="courseBlock" key={idx}>
                                        <Link to={`/viewCourse/${item.id}`} className="courseTitleLink">
                                            <p className="courseTitle">{item.title}</p>
                                        </Link>
                                        <div className="courseWrap">
                                            <div className="courseDesc">
                                                <p>Автор: {item.author}</p>
                                                <p className="coursePrice">Цена: 200$</p>
                                            </div>
                                            <div className="imageContainer">
                                                <img src="../testImage.png" alt="Изображение курса" className="image"/>
                                            </div>
                                        </div>
                                    </div>
                                )
                            })}
                        </div>
                    </div>

            <Modal
                isVisible={isModal}
                title="Добавление секции" content={content}
                footer={<p></p>} onClose={() => setModal(false)}
            />
        </div>
    )
}

export default Course