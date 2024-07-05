import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Login from './pages/Login'
import Home from './pages/Home'
import './App.css';

function App() {
  return (
 
      <div className="App">
        <BrowserRouter>
        <Routes>
        <Route index element={<Home/>} />
        <Route path ="/login" element={<Login/>}/>
        </Routes>
        </BrowserRouter>
      </div>
 
  );
}

export default App;
