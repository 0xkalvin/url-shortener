import React, { useState } from 'react';

import './global.css';
import './App.css'
import './Short.css'
import './Long.css'
import './Main.css'

import api from './services/api';

function App() {

    const [long_url, setLongURL ] = useState('');
    const [short_url, setShortURL ] = useState('');

    async function handleSubmit(e){
        e.preventDefault();

        const response = await api.post('/short', { long_url });

        setShortURL(response.data.short_url);
        
    }

    async function openOriginalURL(e){
        e.preventDefault();
        try {
            const response = await api.get('/long/' + short_url );
            const url = response.data.long_url
            console.log(url);
            const win = window.open(url, '_blank');
            win.focus();
        } catch (error) {
         console.log(error);   
        }
    }


    function copyToClipboard(){
        navigator.clipboard.writeText('http://localhost:8080/long/' + short_url)
    }
    

  return (
    <div id="app">
        
        <h1>Atomic URL</h1>

        <main>
            
            <div id="long-menu">
                <form onSubmit={handleSubmit}>

                    <div className="long-item">
                        <label htmlFor="long_url">Paste a long URL</label>
                        <input 
                            name="long_url"
                            id="long_url"
                            required
                            value={long_url}
                            onChange={e => setLongURL(e.target.value)}
                        />         
                    </div>  
                    <button type="submit">Shorten</button>  
                </form>
            </div>
            
            <div id="short-menu">

                <p>Your shortened URL</p>
                <input 
                    name="short_url"
                    id="short_url"
                    value={ short_url ? 'https://atomic.url/' + short_url : ''}
                    onClick={copyToClipboard}
                />  
                <button 
                onClick={openOriginalURL}
                disabled={short_url === ''}
                type="button"
                >Go</button>  
            </div>
        </main>


        


    </div>
  );
}

export default App;