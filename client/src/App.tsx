import { Route, Routes } from 'react-router-dom'
import './App.css'
import HomePage from './pages/home-page'
import LoginPage from './pages/login-page'
import NavMenu from './components/ui/nav-menu'
import PageContent from './components/ui/page-content'

function App() {

  return (
    <>
      <NavMenu />
      <Routes>
        <Route path="/" element={<HomePage/>}/> {/* 👈 Renders at /app/ */}
        <Route path="/login" element={<LoginPage/>}/> {/* 👈 Renders at /app/ */}
      </Routes>
      
    </>
  )
}

export default App
