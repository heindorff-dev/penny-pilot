import './App.css'

import { Route, Routes } from 'react-router-dom'
import { NavBar } from './components/ui/navigation/desktop/nav-bar'
import { AuthenticationGuard } from "./components/authentication-guard";

import HomePage from './pages/home-page'
import ProfilePage from './pages/profile-page'
import CallbackPage from './pages/callback-page';

function App() {

  return (
    <>
      <NavBar />
      <Routes>
        <Route path="/" element={<HomePage/>}/> {/* 👈 Renders at /app/ */}
        <Route path="/callback" element={<CallbackPage />} />
        <Route
          path="/profile"
          element={<AuthenticationGuard component={ProfilePage} />}
        />
      </Routes>
      
    </>
  )
}

export default App
