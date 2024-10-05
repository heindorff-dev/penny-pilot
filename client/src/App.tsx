import { Route, Routes } from 'react-router-dom'
import './App.css'
import HomePage from './pages/home-page'
import NavMenu from './components/ui/NavMenu'

function App() {

  return (
    <>
      <NavMenu />
      
      <Routes>
        <Route path="/" element={<HomePage/>}/> {/* 👈 Renders at /app/ */}
      </Routes>
    </>
  )
}

export default App
