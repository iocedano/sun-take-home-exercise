import React, { useEffect, useState } from 'react';
import axios from 'axios';
import './App.css';
import Status from './components/status';

function App() {
  let [status, setStatus] = useState([]);
  
  useEffect(() => {
    axios.get('/v1/all-status').then((response) => {
      setStatus(response.data);
    })
  },[])

  return (
    <div className="App">
      <section className="App-header">
        <section>
          <h2>Vendors Status</h2>
          <ul className="status-container">
            {status.map((st) => <Status {...st} />)}
          </ul>
        </section>
      </section>
    </div>
  );
}

export default App;
