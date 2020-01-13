import React from "react"

function Jokes(props) {
    return (
        <div>
            <h3 style={{display: !props.item.question && "none"}}> Question: {props.item.question}</h3>
            <p style={{color: !props.item.question && "#999999"}}> Answer: {props.item.punchLine}</p>
            <br />
            <hr />
        </div>
    )
}

export default Jokes