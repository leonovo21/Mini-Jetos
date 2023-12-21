import { useState } from 'react';
import './App.css'
function FetchCall(){
    const [title, setTitle] = useState("")
    const [email, setEmail] = useState("")
}
function App() {
    return(
        <body className='nb'>
            <style>{'body { background-color: wheat; }'}</style>
            <h1 className="header">Hatsu</h1>

            <input className="title" type='title' placeholder='Title'></input>

            <input className="email" type='email' placeholder='Email'></input>

            <textarea className="Write-Area" placeholder='Write here'></textarea>

            <button className='btn-send' onClick={send}>Send</button>
        </body>
    )
}
function send() {
    console.log("Send");
}
export default App