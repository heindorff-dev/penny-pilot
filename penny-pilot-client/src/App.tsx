import { Route, Routes } from 'react-router-dom'
import './App.css'
import HomePage from './pages/home-page'

function App() {

  return (
    <>
      <Routes>
        <Route path="/" element={<HomePage/>}/> {/* 👈 Renders at /app/ */}
      </Routes>
    </>
  )
}

export default App
