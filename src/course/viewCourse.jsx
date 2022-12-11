import "./courseStyle.css"
import {useNavigate, useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import GetCourse from "./requests/getCourses";
import Modal from "../modal/modalWin";
import AddSection from "./requests/addSection";
import DelCourse from "./requests/delCourse";
import Sections from "./sections";

function ViewCourse() {

    const params = useParams()
    const [course, setCourse] = useState([])
    const navigate = useNavigate()
    const [isModal, setModal] = useState(false);
    const[section,setSection] = useState('')
    const[count,setCount] = useState(0)
    const courseId = params.id
    let content =
        <form action="" className="modalForm" id="modalForm">
            <textarea
                name="section" cols="50" rows="10" required form="modalForm" className="modalFormSection" value={section}
                placeholder="Контентная секция" onChange={event => setSection(event.target.value)}>
            </textarea>
            <button type="submit" className="modalBtn" onClick={(e) =>  formHandler(e)}>Добавить</button>
        </form>

    useEffect(() => {
        request()
    }, [count])

    const request = async() => {
        await GetCourse(setCourse, courseId)
    }

    let formHandler = async(e) => {
        e.preventDefault()
        if (section){
            await AddSection(parseInt(courseId), section)
            setCount(+ 1)
            setModal(false)
        }
    }

    let delHandler = async (e) => {
        e.preventDefault()
        await DelCourse(courseId)
        navigate(-1)
    }

    return (
        <div>
            <p onClick={() => navigate(-1)}>Вернуться назад</p>
            <p onClick={(e) => delHandler(e)}>Удалить курс</p>
            <p>{course.title}</p>
            <p>{course.desc}</p>
            <p>{course.author}</p>
            <button onClick={() => setModal(true)}>Добавить секцию</button>
            <Sections section = {course.sections}/>
                <Modal
                    isVisible={isModal}
                    title="Добавление секции" content={content}
                    footer={<p></p>} onClose={() => setModal(false)}
                />
        </div>
    )
}

export default ViewCourse