import "./edit.css"

function FormEdit(){

    let data = JSON.parse(localStorage.getItem('data'))

    return(
                <div>
                    {console.log(data)}
                </div>
    )
}
export default FormEdit

