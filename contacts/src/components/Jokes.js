import React from "react"

function Jokes(props) {
    return (
        <div>
            <h3 style={{display: !props.question && "none"}}> Question:{props.question}</h3>
            <p style={{color: !props.question && "#999999"}}> Answer: {props.punchLine}</p>
            <br />
            <hr />
        </div>
    )
}

export default Jokes