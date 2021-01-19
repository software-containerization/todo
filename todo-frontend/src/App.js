import React, { useState, useEffect } from 'react';
import logo from './logo.svg';
import './App.css';
import Api from "./Api.js"

const apiUrl = "http://localhost:8080/api"

function App() {
    const [items, setItems] = useState([]);
    const [item, setItem] = useState("");

    const fetchAndSetItems = async () => {
        const items = await Api.getAllItems();
        console.log(items);
        setItems(items);
    }
    
    useEffect(() => {
        fetchAndSetItems();
    }, []);
    
    // async function fetchItems() {
    //     const response = await fetch(apiUrl + '/items');
    //     const json = await response.json();
    //     setItems(json);
    // }

    async function createItem(e, i) {
        const newItems = [...items];
        
    }

    async function handleKeyDown(e){
        if (e.key === 'Enter' && item.title !== "") {
            //create new item
            const newItem = {
                done: false,
                title: item
            }
            var a = await Api.createItem(newItem);
            console.log(a);
            await fetchAndSetItems();
            setItem("");
        }
    }

    async function deleteItem(e, id) {
        e.stopPropagation();
        await Api.deleteItem(id);
        await fetchAndSetItems();
    }

    async function toggleCompleteItem(e, item) {
        item.done = !item.done;
        await Api.updateItem(item.id, item);
        await fetchAndSetItems();
    }

    return (
        <div className="App">
          <header className="App-header">
            <img src={logo} className="App-logo" alt="logo" />
          </header>

          <input
            className="new-item-input"
            type="text"
            value={item}
            placeholder="Add a new item"
            onChange={({ target }) => setItem(target.value)}
            onKeyDown={e => handleKeyDown(e)}
            />
          <div className="item-list">
              {items.map((item, i) => (
                  <div className={`item ${item.done && 'item-is-done'}`}>
                    <div className={'checkbox'} onClick={e => toggleCompleteItem(e, item)}>
                      { item.done && (
                          <span>&#x2714;</span>
                      )}
                  </div>
                      <span className="item-title">{item.title}</span>
                      <span className="delete-item"onClick={e => deleteItem(e, item.id)}>X</span>
                      </div>
              ))}
            </div>
            </div>
    );
}

export default App;
