import Field from "../field/Field";


function Form(props) {
    if (props.item.Fields == null) {
        return (
            <div key={props.key} className=" mx-40 inline-block p-5 border-double border-4 border-cyan-400 max-w-3xl bg-slate-100">
                    <h1 className="text-lg font-semibold pb-5">{props.item.Name}</h1>
                    <p className="pb-5">{props.item.Description}</p>
            </div>
        );
    }

    return (
        <div className=" mx-40 inline-block p-5 border-double border-4 border-cyan-400 max-w-3xl bg-slate-100">
                <h1 className="text-lg font-semibold pb-5">{props.item.Name}</h1>
                <p className="pb-5">{props.item.Description}</p>
            <Field fields={props.item.Fields}/>
        </div>
    );
}

export default Form;

