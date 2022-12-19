import "./courseStyle.css"
import {useState} from "react";
import UpdateSection from "./requests/updateSection";
import Modal from "../modal/modalWin";


function Sections(props) {

    const[section,setSection] = useState('')
    const[id,setId] = useState()
    const [isModal, setModal] = useState(false);
    let content =
        <form className="modalForm" id="modalForm">
            <textarea
                name="section" cols="50" rows="10" required form="modalForm" className="modalFormSection" value={section}
                placeholder="Контентная секция" onChange={event => setSection(event.target.value)}>
            </textarea>
            <button type="submit" className="modalBtn" onClick={(e) =>  formHandler(e)}>Изменить</button>
        </form>

    let formHandler = async (e) => {
        e.preventDefault()
        await UpdateSection(id, section)
        setModal(false)
    }

 return(
     <div className="viewCourseContainer">
         {props.section?.map((item,idx) => {
             console.log(props.section)
             return (
                 <div key={idx} className="courseContent">
                     <p>{item.content}</p>
                     <div className="sectionBtn">
                         <button className="readNext">Читать далее...</button>
                         <button className="updateSection" onClick={() => {
                             setModal(true)
                             setSection(item.content)
                             setId(item.id)
                         }}
                         >Изменить</button>
                     </div>

                 </div>
             )
         })}

         <Modal
             isVisible={isModal}
             title="Изменение секции" content={content}
             footer={<p></p>} onClose={() => setModal(false)}
         />

     </div>
 )

}

export default Sections