import React from "react"

import Contacts from "./Contacts"

function App() {

    const contactList = {
        "display": "flex",
        "flexWrap": "wrap"
    }

    return (
        <div style={contactList}>
            <Contacts
                contact={{
                    name: "Mr. Whiskerson",
                    imgUrl: "http://placekitten.com/300/200",
                    phone: "(212) 555-1234",
                    email: "mr.whiskaz@catnap.meow"
                }}
            />

            <Contacts
                contact={{
                    name: "Fluffykins",
                    imgUrl: "http://placekitten.com/400/200",
                    phone: "(212) 555-2345",
                    email: "fluff@me.com"
                }}
            />

            <Contacts
                contact={{
                    name: "Destroyer",
                    imgUrl: "http://placekitten.com/400/300",
                    phone: "(212) 555-3456",
                    email: "ofworlds@yahoo.com"
                }}
            />

            <Contacts
                contact={{
                    name: "Felix",
                    imgUrl: "http://placekitten.com/200/100",
                    phone: "(212) 555-4567",
                    email: "thecat@hotmail.com"
                }}
            />
        </div>



    )
}

export default App