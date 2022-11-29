import "./modalStyle.css"
import {useState, useEffect} from "react";
import ModalAdd from "./modalAddForm";

function ModalWin(){

    const [isModal, setModal] = useState(false);


    const Modal = ({ isVisible = false, title, onClose }) => {
        const keydownHandler = ({ key }) => {
            switch (key) {
                case 'Escape':
                    onClose();
                    break;
                default:
            }
        }


        useEffect(() => {
            document.addEventListener('keydown', keydownHandler);
            return () => document.removeEventListener('keydown', keydownHandler);
        })

        return !isVisible ? null : (
            <div className="modal" onClick={onClose}>
                <div className="modal-dialog" onClick={e => e.stopPropagation()}>
                    <div className="modal-header">
                        <h3 className="modal-title">{title}</h3>
                        <span className="modal-close" onClick={onClose}>
            &times;
          </span>
                    </div>
                    <div className="modal-body">
                        <div className="modal-content">
                            <ModalAdd setModal = {setModal}/>
                        </div>
                    </div>
                    <div className="modal-footer"></div>
                </div>
            </div>
        )
    }

    return (
        <>
            <button onClick={() => setModal(true)}>+</button>
            <Modal
                isVisible={isModal}
                title="Создание формы"
                onClose={() => setModal(false)}
            />
        </>
    )
}


export default ModalWin
