import "./courseStyle.css"

function Sections(props) {

 return(
     <div>
         {props.section && props.section.map((item,idx) => {
                 return (
                     <div key={idx}>
                         <p className="courseContent">{item.content}</p>
                     </div>
                 )
             })}
     </div>
 )

}

export default Sections