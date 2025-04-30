import { useState } from "react";

import Button from "../components/Button";
import Header from "../components/Header";

const Core = () => {
    const [openedEditor, setOpenedEditor] = useState("html");

    const onTabClick = (editorName: "html" | "css" | "js") => {
        setOpenedEditor(editorName);
    };

    return (
        <div className="App">
        <Header></Header>
        <p>欢迎进入 Web Code Editor ！</p>
        <div className="tab-button-container">
            <Button
                title="HTML"
                onClick={() => {
                    onTabClick("html");
                }}
            />
            <Button
                title="CSS"
                onClick={() => {
                    onTabClick("css");
                }}
            />
            <Button
                title="JavaScript"
                onClick={() => {
                    onTabClick("js");
                }}
            />
        </div>

        <div className="editor-container">{openedEditor === "html" ? <p>HTML editor</p> : openedEditor === "css" ? <p>CSS editor</p> : <p>JavaScript editor</p>}</div>
    </div>
    )
}

export default Core;