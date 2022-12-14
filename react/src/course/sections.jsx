import "./courseStyle.css"

function Sections(props) {

 return(
     <div className="viewCourseContainer">
         {props.section && props.section.map((item,idx) => {
                 return (
                     <div key={idx} className="courseContent">
                         <p>{item.content}</p>
                         <button className="readNext">Читать далее...</button>
                     </div>
                 )
             })}

             
     </div>
 )

}

export default Sections