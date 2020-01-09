import React from "react"

function Footer() {
    const date = new Date()
    const hours = date.getHours()
    let timeOfDay
    const style = { backgroundColor: "#bbbbbb", fontSize: 30 }
    if (hours < 12 && hours > 3) {
        timeOfDay = "morning"
        style.color = "#EE9944"
    } else if (hours >= 12 && hours < 17) {
        timeOfDay = "afternoon"
        style.color = "#33ccaa"
    } else if (hours >= 17 && hours < 21) {
        timeOfDay = "evening"
        style.color = "#EE4499"
    } else {
        timeOfDay = "night"
        style.color = "#114455"
    }


    return (
        <footer>
            <br />
            <h2 style={style}>It is currently about {date.getHours() % 12} o'clock!</h2>
            <br />
            <h2 style={style}>Good {timeOfDay}!</h2>
        </footer>
    )
}
export default Footer