import React from "react"
import Contacts from "./components/Contacts"
import ContactsData from "./components/ContactsData"
import Jokes from "./components/Jokes"
import JokesData from "./components/JokesData"
import Products from "./components/Products"
import ProductsData from "./components/ProductsData"

function App() {

    const contactList = {
        "display": "flex",
        "flexWrap": "wrap"
    }

    const jokeComponents = JokesData.map(joke =>
        <Jokes
            key={joke.id}
            item={joke}
        />
    )

    const contactComponents = ContactsData.map(contact =>
        <Contacts
            key={contact.id}
            item={contact}
        />
    )

    const productsCompomemts = ProductsData.map(product =>
        <Products
            key={product.id}
            item={product}
        />
    )

    return (
        <div>
            <div style={contactList}>
                {contactComponents}
            </div>
            <div>
                <h2 >Jokes: </h2>
                <br />
                {jokeComponents}
            </div>
            <div>
                <h2 >Products:</h2>
                <br />
                {productsCompomemts}
            </div>
        </div>



    )
}

export default App