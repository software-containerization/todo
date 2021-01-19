import React, { useState, useEffect } from 'react';

function Item(props) {
    const [done, setDone] = useState(props.done);
    const [title, setTitle] = useState(props.title);
    
    return (
        <div className={`item $(item.done && 'item-is-done')`}>
          <div className={'checkbox'} onClick={() => alert('click')}>
            <span>&#x2714;</span>
          </div>
          <input
            type="text"
            value={title}
            onKeyDown{e => handleKeyDown(e, i)}
            onChange{e => props.updateItem(e, i)}
        </div>
    );
}
