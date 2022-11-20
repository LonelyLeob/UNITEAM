import "./changeStyle.css"
import ChangeFields from "./changeFields";
import axios from "axios";
import {useState} from "react";

function ChangeForm() {

    const[field, setField] = useState('')

    const data = JSON.parse(localStorage.getItem("data"));

    // if (data != null) {
    //     localStorage.removeItem("data")}

    let uuid = data.Uuid

    let handleSubmit = async () => {
        let res = await axios.post(`http://uni-team-inc.online:8080/api/v1/create/field?form=${uuid}`,
            JSON.stringify(
                {
                    quiz:field
                }),{headers:{
                    Authorization:`Bearer ${localStorage.getItem('access')}`
                }}
        )}


    if (data.Fields !== null) {
        return (
            <div className="changeContainer">
                    <div className="changeContainerContent">
                        <h1>{data.Name}</h1>
                        <p> {data.Description}</p><br/>
                        <div className="wrapperBtn">
                            <input type="text" placeholder="Добавить поле" maxLength="50" value={field} onChange={event => setField(event.target.value)}/>
                            <button className="btn" onClick={(e) => {handleSubmit(e)}}>+</button>
                        </div>
                    </div><br/>
                <ChangeFields fields={data.Fields}/>
            </div>
        )}

    return (
        <div className="changeContainer">
            <div className="changeContainerContent">
                <h1>{data.Name}</h1>
                <p> {data.Description}</p><br/>
                <div className="wrapperBtn">
                    <input type="text" placeholder="Добавить поле" maxLength="50" value={field} onChange={event => setField(event.target.value)}/>
                    <button className="btn" onClick={(e) => {handleSubmit(e)}}>+</button>
                </div>
            </div><br/>
        </div>
    )}

export default ChangeForm