import Field from "../field/Field";
import "./Form.css"

function Form(props) {

        if (props.item.Fields == null) {
            return (

                    <div key={props.key} className="formContainer">
                        <h1 className="formH">{props.item.Name}</h1>
                        <p className="formP">{props.item.Description}</p>
                    </div>

            );
        }

        return (

                <div className="formContainer">
                    <h1 className="formH">{props.item.Name}</h1>
                    <p className="formP">{props.item.Description} {`eto form ${props.Uuid}`}</p>
                    <Field fields={props.item.Fields}/>
                </div>

        );

}

export default Form;

