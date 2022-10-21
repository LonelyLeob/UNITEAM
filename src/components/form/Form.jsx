import Field from "../field/Field";
import Header from "../../header/header";

function Form(props) {
    return (
        <div>
            <Header/>
                {props.json.map((item, idx) => {
                    return (
                    <div key={idx} className=" mx-40 inline-block p-5 border-double border-4 border-cyan-400 max-w-3xl bg-slate-100">
                        <h1 className="text-lg font-semibold pb-5">{item.Name}</h1>
                        <p className="pb-5">{item.Description}</p>
                        <Field fields={item.Fields}/>
                    </div>
                    )
                })}

        </div>
    );
}

export default Form;

