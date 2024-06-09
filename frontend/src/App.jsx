import './App.css'
import Nav from './components/Nav'
import Login from './pages/Login'
import {BrowserRouter, Route, Routes} from 'react-router-dom'
import Register from './pages/Register'
import Home from './pages/Home'
import { useState, useEffect } from 'react'


function App() {

  const [name, setName] = useState('')

    useEffect(() => {

        const fetchData = async() => {
            const response = await fetch('http://localhost:8000/api/user', {
                headers: {'Content-Type': 'application/json'},
                credentials: 'include'
            })

            const user = await response.json()
            setName(user.name)
        }

        fetchData()
    }, [])
  
  return (
    <>
      <BrowserRouter>
        <Nav name={name} setName={setName} />

        <main className="form-signin w-100 m-auto">          
            <Routes>
              <Route path='/' element={<Home name={name} />} />
              <Route path='/login' element={<Login setName={setName} />} />
              <Route path='/register' element={<Register />} />
            </Routes>
        </main>
      </BrowserRouter>
    </>
  )
}

export default App
