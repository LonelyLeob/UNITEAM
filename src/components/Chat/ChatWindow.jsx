import "./chat.css"

function ChatWindow() {
    return (
        <div className="chat_window">
            <div className="black__window">
                <div className="window__user">
                    <input className="inp__mes" type="text" placeholder="Enter message"/>
                    <button className="btn__send">Send</button> 
                </div>
            </div>
        </div>
    )
}

export default ChatWindow