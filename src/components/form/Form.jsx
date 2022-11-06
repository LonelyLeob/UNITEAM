import Field from "../field/Field";

function Form(props) {

    return (
        <div>

                {props.getJson.map((item, idx) => {
                    if (item.Fields <= 1){
                        return(
                            <div key={idx} className="mx-40 inline-block p-5 border-double border-4 border-cyan-400 max-w-3xl bg-slate-100">
                                <h1 className="text-lg font-semibold pb-5">{item.Name}</h1>
                                <p className="pb-5">{item.Description}</p>
                            </div>
                    )
                    }
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

